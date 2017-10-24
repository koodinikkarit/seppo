package services

import (
	"time"

	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
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

	variationVersions := []db.VariationVersion{}
	tx.Where("name = ?", in.Name).
		Where("text = ?", in.Text).
		Find(&variationVersions)

	if len(variationVersions) > 0 {
		variationVersion := variationVersions[0]
		var maxVersion uint32
		tx.Table("variation_versions").
			Where("variation_id = ?", variationVersion.VariationID).
			Count(&maxVersion)

		var variation db.Variation

		if maxVersion == variationVersion.Version {
			tx.First(&variation, variationVersion.VariationID)
			res.Variation = NewVariation(&variation)
		} else {
			newVariation := db.Variation{}
			tx.Create(&newVariation)
			newVariationVersion := &db.VariationVersion{
				VariationID: newVariation.ID,
				Name:        in.Name,
				Text:        in.Text,
				Version:     1,
			}
			tx.Create(&newVariationVersion)
			newVariation.VariationVersionID = &newVariationVersion.ID
			tx.Save(&newVariation)
			newBranch := db.Branch{
				SourceVariationVersionID:      variationVersion.ID,
				DestinationVariationVersionID: newVariationVersion.ID,
			}
			tx.Create(&newBranch)
			res.Variation = NewVariation(&newVariation)
		}
	} else {
		newVariation := db.Variation{}
		tx.Create(&newVariation)
		newVariationVersion := &db.VariationVersion{
			VariationID: newVariation.ID,
			Name:        in.Name,
			Text:        in.Text,
			Version:     1,
		}
		tx.Create(&newVariationVersion)
		newVariation.VariationVersionID = &newVariationVersion.ID
		tx.Save(&newVariation)
		res.Variation = NewVariation(&newVariation)
	}
	tx.Commit()

	return res, nil
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

	now := time.Now()

	var variation db.Variation
	tx.First(&variation, in.VariationId)

	if variation.ID > 0 {
		var newestVariationVersion db.VariationVersion
		tx.Table("variation_versions").
			Where("variation_versions.variation_id = ?", variation.ID).
			Where("variation_versions.version = (select max(version) from variation_versions where variation_versions.variation_id = ?)", variation.ID).
			First(&newestVariationVersion)

		var name string
		var text string
		if in.Name != "" {
			name = in.Name
		} else {
			name = newestVariationVersion.Name
		}

		if in.Text != "" {
			text = in.Text
		} else {
			text = newestVariationVersion.Text
		}

		sameVariationVersions := []db.VariationVersion{}

		tx.Where("name = ?", name).Where("text = ?", text).Find(&sameVariationVersions)

		if len(sameVariationVersions) > 0 {
			sameVariationVersion := sameVariationVersions[0]
			if sameVariationVersion.ID == newestVariationVersion.ID {
				res.Variation = NewVariation(&variation)
			} else {
				MoveVariationVersionReferences(
					tx,
					newestVariationVersion.ID,
					sameVariationVersion.ID,
				)
				newMerge := db.Merge{
					VariationVersion1ID:           newestVariationVersion.ID,
					VariationVersion2ID:           sameVariationVersion.ID,
					DestinationVariationVersionID: sameVariationVersion.ID,
				}
				tx.Create(&newMerge)
				var sameVariation db.Variation
				tx.First(&sameVariation, sameVariationVersion.VariationID)
				newestVariationVersion.DisabledAt = &now
				tx.Save(&newestVariationVersion)
				res.Variation = NewVariation(&sameVariation)
			}
		} else {
			newVariationVersion := db.VariationVersion{
				VariationID: newestVariationVersion.VariationID,
				Name:        name,
				Text:        text,
				Version:     newestVariationVersion.Version + 1,
			}
			tx.Create(&newVariationVersion)
			MoveVariationVersionReferences(
				tx,
				newestVariationVersion.ID,
				newVariationVersion.ID,
			)
			newestVariationVersion.DisabledAt = &now
			tx.Save(&newestVariationVersion)
			variation.VariationVersionID = &newVariationVersion.ID
			tx.Save(&variation)
			res.Variation = NewVariation(&variation)
		}
		tx.Commit()
	}

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

	var newestVariationVersion db.VariationVersion
	tx.Table("variation_versions").
		Where("variation_versions.variation_id = ?", in.VariationId).
		Where("variation_versions.version = (select max(version) from variation_versions where variation_versions.variation_id = ?)", in.VariationId).
		First(&newestVariationVersion)

	if newestVariationVersion.ID > 0 {
		res.Success = true
		now := time.Now()
		newestVariationVersion.DisabledAt = &now
		tx.Save(&newestVariationVersion)
	} else {
		res.Success = false
	}

	tx.Commit()

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
		query = query.Joins("left join tag_variations on tag_variations.variation_version_id = variation_versions.id").
			Where("tag_variations.tag_id = ?", in.TagId)
	}

	if in.SongDatabaseId > 0 {
		query = query.Joins("left join song_database_variations on song_database_variations.variation_version_id = variation_versions.id").
			Where("song_database_variations.song_database_id = ?", in.SongDatabaseId)
	}

	if in.ScheduleId > 0 {
		query = query.Joins("left join schedule_variations on schedule_variations.variation_version_id = variation_versions.id").
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
				v.VariationVersionID,
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
			query = query.Where("variation_versions.name LIKE ?", "%"+in.SearchWord+"%").
				Where("variation_versions.text LIKE ?", "%"+in.SearchWord+"%")
		}
	}

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}

	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	}

	query = query.Select("variations.id, variations.song_id, variations.language_id, variations.variation_version_id").Scan(&variations)

	for i := 0; i < len(variations); i++ {
		res.Variations = append(
			res.Variations,
			NewVariation(&variations[i]),
		)
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
			found = true
			res.VariationVersions = append(
				res.VariationVersions,
				NewVariationVersion(&variationVersions[j]),
			)
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
