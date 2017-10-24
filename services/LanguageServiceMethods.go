package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
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

func (s *SeppoServiceServer) SearchLanguages(
	ctx context.Context,
	in *SeppoService.SearchLanguagesRequest,
) (
	*SeppoService.SearchLanguagesResponse,
	error,
) {
	res := &SeppoService.SearchLanguagesResponse{}
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
				res.Languages = append(
					res.Languages,
					NewLanguage(&languages[i]),
				)
			}
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
