package services

import (
	"database/sql"
	"fmt"
	"time"

	null "gopkg.in/volatiletech/null.v6"

	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/managers"
	"github.com/koodinikkarit/seppo/models"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (s *SeppoServiceServer) CreateVariation(
	ctx context.Context,
	in *SeppoService.CreateVariationRequest,
) (
	*SeppoService.CreateVariationResponse,
	error,
) {
	res := &SeppoService.CreateVariationResponse{}
	newDb := s.getDB()
	defer newDb.Close()
	tx, _ := newDb.Begin()

	sameVariation, _ := models.Variations(
		tx,
		qm.InnerJoin("variation_versions on variations.id = variation_versions.variation_id"),
		qm.Where("variation_Versions.name = ?", in.Name),
		qm.Where("variation_versions.text = ?", in.Text),
		qm.Load("variation_versions"),
	).One()

	if sameVariation == nil {
		newVariation, _ := managers.CreateNewVariation(
			tx,
			in.Name,
			in.Text,
		)
		res.Variation = generators.NewVariation(newVariation)
		tx.Commit()
		return res, nil
	}

	sameVariationVersions, _ := sameVariation.VariationVersions(tx).All()
	sameVariationVersion := managers.FindVariationVersionByNameAndText(
		sameVariationVersions,
		in.Name,
		in.Text,
	)
	newestVariationVersion := managers.FindNewestVariationVersion(
		sameVariationVersions,
	)

	if sameVariationVersion.ID == newestVariationVersion.ID {
		if sameVariationVersion.DisabledAt.Valid == false {
			res.Variation = generators.NewVariation(sameVariation)
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

func HandleVariationUpdateIds(
	tx *sql.Tx,
	in *SeppoService.UpdateVariationRequest,
	variation *models.Variation,
) {
	for _, id := range in.AddTagIds {
		tagVariation := models.TagVariation{
			VariationID: in.VariationId,
			TagID:       id,
		}
		tagVariation.Insert(tx)
	}
	if len(in.RemoveTagIds) > 0 {
		models.TagVariations(
			tx,
			qm.WhereIn("variation_id = ?", in.VariationId),
			qm.Where("tag_id in ?", in.RemoveTagIds),
		).DeleteAll()
	}
	for _, id := range in.AddSongDatabaseIds {
		songDatabaseVariation := models.SongDatabaseVariation{
			VariationID:    in.VariationId,
			SongDatabaseID: id,
		}
		songDatabaseVariation.Insert(tx)
	}
	if len(in.RemoveSongDatabaseIds) > 0 {
		models.SongDatabaseVariations(
			tx,
			qm.Where("variation_id = ?", in.VariationId),
			qm.WhereIn("song_database_id in ?", in.RemoveSongDatabaseIds),
		)
	}

	if in.SongId > 0 {
		variation.SongID = null.NewUint64(in.SongId, true)
	}

	if in.LanguageId > 0 {
		variation.LanguageID = null.NewUint64(in.LanguageId, true)
	}
	if in.SongId > 0 &&
		in.LanguageId > 0 {
		variation.Update(tx, "song_id", "language_id")
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
	newDb := s.getDB()
	defer newDb.Close()
	tx, _ := newDb.Begin()
	variation, _ := models.Variations(
		tx,
		qm.Load("VariationVersions"),
		qm.Where("id = ?", in.VariationId),
	).One()

	if variation == nil {
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
		res.Variation = generators.NewVariation(variation)
		return res, nil
	}

	variationVersions, _ := variation.VariationVersions(tx).All()

	currentNewestVariationVersion := managers.FindNewestVariationVersion(
		variationVersions,
	)

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
		res.Variation = generators.NewVariation(variation)
		return res, nil
	}

	sameVariation, _ := models.Variations(
		tx,
		qm.InnerJoin("variation_versions on variations.id = variation_versions.variation_id"),
		qm.Where("variation_Versions.name = ?", in.Name),
		qm.Where("variation_versions.text = ?", in.Text),
		qm.Load("variation_versions"),
	).One()

	if sameVariation == nil {
		currentNewestVariationVersion.DisabledAt = null.NewTime(time.Now(), true)
		currentNewestVariationVersion.Update(tx)
		newVariationVersion := models.VariationVersion{
			VariationID: variation.ID,
			Name:        name,
			Text:        text,
			Version:     currentNewestVariationVersion.Version + 1,
		}
		newVariationVersion.Insert(tx)
		HandleVariationUpdateIds(
			tx,
			in,
			variation,
		)
		res.Variation = generators.NewVariation(variation)
		tx.Commit()
		return res, nil
	}

	sameVariationVersions, _ := sameVariation.VariationVersions(tx).All()
	sameVariationVersion := managers.FindVariationVersionByNameAndText(
		sameVariationVersions,
		name,
		text,
	)
	newestVariationVersion := managers.FindNewestVariationVersion(
		sameVariationVersions,
	)

	if sameVariationVersion.ID == newestVariationVersion.ID {
		currentNewestVariationVersion.DisabledAt = null.NewTime(time.Now(), true)
		currentNewestVariationVersion.Update(tx)
		newMerge := models.Merge{
			VariationVersion1ID:           currentNewestVariationVersion.ID,
			VariationVersion2ID:           sameVariationVersion.ID,
			DestinationVariationVersionID: sameVariationVersion.ID,
		}
		newMerge.Insert(tx)
		managers.MoveVariationReferences(
			tx,
			currentNewestVariationVersion.VariationID,
			sameVariation.ID,
		)

		res.Variation = generators.NewVariation(sameVariation)
		res.Success = true
		tx.Commit()
		return res, nil
	}

	currentNewestVariationVersion.DisabledAt = null.NewTime(time.Now(), true)
	currentNewestVariationVersion.Update(tx)

	newVariationVersion := models.VariationVersion{
		VariationID: variation.ID,
		Name:        name,
		Text:        text,
		Version:     currentNewestVariationVersion.Version + 1,
	}
	newVariationVersion.Insert(tx)
	HandleVariationUpdateIds(
		tx,
		in,
		variation,
	)
	tx.Commit()
	res.Variation = generators.NewVariation(variation)
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
	newDb := s.getDB()
	defer newDb.Close()

	variation, _ := models.FindVariation(
		newDb,
		in.VariationId,
	)

	if variation == nil {
		res.Success = false
		return res, nil
	}

	variation.VariationVersions(
		newDb,
		qm.Where("disabled_at is null"),
	).UpdateAll(models.M{
		"disabled_at": time.Now(),
	})
	res.Success = true
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

	variations, _ := models.Variations(
		newDb,
		qm.WhereIn("id in ?", in.VariationIds),
	).All()

	for _, variationID := range in.VariationIds {
		var found bool
		for _, variation := range variations {
			if variationID == variation.ID {
				continue
			}
			res.Variations = append(
				res.Variations,
				generators.NewVariation(variation),
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
	boil.DebugMode = true
	res := &SeppoService.SearchVariationsResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	var queryMods []qm.QueryMod

	queryMods = append(
		queryMods,
		qm.InnerJoin("variation_versions on variation_versions.variation_id = variations.id"),
		qm.Where("variation_versions.disabled_at is null"),
	)

	// if in.OrderBy > 0 {
	// 	switch in.OrderBy {
	// 	case 1:
	// 		queryMods = append(
	// 			queryMods,
	// 			qm.Order("name"),
	// 		)
	// 	case 2:
	// 		queryMods = append(
	// 			queryMods,
	// 			qm.Order("name desc"),
	// 		)
	// 	}

	// }

	if in.TagId > 0 {
		queryMods = append(
			queryMods,
			qm.InnerJoin("tag_variations tv on tv.variation_id = variations.id"),
			qm.Where("tag_variations.tag_id = ?", in.TagId),
		)
	}

	if in.SongDatabaseId > 0 {
		queryMods = append(
			queryMods,
			qm.InnerJoin("song_database_variations sdv on sdv.variation_id = variations.id"),
			qm.Where("sdv.song_database_id = ?", in.SongDatabaseId),
		)
	}

	if in.ScheduleId > 0 {
		queryMods = append(
			queryMods,
			qm.InnerJoin("schedule_variations sv on sv.variation_id = variations.id"),
			qm.Where("sv.schedule_id = ?", in.ScheduleId),
		)
	}

	if in.LanguageId > 0 {
		queryMods = append(
			queryMods,
			qm.Where("variations.language_id = ?", in.LanguageId),
		)
	}

	// if len(in.SkipVariationIds) > 0 {
	// 	query = query.Not("id", in.SkipVariationIds)
	// }

	if in.SongDatabaseFilterId > 0 {
		queryMods = append(
			queryMods,
			qm.InnerJoin("song_database_variations sdv2 on sdv2.variation_id = variations.id"),
			qm.Where("sdv2.song_database_id != ?", in.SongDatabaseId),
		)

		// var filterSongDatabaseVariationVersionIds []uint32
		// filterSongDatabaseVariations := []db.SongDatabaseVariation{}
		// newDb.Where("song_database_id = ?", in.SongDatabaseFilterId).
		// 	Select("variation_version_id").Find(&filterSongDatabaseVariations)
		// for _, v := range filterSongDatabaseVariations {
		// 	filterSongDatabaseVariationVersionIds = append(
		// 		filterSongDatabaseVariationVersionIds,
		// 		v.VariationID,
		// 	)
		// }
		// if filterSongDatabaseVariationVersionIds != nil {
		// 	query = query.Not(
		// 		"variation_versions.id",
		// 		filterSongDatabaseVariationVersionIds,
		// 	)
		// }
	}

	c, _ := models.VariationVersions(
		newDb,
		queryMods...,
	).Count()
	res.MaxVariations = uint64(c)

	if in.SearchWord != "" {
		if in.SearchFrom > 0 {
			switch in.SearchFrom {
			case 1:
				queryMods = append(
					queryMods,
					qm.Where("variation_versions.name LIKE ?", "%"+in.SearchWord+"%"),
				)
			case 2:
				queryMods = append(
					queryMods,
					qm.Where("variations.name LIKE ? or variation_texts.text LIKE ?", "%"+in.SearchWord+"%", "%"+in.SearchWord+"%"),
				)
			}
		} else {
			queryMods = append(
				queryMods,
				qm.Where("variation_versions.name LIKE ? OR variation_versions.text LIKE ?", "%"+in.SearchWord+"%", "%"+in.SearchWord+"%"),
			)
		}
	}

	if in.Offset > 0 {
		queryMods = append(
			queryMods,
			qm.Offset(int(in.Offset)),
		)
	}

	if in.Limit > 0 {
		queryMods = append(
			queryMods,
			qm.Limit(int(in.Limit)),
		)
	} else {
		queryMods = append(
			queryMods,
			qm.Limit(10000),
		)
	}

	variations, _ := models.Variations(
		newDb,
		queryMods...,
	).All()

	for _, variation := range variations {
		res.Variations = append(
			res.Variations,
			generators.NewVariation(variation),
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

	variationVersions, err := models.VariationVersions(
		newDb,
		qm.Where("disabled_at is null"),
		qm.WhereIn("variation_id in ?", in.VariationIds),
		// qm.InnerJoin("left join variation_versions vv on (vv.variation_id = variation_versions.variation_id and variation_versions.version < vv.versions"),
		// qm.Where("vv.id is null"),
		// qm.Where("variation_versions.variation_id in ?", in.VariationIds),
	).All()

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
				generators.NewVariationVersion(variationVersion),
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
	newDb := s.getDB()
	defer newDb.Close()

	variationVersions, _ := models.VariationVersions(
		newDb,
		qm.WhereIn("id in ?", in.VariationVersionIds),
	).All()

	for _, variationVersionID := range in.VariationVersionIds {
		found := false
		for _, variationVersion := range variationVersions {
			if variationVersion.ID != variationVersionID {
				continue
			}
			found = true
			res.VariationVersions = append(
				res.VariationVersions,
				generators.NewVariationVersion(variationVersion),
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
