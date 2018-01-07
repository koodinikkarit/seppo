package service

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	"github.com/jinzhu/gorm"
	SeppoService "github.com/koodinikkarit/go-clientlibs/seppo"
	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/managers"
	"github.com/koodinikkarit/seppo/models"
)

func (s *SeppoServiceServer) CreateVariation(
	ctx context.Context,
	in *SeppoService.CreateVariationRequest,
) (
	*SeppoService.CreateVariationResponse,
	error,
) {
	res := &SeppoService.CreateVariationResponse{}
	tx := s.getDB().Begin()

	var sameVariation models.Variation
	tx.Table("variations").
		Joins("left join variation_versions on variations.id = variation_versions.variation_id").
		Preload("VariationVersions").
		Where("variation_versions.name = ?", in.Name).
		Where("variation_versions.text = ?", in.Text).
		First(&sameVariation)

	if sameVariation.ID == 0 {
		newVariation, _ := managers.CreateNewVariation(
			tx,
			in.Name,
			in.Text,
		)
		res.Variation = generators.NewVariation(&newVariation)
		tx.Commit()
		return res, nil
	}

	sameVariationVersion := sameVariation.FindVariationVersionByNameAndText(
		in.Name,
		in.Text,
	)
	newestVariationVersion := sameVariation.FindNewestVersion()

	if sameVariationVersion.ID == newestVariationVersion.ID {
		if sameVariationVersion.DisabledAt == nil {
			res.Variation = generators.NewVariation(&sameVariation)
			return res, nil
		}
		newVariation, _ := managers.CreateBranchAndVariation(
			tx,
			sameVariationVersion.ID,
			in.Name,
			in.Text,
		)
		res.Variation = generators.NewVariation(newVariation)
		tx.Commit()
		return res, nil
	}

	newVariation, _ := managers.CreateBranchAndVariation(
		tx,
		sameVariationVersion.ID,
		in.Name,
		in.Text,
	)
	tx.Commit()
	res.Variation = generators.NewVariation(newVariation)
	return res, nil
}

func handleVariationUpdateIds(
	tx *gorm.DB,
	in *SeppoService.UpdateVariationRequest,
	variation *models.Variation,
) {
	if len(in.AddTagIds) > 0 {
		var tagVariations []models.TagVariation
		for _, id := range in.AddTagIds {
			tagVariations = append(
				tagVariations,
				models.TagVariation{
					VariationID: in.VariationId,
					TagID:       id,
				},
			)
		}
		managers.BatchAddTagsToVariation(
			tx,
			tagVariations,
		)
	}
	if len(in.RemoveTagIds) > 0 {
		tx.Where("variation_id = ?", in.VariationId).
			Where("tag_id in (?)", in.RemoveTagIds).
			Delete(&models.TagVariation{})
	}
	if len(in.AddSongDatabaseIds) > 0 {
		var songDatabaseVariations []models.SongDatabaseVariation
		for _, id := range in.AddSongDatabaseIds {
			songDatabaseVariations = append(
				songDatabaseVariations,
				models.SongDatabaseVariation{
					VariationID:    in.VariationId,
					SongDatabaseID: id,
				},
			)
		}
		managers.BatchAddVariationsToSongDatabase(
			tx,
			songDatabaseVariations,
		)
	}
	if len(in.RemoveSongDatabaseIds) > 0 {
		tx.Where("variation_id = (?)", in.VariationId).
			Where("song_database_id in (?)", in.RemoveSongDatabaseIds).
			Delete(&models.SongDatabaseVariation{})
	}

	variationUpdateMap := make(map[string]interface{})

	if in.SongId > 0 {
		variationUpdateMap["song_id"] = in.SongId
	}

	if in.LanguageId > 0 {
		variationUpdateMap["language_id"] = in.LanguageId
	}

	if len(variationUpdateMap) > 0 {
		tx.Model(variation).Updates(variationUpdateMap)
	}
}

func (s *SeppoServiceServer) UpdateVariation(
	ctx context.Context,
	in *SeppoService.UpdateVariationRequest,
) (
	*SeppoService.UpdateVariationResponse,
	error,
) {
	res := &SeppoService.UpdateVariationResponse{}
	tx := s.getDB().Begin()

	var variation models.Variation
	tx.Preload("VariationVersions").
		First(&variation, in.VariationId)

	if variation.ID == 0 {
		res.Success = false
		return res, nil
	}

	if in.Name == "" && in.Text == "" {
		handleVariationUpdateIds(
			tx,
			in,
			&variation,
		)
		tx.Commit()
		res.Success = true
		res.Variation = generators.NewVariation(&variation)
		return res, nil
	}

	currentNewestVariationVersion := variation.FindNewestVersion()

	var name string
	var text string
	if in.Name != "" {
		name = in.Name
	} else {
		name = currentNewestVariationVersion.Name
	}

	if in.Text != "" {
		text = in.Text
	} else {
		text = currentNewestVariationVersion.Text
	}

	if name == currentNewestVariationVersion.Name &&
		text == currentNewestVariationVersion.Text {

		handleVariationUpdateIds(
			tx,
			in,
			&variation,
		)
		tx.Commit()
		res.Success = true
		res.Variation = generators.NewVariation(&variation)
		return res, nil
	}

	var sameVariation models.Variation
	tx.Table("variations").
		Joins("left join variation_versions on variations.id = variation_versions.variation_id").
		Preload("VariationVersions").
		Where("variation_versions.name = ?", name).
		Where("variation_versions.text = ?", text).
		First(&sameVariation)

	if sameVariation.ID == 0 {
		tx.Model(&currentNewestVariationVersion).
			Update("disabled_at", time.Now())
		newVariationVersion := models.VariationVersion{
			VariationID: variation.ID,
			Name:        name,
			Text:        text,
			Version:     currentNewestVariationVersion.Version + 1,
		}
		tx.Create(&newVariationVersion)
		handleVariationUpdateIds(
			tx,
			in,
			&variation,
		)
		res.Variation = generators.NewVariation(&variation)
		tx.Commit()
		return res, nil
	}

	sameVariationVersion := sameVariation.FindVariationVersionByNameAndText(
		name,
		text,
	)
	newestVariationVersion := sameVariation.FindNewestVersion()

	if sameVariationVersion.ID == newestVariationVersion.ID {
		tx.Model(&currentNewestVariationVersion).
			Update("disabled_at", time.Now())
		newMerge := models.Merge{
			VariationVersion1ID:           currentNewestVariationVersion.ID,
			VariationVersion2ID:           sameVariationVersion.ID,
			DestinationVariationVersionID: sameVariationVersion.ID,
		}
		tx.Create(&newMerge)
		managers.MoveVariationReferences(
			tx,
			currentNewestVariationVersion.VariationID,
			sameVariation.ID,
		)

		res.Variation = generators.NewVariation(&sameVariation)
		res.Success = true
		tx.Commit()
		return res, nil
	}

	tx.Model(&currentNewestVariationVersion).
		Update("disabled_at", time.Now())

	newVariationVersion := models.VariationVersion{
		VariationID: variation.ID,
		Name:        name,
		Text:        text,
		Version:     currentNewestVariationVersion.Version + 1,
	}
	tx.Create(&newVariationVersion)
	handleVariationUpdateIds(
		tx,
		in,
		&variation,
	)
	tx.Commit()
	res.Variation = generators.NewVariation(&variation)
	res.Success = true
	return res, nil
}

func (s *SeppoServiceServer) RemoveVariation(
	ctx context.Context,
	in *SeppoService.RemoveVariationRequest,
) (
	*SeppoService.RemoveVariationResponse,
	error,
) {
	res := &SeppoService.RemoveVariationResponse{}
	tx := s.getDB().Begin()

	var n uint32
	tx.Table("variations").Where("id = ?", in.VariationId).Count(&n)
	if n > 0 {
		tx.Table("variation_versions").
			Where("variation_id = ?", in.VariationId).
			Where("disabled_at is null").
			Update("disabled_at", time.Now())
		res.Success = true
		tx.Commit()
	}
	res.Success = false
	return res, nil
}

func (s *SeppoServiceServer) FetchVariationById(
	ctx context.Context,
	in *SeppoService.FetchVariationByIdRequest,
) (
	*SeppoService.FetchVariationByIdResponse,
	error,
) {
	res := &SeppoService.FetchVariationByIdResponse{}
	db := s.getDB()

	var variations []models.Variation
	db.Where("id in (?)", in.VariationIds).
		Find(&variations)

	for _, variationID := range in.VariationIds {
		var found bool
		for _, variation := range variations {
			if variationID != variation.ID {
				continue
			}
			res.Variations = append(
				res.Variations,
				generators.NewVariation(&variation),
			)
			found = true
			break
		}
		if found == false {
			res.Variations = append(
				res.Variations,
				&SeppoService.Variation{},
			)
		}
	}

	return res, nil
}

func (s *SeppoServiceServer) SearchVariations(
	ctx context.Context,
	in *SeppoService.SearchVariationsRequest,
) (
	*SeppoService.SearchVariationsResponse,
	error,
) {
	res := &SeppoService.SearchVariationsResponse{}
	db := s.getDB()

	query := db.Table("variation_versions").Where("variation_versions.disabled_at is NULL").
		Joins("left join variations on variations.id = variation_versions.variation_id")

	if in.OrderBy > 0 {
		switch in.OrderBy {
		case 1:
			query = query.Order("name")
		case 2:
			query = query.Order("name desc")
		}
	}

	if in.TagId > 0 {
		query = query.Joins("left join tag_variations on tag_variations.variation_id = variations.id").
			Where("tag_variations.tag_id = ?", in.TagId)
	}

	if in.SongDatabaseId > 0 {
		query = query.Joins("left join song_database_variations on song_database_variations.variation_id = variations.id").
			Where("song_database_variations.song_database_id = ?", in.SongDatabaseId)
	}

	if in.ScheduleId > 0 {
		query = query.Joins("left join schedule_variations on schedule_variations.variation_id = variations.id").
			Where("schedule_variations.schedule_id = ?", in.ScheduleId)
	}

	if in.LanguageId > 0 {
		query = query.Where("variations.language_id = ?", in.LanguageId)
	}

	if len(in.SkipVariationIds) > 0 {
		query = query.Not("id", in.SkipVariationIds)
	}

	if in.SongDatabaseFilterId > 0 {
		query = query.Joins("song_database_variations sdv2 on sdv2.variation_id = variations.id").
			Not("sdv2.song_database_id = ?", in.SongDatabaseId)
	}

	query.Count(&res.MaxVariations)

	if in.SearchWord != "" {
		if in.SearchFrom > 0 {
			switch in.SearchFrom {
			case 1:
				query = query.Where("variation_versions.name LIKE ?", "%"+in.SearchWord+"%")
			case 2:
				query = query.
					Where("variations.name LIKE ? or variation_texts.text LIKE ?", "%"+in.SearchWord+"%", "%"+in.SearchWord+"%")
			}
		} else {
			query = query.Where("variation_versions.name LIKE ? OR variation_versions.text LIKE ?", "%"+in.SearchWord+"%", "%"+in.SearchWord+"%")
		}
	}

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}

	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	} else {
		query = query.Limit(10000)
	}

	var variations []models.Variation
	query = query.Select("variations.id, variations.song_id, variations.language_id").
		Scan(&variations)

	for _, variation := range variations {
		res.Variations = append(
			res.Variations,
			generators.NewVariation(&variation),
		)
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchNewestVariationVersionByVariationId(
	ctx context.Context,
	in *SeppoService.FetchNewestVariationVersionByVariationIdRequest,
) (
	*SeppoService.FetchNewestVariationVersionByVariationIdResponse,
	error,
) {
	res := &SeppoService.FetchNewestVariationVersionByVariationIdResponse{}
	db := s.getDB()

	var variationVersions []models.VariationVersion
	err := db.Table("variation_versions").
		Joins("left join variation_versions vv2 on (variation_versions.variation_id = vv2.variation_id and variation_versions.version < vv2.version)").
		Where("vv2.id is null").
		Where("variation_versions.variation_id in (?)", in.VariationIds).
		Find(&variationVersions).Error

	if err != nil {
		fmt.Println("err", err)
	}

	for _, variationID := range in.VariationIds {
		found := false
		for _, variationVersion := range variationVersions {
			if variationVersion.VariationID != variationID {
				continue
			}
			found = true
			res.VariationVersions = append(
				res.VariationVersions,
				generators.NewVariationVersion(&variationVersion),
			)
			break
		}
		if found == false {
			res.VariationVersions = append(
				res.VariationVersions,
				&SeppoService.VariationVersion{},
			)
		}
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchVariationVersionById(
	ctx context.Context,
	in *SeppoService.FetchVariationVersionByIdRequest,
) (
	*SeppoService.FetchVariationVersionByIdResponse,
	error,
) {
	res := &SeppoService.FetchVariationVersionByIdResponse{}
	db := s.getDB()

	var variationVersions []models.VariationVersion
	db.Where("id in (?)", in.VariationVersionIds).
		Find(&variationVersions)

	for _, variationVersionID := range in.VariationVersionIds {
		found := false
		for _, variationVersion := range variationVersions {
			if variationVersion.ID != variationVersionID {
				continue
			}
			found = true
			res.VariationVersions = append(
				res.VariationVersions,
				generators.NewVariationVersion(&variationVersion),
			)
			break
		}
		if found == false {
			res.VariationVersions = append(
				res.VariationVersions,
				&SeppoService.VariationVersion{},
			)
		}
	}

	return res, nil
}
