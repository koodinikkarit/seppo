package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/models"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
	"github.com/volatiletech/sqlboiler/queries/qm"
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

	language := models.Language{
		Name: in.Name,
	}

	language.Insert(newDb)
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
	newDb := s.getDB()
	defer newDb.Close()

	language, _ := models.FindLanguage(
		newDb,
		in.LanguageId,
	)

	if language == nil {
		res.Success = false
		return res, nil
	}

	if in.Name != "" {
		language.Name = in.Name
	}
	language.Update(newDb)
	res.Language = generators.NewLanguage(language)
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
	newDb := s.getDB()
	defer newDb.Close()

	language, _ := models.FindLanguage(
		newDb,
		in.LanguageId,
	)

	if language == nil {
		res.Success = false
		return res, nil
	}

	language.Delete(newDb)
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
	newDb := s.getDB()
	defer newDb.Close()

	var queryMods []qm.QueryMod

	c, _ := models.Languages(newDb).Count()
	res.MaxLanguages = uint64(c)

	if in.SearchWord != "" {
		queryMods = append(
			queryMods,
			qm.Where("languages.name LIKE ?", "%"+in.SearchWord+"%"),
		)
	}

	if in.Offset > 0 {
		queryMods = append(
			queryMods,
			qm.Offset(int(in.Offset)),
		)
	} else {
		queryMods = append(
			queryMods,
			qm.Offset(10000),
		)
	}
	if in.Limit > 0 {
		queryMods = append(
			queryMods,
			qm.Limit(int(in.Limit)),
		)
	}

	languages, _ := models.Languages(
		newDb,
		queryMods...,
	).All()

	for _, language := range languages {
		res.Languages = append(
			res.Languages,
			generators.NewLanguage(language),
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

	languages, _ := models.Languages(
		newDb,
		qm.WhereIn("id in ?", in.LanguageIds),
	).All()

	for _, languageID := range in.LanguageIds {
		found := false
		for _, language := range languages {
			if languageID == language.ID {
				found = true
				res.Languages = append(
					res.Languages,
					generators.NewLanguage(language),
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
