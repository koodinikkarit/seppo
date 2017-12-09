package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/generators"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
)

func (s SeppoServiceServer) CreateLanguage(
	ctx context.Context,
	in *SeppoService.CreateLanguageRequest,
) (
	*SeppoService.CreateLanguageResponse,
	error,
) {
	res := &SeppoService.CreateLanguageResponse{}
	newDB := s.getGormDB()
	defer newDB.Close()

	language := db.Language{
		Name: in.Name,
	}

	newDB.Create(&language)

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
	newDB := s.getGormDB()
	defer newDB.Close()

	var language db.Language

	newDB.First(&language, in.LanguageId)

	if language.ID == 0 {
		res.Success = false
		return res, nil
	}

	if in.Name != "" {
		language.Name = in.Name
	}
	newDB.Save(&language)
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
	newDB := s.getGormDB()
	defer newDB.Close()

	var language db.Language

	newDB.Select("id").First(&language, in.LanguageId)

	if language.ID > 0 {
		res.Success = false
		return res, nil
	}

	newDB.Delete(&language)
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
	newDB := s.getGormDB()
	defer newDB.Close()

	languages := []db.Language{}

	query := newDB.Table("languages")

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
	newDB := s.getGormDB()
	defer newDB.Close()

	languages := []db.Language{}

	newDB.Where("id in (?)", in.LanguageIds).Find(&languages)

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
