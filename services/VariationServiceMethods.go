package services

import (
	"time"

	"golang.org/x/net/context"

	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/managers"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
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
	defer tx.Close()

	var sameVariation db.Variation

	tx.Table("variations").
		Joins("left join variation_versions on variations.id = variation_versions.variation_id").
		Preload("VariationVersions").
		Where("variation_versions.name = ?", in.Name).
		Where("variation_versions.text = ?", in.Text).
		First(&sameVariation)

	if sameVariation.ID > 0 {
		sameVariationVersion := sameVariation.FindVariationVersionByNameAndText(
			in.Name,
			in.Text,
		)
		newestVariationVersion := sameVariation.FindNewestVersion()

		if sameVariationVersion.ID == newestVariationVersion.ID {
			if sameVariationVersion.DisabledAt != nil {
				tx.Model(&sameVariationVersion).Update("disabled_at", nil)
			}
			res.Variation = NewVariation(&sameVariation)
		} else {
			newVariation := db.CreateNewVariationAndVersion(
				in.Name,
				in.Text,
			)
			tx.Create(&newVariation)
			newBranch := db.Branch{
				SourceVariationVersionID:      sameVariationVersion.ID,
				DestinationVariationVersionID: newVariation.VariationVersions[0].ID,
			}
			tx.Create(&newBranch)
			res.Variation = NewVariation(&newVariation)
		}

	} else {
		newVariation := db.CreateNewVariationAndVersion(
			in.Name,
			in.Text,
		)
		tx.Create(&newVariation)
		res.Variation = NewVariation(&newVariation)
	}
	tx.Commit()
	return res, nil
}

func HandleVariationUpdateIds(
	tx *gorm.DB,
	in *SeppoService.UpdateVariationRequest,
	variation db.Variation,
) {
	if len(in.AddTagIds) > 0 {
		var tagVariations []db.TagVariation
		for _, id := range in.AddTagIds {
			tagVariations = append(
				tagVariations,
				db.TagVariation{
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
			Delete(&db.TagVariation{})
	}
	if len(in.AddSongDatabaseIds) > 0 {
		var songDatabaseVariations []db.SongDatabaseVariation
		for _, id := range in.AddSongDatabaseIds {
			songDatabaseVariations = append(
				songDatabaseVariations,
				db.SongDatabaseVariation{
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
			Delete(&db.SongDatabaseVariation{})
	}

	variationUpdateMap := make(map[string]interface{})

	if in.SongId > 0 {
		variationUpdateMap["song_id"] = in.SongId
	}

	if in.LanguageId > 0 {
		variationUpdateMap["language_id"] = in.LanguageId
	}

	if len(variationUpdateMap) > 0 {
		tx.Model(&variation).Updates(variationUpdateMap)
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
	defer tx.Close()

	var variation db.Variation

	tx.Preload("VariationVersions").First(&variation, in.VariationId)

	if variation.ID == 0 {
		res.Success = false
		return res, nil
	}

	if in.Name == "" && in.Text == "" {
		HandleVariationUpdateIds(
			tx,
			in,
			variation,
		)
		tx.Commit()
		res.Success = true
		res.Variation = NewVariation(&variation)
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

		HandleVariationUpdateIds(
			tx,
			in,
			variation,
		)
		tx.Commit()
		res.Success = true
		res.Variation = NewVariation(&variation)
		return res, nil
	}

	var sameVariation db.Variation

	tx.Table("variations").
		Joins("left join variation_versions on variations.id = variation_versions.variation_id").
		Preload("VariationVersions").
		Where("variation_versions.name = ?", name).
		Where("variation_versions.text = ?", text).
		First(&sameVariation)

	if sameVariation.ID > 0 {
		sameVariationVersion := sameVariation.FindVariationVersionByNameAndText(
			name,
			text,
		)
		newestVariationVersion := sameVariation.FindNewestVersion()
		if sameVariationVersion.ID != currentNewestVariationVersion.ID {
			if sameVariationVersion.ID == newestVariationVersion.ID {
				tx.Model(&currentNewestVariationVersion).Update(
					"disabled_at",
					time.Now(),
				)
				newMerge := db.Merge{
					VariationVersion1ID:           currentNewestVariationVersion.ID,
					VariationVersion2ID:           sameVariationVersion.ID,
					DestinationVariationVersionID: sameVariationVersion.ID,
				}
				tx.Create(&newMerge)
				db.MoveVariationReferences(
					tx,
					currentNewestVariationVersion.VariationID,
					sameVariation.ID,
				)
			} else {
				tx.Model(&currentNewestVariationVersion).Update(
					"disabled_at",
					time.Now(),
				)
				newVariationVersion := db.VariationVersion{
					VariationID: in.VariationId,
					Name:        name,
					Text:        text,
					Version:     currentNewestVariationVersion.Version + 1,
				}
				tx.Create(&newVariationVersion)
				HandleVariationUpdateIds(
					tx,
					in,
					variation,
				)
			}
		}
	} else {
		tx.Model(&currentNewestVariationVersion).Update(
			"disabled_at",
			time.Now(),
		)
		newVariationVersion := db.VariationVersion{
			VariationID: in.VariationId,
			Name:        name,
			Text:        text,
			Version:     currentNewestVariationVersion.Version + 1,
		}
		tx.Create(&newVariationVersion)
		HandleVariationUpdateIds(
			tx,
			in,
			variation,
		)
	}

	tx.Commit()

	res.Variation = NewVariation(&variation)
	res.Success = true

	// now := time.Now()

	// var variation db.Variation
	// tx.First(&variation, in.VariationId)

	// if variation.ID == 0 {
	// 	res.Success = false
	// 	return res, nil
	// }

	// if in.SongId > 0 {
	// 	variation.SongID = &in.SongId
	// }

	// if in.LanguageId > 0 {
	// 	variation.LanguageID = &in.LanguageId
	// }

	// tx.Save(&variation)

	// var newestVariationVersion db.VariationVersion
	// tx.Table("variation_versions").
	// 	Where("variation_versions.variation_id = ?", variation.ID).
	// 	Where("variation_versions.version = (select max(version) from variation_versions where variation_versions.variation_id = ?)", variation.ID).
	// 	First(&newestVariationVersion)

	// var name string
	// var text string
	// if in.Name != "" {
	// 	name = in.Name
	// } else {
	// 	name = newestVariationVersion.Name
	// }

	// if in.Text != "" {
	// 	text = in.Text
	// } else {
	// 	text = newestVariationVersion.Text
	// }

	// sameVariationVersions := []db.VariationVersion{}

	// tx.Where("name = ?", name).Where("text = ?", text).Find(&sameVariationVersions)

	// if len(sameVariationVersions) > 0 {
	// 	sameVariationVersion := sameVariationVersions[0]
	// 	if sameVariationVersion.ID == newestVariationVersion.ID {
	// 		res.Variation = NewVariation(&variation)
	// 	} else {

	// 		MoveVariationVersionReferences(
	// 			tx,
	// 			newestVariationVersion.ID,
	// 			sameVariationVersion.ID,
	// 		)
	// 		newMerge := db.Merge{
	// 			VariationVersion1ID:           newestVariationVersion.ID,
	// 			VariationVersion2ID:           sameVariationVersion.ID,
	// 			DestinationVariationVersionID: sameVariationVersion.ID,
	// 		}
	// 		tx.Create(&newMerge)
	// 		var sameVariation db.Variation
	// 		tx.First(&sameVariation, sameVariationVersion.VariationID)
	// 		newestVariationVersion.DisabledAt = &now
	// 		tx.Save(&newestVariationVersion)
	// 		res.Variation = NewVariation(&sameVariation)
	// 	}
	// } else {
	// 	newVariationVersion := db.VariationVersion{
	// 		VariationID: newestVariationVersion.VariationID,
	// 		Name:        name,
	// 		Text:        text,
	// 		Version:     newestVariationVersion.Version + 1,
	// 	}
	// 	tx.Create(&newVariationVersion)
	// 	MoveVariationVersionReferences(
	// 		tx,
	// 		newestVariationVersion.ID,
	// 		newVariationVersion.ID,
	// 	)
	// 	newestVariationVersion.DisabledAt = &now
	// 	tx.Save(&newestVariationVersion)
	// 	tx.Save(&variation)
	// 	res.Variation = NewVariation(&variation)
	// }
	// tx.Commit()
	// res.Success = true

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
	defer tx.Close()

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

	// var newestVariationVersion db.VariationVersion
	// tx.Table("variation_versions").
	// 	Where("variation_versions.variation_id = ?", in.VariationId).
	// 	Where("variation_versions.version = (select max(version) from variation_versions where variation_versions.variation_id = ?)", in.VariationId).
	// 	First(&newestVariationVersion)

	// if newestVariationVersion.ID > 0 {
	// 	res.Success = true
	// 	now := time.Now()
	// 	newestVariationVersion.DisabledAt = &now
	// 	tx.Save(&newestVariationVersion)
	// } else {
	// 	res.Success = false
	// }

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
	newDb := s.getDB()
	defer newDb.Close()

	variations := []db.Variation{}
	newDb.Where("id in (?)", in.VariationIds).Find(&variations)

	for _, variationId := range in.VariationIds {
		var found bool
		for _, variation := range variations {
			if variation.ID == variationId {
				res.Variations = append(
					res.Variations,
					NewVariation(&variation),
				)
				found = true
				break
			}
		}
		if found == false {
			res.Variations = append(res.Variations, &SeppoService.Variation{
				Id: 0,
			})
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
	newDb := s.getDB()
	defer newDb.Close()

	variations := []db.Variation{}

	query := newDb.Table("variation_versions").Where("variation_versions.disabled_at is NULL").
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
		var filterSongDatabaseVariationVersionIds []uint32
		filterSongDatabaseVariations := []db.SongDatabaseVariation{}
		newDb.Where("song_database_id = ?", in.SongDatabaseFilterId).
			Select("variation_version_id").Find(&filterSongDatabaseVariations)
		for _, v := range filterSongDatabaseVariations {
			filterSongDatabaseVariationVersionIds = append(
				filterSongDatabaseVariationVersionIds,
				v.VariationID,
			)
		}
		if filterSongDatabaseVariationVersionIds != nil {
			query = query.Not(
				"variation_versions.id",
				filterSongDatabaseVariationVersionIds,
			)
		}
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
	}

	query = query.Select("variations.id, variations.song_id, variations.language_id").Scan(&variations)

	for i := 0; i < len(variations); i++ {
		res.Variations = append(
			res.Variations,
			NewVariation(&variations[i]),
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
	newDb := s.getDB()
	defer newDb.Close()

	variationVersions := []db.VariationVersion{}

	// select v1.id,v1.name,v1.version,v2.id,v2.name,v2.version
	// from variation_versions v1
	// left join variation_versions v2
	// on v1.variation_id = v2.variation_id
	// and v1.version < v2.version
	// where v2.id is null
	// and v1.variation_id in (1,2)

	// newDb.Table("variation_versions as vv1").
	// 	Joins("left join variation_versions as vv2 on vv1.variation_id = vv2.variation_id").
	// 	Where("vv1.version >= vv2.version").
	// 	Where("vv1.variation_id in (?)", in.VariationIds).
	// 	Select("vv1.*").
	// 	Find(&variationVersions)

	newDb.Where("variation_id in (?)", in.VariationIds).
		Where("disabled_at is null").
		Find(&variationVersions)

	for _, variationID := range in.VariationIds {
		found := false
		for _, variationVersion := range variationVersions {
			if variationVersion.VariationID == variationID {
				found = true
				res.VariationVersions = append(
					res.VariationVersions,
					NewVariationVersion(&variationVersion),
				)
			}
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
	newDb := s.getDB()
	defer newDb.Close()

	variationVersions := []db.VariationVersion{}
	newDb.Where("id in (?)", in.VariationVersionIds).Find(&variationVersions)

	for i := 0; i < len(in.VariationVersionIds); i++ {
		found := false
		for j := 0; j < len(variationVersions); j++ {
			if in.VariationVersionIds[i] == variationVersions[j].ID {
				found = true
				res.VariationVersions = append(
					res.VariationVersions,
					NewVariationVersion(&variationVersions[j]),
				)
			}
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
