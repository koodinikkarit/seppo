package services

import (
	"strconv"

	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/logs"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
)

func (s *SeppoServiceServer) CreateSongDatabase(
	ctx context.Context,
	in *SeppoService.CreateSongDatabaseRequest,
) (
	*SeppoService.CreateSongDatabaseResponse,
	error,
) {
	res := &SeppoService.CreateSongDatabaseResponse{}
	return res, nil
}

func (s *SeppoServiceServer) UpdateSongDatabase(
	ctx context.Context,
	in *SeppoService.UpdateSongDatabaseRequest,
) (
	*SeppoService.UpdateSongDatabaseResponse,
	error,
) {
	res := &SeppoService.UpdateSongDatabaseResponse{}
	return res, nil
}

func (s *SeppoServiceServer) RemoveSongDatabase(
	ctx context.Context,
	in *SeppoService.RemoveSongDatabaseRequest,
) (
	*SeppoService.RemoveSongDatabaseResponse,
	error,
) {
	res := &SeppoService.RemoveSongDatabaseResponse{}
	return res, nil
}

func (s *SeppoServiceServer) CreateEwDatabase(
	ctx context.Context,
	in *SeppoService.CreateEwDatabaseRequest,
) (
	*SeppoService.CreateEwDatabaseResponse,
	error,
) {
	res := &SeppoService.CreateEwDatabaseResponse{}
	return res, nil
}

func (s *SeppoServiceServer) UpdateEwDatabase(
	ctx context.Context,
	in *SeppoService.UpdateEwDatabaseRequest,
) (
	*SeppoService.UpdateEwDatabaseResponse,
	error,
) {
	res := &SeppoService.UpdateEwDatabaseResponse{}
	return res, nil
}

func (s *SeppoServiceServer) RemoveEwDatabase(
	ctx context.Context,
	in *SeppoService.RemoveEwDatabaseRequest,
) (
	*SeppoService.RemoveEwDatabaseResponse,
	error,
) {
	res := &SeppoService.RemoveEwDatabaseResponse{}
	return res, nil
}

func (s *SeppoServiceServer) AddVariationToSongDatabase(
	ctx context.Context,
	in *SeppoService.AddVariationToSongDatabaseRequest,
) (
	*SeppoService.AddVariationToSongDatabaseResponse,
	error,
) {
	res := &SeppoService.AddVariationToSongDatabaseResponse{}
	return res, nil
}

func (s *SeppoServiceServer) RemoveVariationFromSongDatabase(
	ctx context.Context,
	in *SeppoService.RemoveVariationFromSongDatabaseRequest,
) (
	*SeppoService.RemoveVariationFromSongDatabaseResponse,
	error,
) {
	res := &SeppoService.RemoveVariationFromSongDatabaseResponse{}
	return res, nil
}

func (s SeppoServiceServer) CreateTag(
	ctx context.Context,
	in *SeppoService.CreateTagRequest,
) (
	*SeppoService.CreateTagResponse,
	error,
) {
	res := &SeppoService.CreateTagResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	tag := db.Tag{
		Name: in.Name,
	}

	newDb.Create(&tag)

	logs.InsertLog(
		newDb,
		1,
		"Created tag named "+in.Name,
	)

	res.Tag = &SeppoService.Tag{
		Id:   tag.ID,
		Name: tag.Name,
	}

	return res, nil
}

func (s SeppoServiceServer) UpdateTag(
	ctx context.Context,
	in *SeppoService.UpdateTagRequest,
) (
	*SeppoService.UpdateTagResponse,
	error,
) {
	res := &SeppoService.UpdateTagResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	tag := db.Tag{}

	newDb.First(&tag, in.TagId)

	if tag.ID > 0 {
		if in.Name != "" {
			tag.Name = in.Name
		}

		newDb.Save(&tag)
		res.Tag = &SeppoService.Tag{
			Id:   tag.ID,
			Name: tag.Name,
		}
		logs.InsertLog(newDb, 1, "Updated tag named "+tag.Name)
		res.Success = true
	} else {
		res.Success = false
	}

	return res, nil
}

func (s SeppoServiceServer) RemoveTag(
	ctx context.Context,
	in *SeppoService.RemoveTagRequest,
) (
	*SeppoService.RemoveTagResponse,
	error,
) {
	res := &SeppoService.RemoveTagResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	tag := db.Tag{}

	newDb.Select("id").First(&tag, in.TagId)

	if tag.ID > 0 {
		newDb.Delete(&tag)
		logs.InsertLog(newDb, 1, "Removed tag with id "+strconv.Itoa(int(in.TagId)))
		res.Success = true
	} else {
		res.Success = false
	}

	return res, nil
}

func (s SeppoServiceServer) CreateLanguage(
	ctx context.Context,
	in *SeppoService.CreateLanguageRequest,
) (
	*SeppoService.CreateLanguageResponse,
	error,
) {
	res := &SeppoService.CreateLanguageResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	language := db.Language{
		Name: in.Name,
	}

	newDb.Create(&language)

	res.Language = NewLanguage(&language)

	return res, nil
}

func (s SeppoServiceServer) UpdateLanguage(
	ctx context.Context,
	in *SeppoService.UpdateLanguageRequest,
) (
	*SeppoService.UpdateLanguageResponse,
	error,
) {
	res := &SeppoService.UpdateLanguageResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	var language db.Language

	newDb.First(&language, in.LanguageId)

	if language.ID > 0 {
		language.Name = in.Name
		newDb.Save(&language)
		res.Language = NewLanguage(&language)
		res.Success = true
	} else {
		res.Success = false
	}

	return res, nil
}

func (s SeppoServiceServer) RemoveLanguage(
	ctx context.Context,
	in *SeppoService.RemoveLanguageRequest,
) (
	*SeppoService.RemoveLanguageResponse,
	error,
) {
	res := &SeppoService.RemoveLanguageResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	var language db.Language

	newDb.Select("id").First(&language, in.LanguageId)

	if language.ID > 0 {
		newDb.Delete(&language)
		res.Success = true
	} else {
		res.Success = false
	}

	return res, nil
}

func (s SeppoServiceServer) AddTagToVariation(
	ctx context.Context,
	in *SeppoService.AddTagToVariationRequest,
) (
	*SeppoService.AddTagToVariationResponse,
	error,
) {
	res := &SeppoService.AddTagToVariationResponse{}
	return res, nil
}

func (s SeppoServiceServer) RemoveTagFromVariation(
	ctx context.Context,
	in *SeppoService.RemoveTagFromVariationRequest,
) (
	*SeppoService.RemoveTagFromVariationResponse,
	error,
) {
	res := &SeppoService.RemoveTagFromVariationResponse{}
	return res, nil
}

func (s SeppoServiceServer) AddTagToSongDatabase(
	ctx context.Context,
	in *SeppoService.AddTagToSongDatabaseRequest,
) (
	*SeppoService.AddTagToSongDatabaseResponse,
	error,
) {
	res := &SeppoService.AddTagToSongDatabaseResponse{}
	return res, nil
}

func (s SeppoServiceServer) RemoveTagFromSongDatabase(
	ctx context.Context,
	in *SeppoService.RemoveTagFromSongDatabaseRequest,
) (
	*SeppoService.RemoveTagFromSongDatabaseResponse,
	error,
) {
	res := &SeppoService.RemoveTagFromSongDatabaseResponse{}
	return res, nil
}

func (s SeppoServiceServer) CreateSchedule(
	ctx context.Context,
	in *SeppoService.CreateScheduleRequest,
) (
	*SeppoService.CreateScheduleResponse,
	error,
) {
	res := &SeppoService.CreateScheduleResponse{}
	return res, nil
}

func (s SeppoServiceServer) UpdateSchedule(
	ctx context.Context,
	in *SeppoService.UpdateScheduleRequest,
) (
	*SeppoService.UpdateScheduleResponse,
	error,
) {
	res := &SeppoService.UpdateScheduleResponse{}
	return res, nil
}

func (s SeppoServiceServer) RemoveSchedule(
	ctx context.Context,
	in *SeppoService.RemoveScheduleRequest,
) (
	*SeppoService.RemoveScheduleResponse,
	error,
) {
	res := &SeppoService.RemoveScheduleResponse{}
	return res, nil
}
