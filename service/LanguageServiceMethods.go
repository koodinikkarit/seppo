package service

import (
	"golang.org/x/net/context"

	SeppoService "github.com/koodinikkarit/go-clientlibs/seppo"
	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/models"
)

func (s SeppoServiceServer) CreateLanguage(
	ctx context.Context,
	in *SeppoService.CreateLanguageRequest,
) (
	*SeppoService.CreateLanguageResponse,
	error,
) {
	res := &SeppoService.CreateLanguageResponse{}
	db := s.getDB()

	language := models.Language{
		Name: in.Name,
	}

	db.Create(&language)

	res.Language = generators.NewLanguage(&language)

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
	db := s.getDB()

	var language models.Language

	db.First(&language, in.LanguageId)

	if language.ID == 0 {
		res.Success = false
		return res, nil
	}

	if in.Name != "" {
		language.Name = in.Name
	}
	db.Save(&language)
	res.Language = generators.NewLanguage(&language)
	res.Success = true

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
	db := s.getDB()

	var language models.Language

	db.Select("id").First(&language, in.LanguageId)

	if language.ID > 0 {
		res.Success = false
		return res, nil
	}

	db.Delete(&language)
	res.Success = true

	return res, nil
}

func (s *SeppoServiceServer) SearchLanguages(
	ctx context.Context,
	in *SeppoService.SearchLanguagesRequest,
) (
	*SeppoService.SearchLanguagesResponse,
	error,
) {
	res := &SeppoService.SearchLanguagesResponse{}
	db := s.getDB()

	languages := []models.Language{}

	query := db.Table("languages")

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

	for _, language := range languages {
		res.Languages = append(
			res.Languages,
			generators.NewLanguage(&language),
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
	db := s.getDB()

	languages := []models.Language{}

	db.Where("id in (?)", in.LanguageIds).Find(&languages)

	for _, languageID := range in.LanguageIds {
		found := false
		for _, language := range languages {
			if languageID != language.ID {
				continue
			}
			found = true
			res.Languages = append(
				res.Languages,
				generators.NewLanguage(&language),
			)
			break
		}
		if found == false {
			res.Languages = append(
				res.Languages,
				&SeppoService.Language{},
			)
		}
	}

	return res, nil
}
