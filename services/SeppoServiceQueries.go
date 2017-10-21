package services

import (
	"github.com/koodinikkarit/seppo/db"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
	"golang.org/x/net/context"
)

func createQuestionMarks(n int) string {
	str := ""
	for i := 0; i < n; i++ {
		str += "?, "
	}
	return str
}

func (s *SeppoServiceServer) ListenForChangedEwSong(
	in *SeppoService.ListenForChangedEwSongRequest,
	stream SeppoService.Seppo_ListenForChangedEwSongServer,
) error {

	return nil
}

func (s *SeppoServiceServer) FetchSongDatabases(
	ctx context.Context,
	in *SeppoService.FetchSongDatabasesRequest,
) (
	*SeppoService.SongDatabasesConnection,
	error,
) {
	res := &SeppoService.SongDatabasesConnection{}
	return res, nil
}

func (s *SeppoServiceServer) FetchSongDatabaseById(
	ctx context.Context,
	in *SeppoService.FetchSongDatabaseByIdRequest,
) (
	*SeppoService.FetchSongDatabaseByIdResponse,
	error,
) {
	res := &SeppoService.FetchSongDatabaseByIdResponse{}
	return res, nil
}

func (s *SeppoServiceServer) FetchEwDatabases(
	ctx context.Context,
	in *SeppoService.FetchEwDatabasesRequest,
) (
	*SeppoService.EwDatabasesConnection,
	error,
) {
	res := &SeppoService.EwDatabasesConnection{}
	return res, nil
}

func (s *SeppoServiceServer) FetchEwDatabaseById(
	ctx context.Context,
	in *SeppoService.FetchEwDatabaseByIdRequest,
) (
	*SeppoService.FetchEwDatabaseByIdResponse,
	error,
) {
	res := &SeppoService.FetchEwDatabaseByIdResponse{}
	return res, nil
}

func (s *SeppoServiceServer) FetchVariationsBySongDatabaseId(
	ctx context.Context,
	in *SeppoService.FetchVariationsBySongDatabaseIdRequest,
) (
	*SeppoService.FetchVariationsBySongDatabaseIdResponse,
	error,
) {
	res := &SeppoService.FetchVariationsBySongDatabaseIdResponse{}
	return res, nil
}

func (s *SeppoServiceServer) FetchVariationTextByVariationId(
	ctx context.Context,
	in *SeppoService.FetchVariationTextByVariationIdRequest,
) (
	*SeppoService.FetchVariationTextByVariationIdResponse,
	error,
) {
	res := &SeppoService.FetchVariationTextByVariationIdResponse{}
	return res, nil
}

func (s *SeppoServiceServer) SearchTags(
	ctx context.Context,
	in *SeppoService.SearchTagsRequest,
) (
	*SeppoService.TagsConnection,
	error,
) {
	res := &SeppoService.TagsConnection{}
	newDb := s.getDB()
	defer newDb.Close()

	query := newDb.Select("tags.id, tags.name").Table("tags")

	if in.SongDatabaseId > 0 {
		query = query.Joins("song_database_tags on tags.id = song_databases_tags.tag_id").
			Where("song_database_tags.song_database_id = ?", in.SongDatabaseId)
	}

	if in.VariationId > 0 {
		query = query.Joins("tag_variations on tags.id = tag_variations.tag_id").
			Where("tag_variations.variation_id = ?", in.VariationId)
	}

	if in.SearchWord != "" {
		query = query.Where("tags.name LIKE ?", "%"+in.SearchWord+"%")
	}

	query.Count(&res.MaxTags)

	query = query.Limit(5000)

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}

	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	}

	tags := []db.Tag{}

	query.Find(&tags)

	for i := 0; i < len(tags); i++ {
		res.Tags = append(
			res.Tags,
			&SeppoService.Tag{
				Id:   tags[i].ID,
				Name: tags[i].Name,
			},
		)
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchTagById(
	ctx context.Context,
	in *SeppoService.FetchTagByIdRequest,
) (
	*SeppoService.FetchTagByIdResponse,
	error,
) {
	res := &SeppoService.FetchTagByIdResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	tags := []db.Tag{}
	newDb.Where("id in (?)", in.TagIds).Find(&tags)

	for _, tagID := range in.TagIds {
		found := false
		for i := 0; i < len(tags); i++ {
			if tagID == tags[i].ID {
				found = true
				res.Tags = append(res.Tags, NewTag(&tags[i]))
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

func (s *SeppoServiceServer) SearchLanguages(
	ctx context.Context,
	in *SeppoService.SearchLanguagesRequest,
) (
	*SeppoService.LanguagesConnection,
	error,
) {
	res := &SeppoService.LanguagesConnection{}
	newDb := s.getDB()
	defer newDb.Close()

	languages := []db.Language{}

	query := newDb.Table("languages")

	query.Count(&res.MaxLanguages)

	if in.SearchWord != "" {
		query = query.Where("languages.name LIKE ?", "%"+in.SearchWord+"%")
	}

	query = query.Limit(5000)

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}
	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	}

	query.Find(&languages)

	for i := 0; i < len(languages); i++ {
		res.Languages = append(
			res.Languages,
			NewLanguage(&languages[i]),
		)
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchLanguageById(
	ctx context.Context,
	in *SeppoService.FetchLanguageByIdRequest,
) (
	*SeppoService.FetchLanguageByIdResponse,
	error,
) {
	res := &SeppoService.FetchLanguageByIdResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	languages := []db.Language{}

	newDb.Where("id in (?)", in.LanguageIds).Find(&languages)

	for _, languageID := range in.LanguageIds {
		found := false
		for i := 0; i < len(languages); i++ {
			if languageID == languages[i].ID {
				found = true
				res.Languages = append(res.Languages, NewLanguage(&languages[i]))
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
	return res, nil
}

func (s *SeppoServiceServer) FetchTagSongDatabases(
	ctx context.Context,
	in *SeppoService.FetchTagSongDatabasesRequest,
) (
	*SeppoService.FetchTagSongDatabasesResponse,
	error,
) {
	res := &SeppoService.FetchTagSongDatabasesResponse{}
	return res, nil
}

func (s *SeppoServiceServer) FetchSongDatabaseTags(
	ctx context.Context,
	in *SeppoService.FetchSongDatabaseTagsRequest,
) (
	*SeppoService.FetchSongDatabaseTagsResponse,
	error,
) {
	res := &SeppoService.FetchSongDatabaseTagsResponse{}
	return res, nil
}

func (s *SeppoServiceServer) FetchLanguageVariations(
	ctxt context.Context,
	in *SeppoService.FetchLanguageVariationsRequest,
) (
	*SeppoService.FetchLanguageVariationsResponse,
	error,
) {
	res := &SeppoService.FetchLanguageVariationsResponse{}
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
	return res, nil
}

func (s *SeppoServiceServer) SearchLogs(
	ctx context.Context,
	in *SeppoService.SearchLogsRequest,
) (
	*SeppoService.SearchLogsResponse,
	error,
) {
	res := &SeppoService.SearchLogsResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	query := newDb.Table("logs")

	if in.SearchWord != "" {
		query = query.Where("logs.message LIKE ?", "%"+in.SearchWord+"%")
	}

	if in.StartDate > 0 {
		//d := time.Unix(in.StartDate, 0)
		query = query.Where("logs.message_date > ?", in.StartDate)
	}

	if in.EndDate > 0 {
		//d := time.Unix(in.EndDate, 0)
		query = query.Where("logs.message_date < ?", in.EndDate)
	}

	query.Count(&res.MaxLogs)

	query = query.Limit(5000).Order("message_date desc")

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}

	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	}

	logs := []db.Log{}

	query.Find(&logs)

	for i := 0; i < len(logs); i++ {
		res.Logs = append(
			res.Logs,
			NewLog(&logs[i]),
		)
	}

	return res, nil
}
