package seppo

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
)

func (s *SeppoServiceServer) FetchVariationById(ctx context.Context, in *SeppoService.FetchVariationByIdRequest) (*SeppoService.FetchVariationByIdResponse, error) {
	res := &SeppoService.FetchVariationByIdResponse{}

	variations := []SeppoDB.Variation{}
	s.databaseService.GetDb().Where("id in (?)", in.VariationIds).Find(&variations)
	for _, variationId := range in.VariationIds {
		var found bool
		for _, variation := range variations {
			if variation.ID == variationId {
				res.Variations = append(res.Variations, NewVariationToServiceType(&variation))
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
	variations := []SeppoDB.Variation{}

	query := s.databaseService.GetDb().Table("variations")

	if in.TagId > 0 {
		query = query.Joins("JOIN tag_variations on tag_variations.variation_id = variations.id").
			Where("tag_variations.tag_id = ?", in.TagId)
	}

	if in.SongDatabaseId > 0 {
		query = query.Joins("JOIN song_database_variations on song_database_variations.variation_id = variations.id").
			Where("song_database_variations.song_database_id = ?", in.SongDatabaseId)
	}

	if in.ScheduleId > 0 {
		query = query.Joins("JOIN schedule_variations on schedule_variations.variation_id = variations.id").
			Where("schedule_variations.schedule_id = ?", in.ScheduleId)
	}

	if in.LanguageId > 0 {
		query = query.Where("variations.language_id = ?", in.LanguageId)
	}

	if len(in.SkipVariationIds) > 0 {
		query = query.Not("id", in.SkipVariationIds)
	}

	if in.SongDatabaseFilterId > 0 {
		var filterSongDatabaseVariationsIds []uint32
		filterSongDatabaseVariations := []SeppoDB.SongDatabaseVariation{}
		s.databaseService.GetDb().
			Where("song_database_id = ?", in.SongDatabaseFilterId).
			Select("variation_id").Find(&filterSongDatabaseVariations)
		for _, v := range filterSongDatabaseVariations {
			filterSongDatabaseVariationsIds = append(filterSongDatabaseVariationsIds, v.VariationID)
		}
		if filterSongDatabaseVariationsIds != nil {
			query = query.Not("id", filterSongDatabaseVariationsIds)
		}
	}

	query.Count(&res.MaxVariations)

	if in.SearchWord != "" {
		query = query.Joins("JOIN variation_texts ON variation_texts.variation_id = variations.id").
			Where("variations.name LIKE ? or variation_texts.text LIKE ?", "%"+in.SearchWord+"%", "%"+in.SearchWord+"%")
	}

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}

	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	}

	query = query.Select("variations.id, variations.name, variations.song_id, variations.version").Find(&variations)

	for _, variation := range variations {
		res.Variations = append(res.Variations, &SeppoService.Variation{
			Id:      variation.ID,
			Name:    variation.Name,
			SongId:  variation.SongID,
			Version: variation.Version,
		})
	}

	return res, nil
}

func (s *SeppoServiceServer) ListenForChangedEwSong(in *SeppoService.ListenForChangedEwSongRequest, stream SeppoService.Seppo_ListenForChangedEwSongServer) error {

	return nil
}

func (s *SeppoServiceServer) FetchSongDatabases(ctx context.Context, in *SeppoService.FetchSongDatabasesRequest) (*SeppoService.SongDatabasesConnection, error) {
	res := &SeppoService.SongDatabasesConnection{}
	songDatabases := []SeppoDB.SongDatabase{}

	query := s.databaseService.GetDb().Table("song_databases")

	if in.VariationId > 0 {
		res.Id = in.VariationId
		query = query.Joins("JOIN song_database_variations on song_database_variations.song_database_id = song_databases.id").
			Where("song_database_variations.variation_id = ?", in.VariationId)
	}

	query.Count(&res.MaxSongDatabases)

	if in.SearchWord != "" {
		query = query.Where("song_databases.name LIKE ?", "%"+in.SearchWord+"%")
	}

	query = query.Find(&songDatabases)

	for _, songDatabase := range songDatabases {
		res.SongDatabases = append(res.SongDatabases, NewSongDatabaseToServiceType(&songDatabase))
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchSongDatabaseById(ctx context.Context, in *SeppoService.FetchSongDatabaseByIdRequest) (*SeppoService.FetchSongDatabaseByIdResponse, error) {
	res := &SeppoService.FetchSongDatabaseByIdResponse{}

	songDatabases := []SeppoDB.SongDatabase{}
	s.databaseService.GetDb().Where("id in (?)", in.SongDatabaseIds).Find(&songDatabases)

	for _, songDatabaseId := range in.SongDatabaseIds {
		found := false
		for _, songDatabase := range songDatabases {
			if songDatabase.ID == songDatabaseId {
				found = true
				res.SongDatabases = append(res.SongDatabases, NewSongDatabaseToServiceType(&songDatabase))
			}
		}
		if found == false {
			res.SongDatabases = append(res.SongDatabases, &SeppoService.SongDatabase{
				Id: 0,
			})
		}
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchEwDatabases(ctx context.Context, in *SeppoService.FetchEwDatabasesRequest) (*SeppoService.EwDatabasesConnection, error) {
	res := &SeppoService.EwDatabasesConnection{}

	ewDatabases := []SeppoDB.EwDatabase{}
	s.databaseService.GetDb().Find(&ewDatabases)

	for _, ewDatabase := range ewDatabases {
		res.EwDatabases = append(res.EwDatabases, NewEwDatabaseToServiceType(&ewDatabase))
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchEwDatabaseById(ctx context.Context, in *SeppoService.FetchEwDatabaseByIdRequest) (*SeppoService.FetchEwDatabaseByIdResponse, error) {
	res := &SeppoService.FetchEwDatabaseByIdResponse{}

	ewDatabases := []SeppoDB.EwDatabase{}
	s.databaseService.GetDb().Where("id in (?)", in.EwDatabaseIds).Find(&ewDatabases)
	for _, ewDatabaseId := range in.EwDatabaseIds {
		found := false
		for _, ewDatabase := range ewDatabases {
			if ewDatabaseId == ewDatabase.ID {
				found = true
				res.EwDatabases = append(res.EwDatabases, NewEwDatabaseToServiceType(&ewDatabase))
			}
		}
		if found == false {
			res.EwDatabases = append(res.EwDatabases, &SeppoService.EwDatabase{
				Id: 0,
			})
		}
	}
	return res, nil
}

func (s *SeppoServiceServer) FetchVariationsBySongDatabaseId(ctx context.Context, in *SeppoService.FetchVariationsBySongDatabaseIdRequest) (*SeppoService.FetchVariationsBySongDatabaseIdResponse, error) {
	res := &SeppoService.FetchVariationsBySongDatabaseIdResponse{}

	fetchedSongDatabaseVariations := []SeppoDB.SongDatabaseVariation{}
	s.databaseService.GetDb().Where("song_database_id in (?)", in.SongDatabaseIds).Find(&fetchedSongDatabaseVariations)

	variationIds := []uint32{}
	for _, v := range fetchedSongDatabaseVariations {
		variationIds = append(variationIds, v.VariationID)
	}
	variations := []SeppoDB.Variation{}
	s.databaseService.GetDb().Where("id in (?)", variationIds).Find(&variations)

	for _, songDatabaseId := range in.SongDatabaseIds {
		databaseVariations := SeppoService.SongDatabaseVariations{}
		databaseVariations.SongDatabaseId = songDatabaseId
		for _, songDatabaseVariation := range fetchedSongDatabaseVariations {
			if songDatabaseId == songDatabaseVariation.SongDatabaseID {
				for _, variation := range variations {
					if songDatabaseVariation.VariationID == variation.ID {
						databaseVariations.Variations = append(databaseVariations.Variations, NewVariationToServiceType(&variation))
					}
				}
			}
		}
		res.SongDatabaseVariations = append(res.SongDatabaseVariations, &databaseVariations)
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchVariationTextByVariationId(ctx context.Context, in *SeppoService.FetchVariationTextByVariationIdRequest) (*SeppoService.FetchVariationTextByVariationIdResponse, error) {
	res := &SeppoService.FetchVariationTextByVariationIdResponse{}

	variationTexts := []SeppoDB.VariationText{}

	s.databaseService.GetDb().Where("variation_id in (?)", in.VariationIds).Find(&variationTexts)

	for _, variationId := range in.VariationIds {
		found := false
		for _, variationText := range variationTexts {
			if variationText.VariationID == variationId {
				found = true
				res.VariationTexts = append(res.VariationTexts, NewVariationTextToServiceType(&variationText))
			}
		}

		if found == false {
			res.VariationTexts = append(res.VariationTexts, &SeppoService.VariationText{
				Id: 0,
			})
		}
	}

	return res, nil
}

func (s *SeppoServiceServer) SearchTags(ctx context.Context, in *SeppoService.SearchTagsRequest) (*SeppoService.TagsConnection, error) {
	res := &SeppoService.TagsConnection{}

	tags := []SeppoDB.Tag{}

	query := s.databaseService.GetDb().Table("tags")

	if in.SongDatabaseId > 0 {
		res.Id = in.SongDatabaseId
		query = query.Joins("JOIN song_database_tags on song_database_tags.tag_id = tags.id").
			Where("song_database_tags.song_database_id = ?", in.SongDatabaseId)
	}

	if in.VariationId > 0 {
		res.Id = in.VariationId
		query = query.Joins("JOIN tag_variations on tag_variations.tag_id = tags.id").
			Where("tag_variations.variation_id = ?", in.VariationId)
	}

	query.Count(&res.MaxTags)

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}

	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	}

	query.Find(&tags)

	for i := 0; i < len(tags); i++ {
		res.Tags = append(res.Tags, NewTagToServiceType(&tags[i]))
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchTagById(ctx context.Context, in *SeppoService.FetchTagByIdRequest) (*SeppoService.FetchTagByIdResponse, error) {
	res := &SeppoService.FetchTagByIdResponse{}

	tags := []SeppoDB.Tag{}

	s.databaseService.GetDb().Where("id in (?)", in.TagIds).Find(&tags)

	for _, tagID := range in.TagIds {
		found := false
		for i := 0; i < len(tags); i++ {
			if tagID == tags[i].ID {
				found = true
				res.Tags = append(res.Tags, NewTagToServiceType(&tags[i]))
			}
		}
		if found == false {
			res.Tags = append(res.Tags, &SeppoService.Tag{
				Id: 0,
			})
		}
	}

	return res, nil
}

func (s *SeppoServiceServer) SearchLanguages(ctx context.Context, in *SeppoService.SearchLanguagesRequest) (*SeppoService.LanguagesConnection, error) {
	res := &SeppoService.LanguagesConnection{}

	languages := []SeppoDB.Language{}

	q := s.databaseService.GetDb().Find(&languages)

	q.Count(&res.MaxLanguages)

	for i := 0; i < len(languages); i++ {
		res.Languages = append(res.Languages, NewLanguageToServiceType(&languages[i]))
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchLanguageById(ctx context.Context, in *SeppoService.FetchLanguageByIdRequest) (*SeppoService.FetchLanguageByIdResponse, error) {
	res := &SeppoService.FetchLanguageByIdResponse{}

	languages := []SeppoDB.Language{}

	s.databaseService.GetDb().Where("id in (?)", in.LanguageIds).Find(&languages)

	for _, languageID := range in.LanguageIds {
		found := false
		for i := 0; i < len(languages); i++ {
			if languageID == languages[i].ID {
				found = true
				res.Languages = append(res.Languages, NewLanguageToServiceType(&languages[i]))
			}
		}
		if found == false {
			res.Languages = append(res.Languages, &SeppoService.Language{
				Id: 0,
			})
		}
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchVariationTags(
	ctx context.Context,
	in *SeppoService.FetchVariationTagsRequest,
) (
	*SeppoService.FetchVariationTagsResponse,
	error,
) {
	res := &SeppoService.FetchVariationTagsResponse{}

	tags := []SeppoDB.Tag{}
	tagVariations := []SeppoDB.TagVariation{}

	s.databaseService.GetDb().
		Where("variation_id in (?)", in.VariationIds).
		Find(&tagVariations)

	var tagIds []uint32
	for i := 0; i < len(tagVariations); i++ {
		tagIds = append(tagIds, tagVariations[i].TagID)
	}

	s.databaseService.GetDb().Where("id in (?)", tagIds).Find(&tags)

	for _, variationId := range in.VariationIds {
		variationTags := &SeppoService.VariationTags{
			VariationId: variationId,
		}
		for i := 0; i < len(tagVariations); i++ {
			if tagVariations[i].VariationID == variationId {
				for j := 0; j < len(tags); j++ {
					if tags[j].ID == tagVariations[i].TagID {
						variationTags.Tags = append(variationTags.Tags, NewTagToServiceType(&tags[j]))
						break
					}
				}
			}
		}
		res.VariationTags = append(res.VariationTags, variationTags)
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchTagVariations(
	ctx context.Context,
	in *SeppoService.FetchTagVariationsRequest,
) (
	*SeppoService.FetchTagVariationsResponse,
	error,
) {
	res := &SeppoService.FetchTagVariationsResponse{}

	tagVariations := []SeppoDB.TagVariation{}
	variations := []SeppoDB.Variation{}

	s.databaseService.GetDb().
		Where("tag_id in (?)", in.TagIds).
		Find(&tagVariations)

	var variationIds []uint32
	for i := 0; i < len(tagVariations); i++ {
		variationIds = append(
			variationIds,
			tagVariations[i].VariationID,
		)
	}

	s.databaseService.GetDb().
		Where("id in (?)", variationIds).
		Find(&variations)

	for _, tagId := range in.TagIds {
		tagVariations2 := &SeppoService.TagVariations{
			TagId: tagId,
		}
		for i := 0; i < len(tagVariations); i++ {
			if tagVariations[i].TagID == tagId {
				for j := 0; j < len(variations); j++ {
					if variations[j].ID == tagVariations[i].VariationID {
						tagVariations2.Variations = append(tagVariations2.Variations, NewVariationToServiceType(&variations[j]))
						break
					}
				}
			}
		}
		res.TagVariations = append(res.TagVariations, tagVariations2)
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchTagSongDatabases(ctx context.Context, in *SeppoService.FetchTagSongDatabasesRequest) (*SeppoService.FetchTagSongDatabasesResponse, error) {
	res := &SeppoService.FetchTagSongDatabasesResponse{}

	//songDatabases := []SeppoDB.SongDatabase{}

	// s.databaseService.GetDb().Model(&SeppoDB.SongDatabaseTag{}).Where("song_database_tags.tag_id = ?", in.TagId).Related(&songDatabases)

	// for i := 0; i < len(songDatabases); i++ {
	// 	res.SongDatabases = append(res.SongDatabases, NewSongDatabaseToServiceType(&songDatabases[i]))
	// }

	return res, nil
}

func (s *SeppoServiceServer) FetchSongDatabaseTags(ctx context.Context, in *SeppoService.FetchSongDatabaseTagsRequest) (*SeppoService.FetchSongDatabaseTagsResponse, error) {
	res := &SeppoService.FetchSongDatabaseTagsResponse{}

	//tags := []SeppoDB.Tag{}

	// s.databaseService.GetDb().Model(&SeppoDB.SongDatabaseTag{}).Where("song_database_tags.song_database_Id = ?", in.SongDatabaseId).Related(&tags)

	// for i := 0; i < len(tags); i++ {
	// 	res.Tags = append(res.Tags, NewTagToServiceType(&tags[i]))
	// }

	return res, nil
}

func (s *SeppoServiceServer) FetchLanguageVariations(ctxt context.Context, in *SeppoService.FetchLanguageVariationsRequest) (*SeppoService.FetchLanguageVariationsResponse, error) {
	res := &SeppoService.FetchLanguageVariationsResponse{}

	//variations := []SeppoDB.Variation{}

	// s.databaseService.GetDb().Model(&SeppoDB.Language{}).Where("variations.language_id = ?", in.LanguageId).Related(&variations)

	return res, nil
}

func (s *SeppoServiceServer) FetchTagsBySongDatabaseById(
	ctx context.Context,
	in *SeppoService.FetchTagsBySongDatabaseIdRequest,
) (
	*SeppoService.FetchTagsBySongDatabaseIdResponse,
	error,
) {
	res := &SeppoService.FetchTagsBySongDatabaseIdResponse{}

	return res, nil
}

func (s *SeppoServiceServer) SearchSchedules(
	ctx context.Context,
	in *SeppoService.SearchSchedulesRequest,
) (
	*SeppoService.SearchSchedulesResponse,
	error,
) {
	res := &SeppoService.SearchSchedulesResponse{}

	schedules := []SeppoDB.Schedule{}

	query := s.databaseService.GetDb().Table("schedules")

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}
	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	}

	query.Count(&res.MaxSchedules)
	query.Find(&schedules)

	for i := 0; i < len(schedules); i++ {
		res.Schedules = append(res.Schedules, NewScheduleToServiceType(&schedules[i]))
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchScheduleById(
	ctx context.Context,
	in *SeppoService.FetchScheduleByIdRequest,
) (
	*SeppoService.FetchScheduleByIdResponse,
	error,
) {
	res := &SeppoService.FetchScheduleByIdResponse{}

	schedules := []SeppoDB.Schedule{}

	s.databaseService.GetDb().Where("schedules.id in (?)", in.ScheduleIds).Find(&schedules)

	for i := 0; i < len(in.ScheduleIds); i++ {
		found := false
		for j := 0; j < len(schedules); j++ {
			if in.ScheduleIds[i] == schedules[j].ID {
				found = true
				res.Schedules = append(res.Schedules, NewScheduleToServiceType(&schedules[j]))
			}
		}
		if found == false {
			res.Schedules = append(res.Schedules, &SeppoService.Schedule{
				Id: 0,
			})
		}
	}

	return res, nil
}
