package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/models"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (s *SeppoServiceServer) CreateCopyright(
	ctx context.Context,
	in *SeppoService.CreateCopyrightRequest,
) (
	*SeppoService.CreateCopyrightResponse,
	error,
) {
	res := &SeppoService.CreateCopyrightResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	copyright := models.Copyright{
		Name: in.Name,
	}
	copyright.Insert(newDb)

	res.Copyright = generators.NewCopyright(&copyright)

	return res, nil
}

func (s *SeppoServiceServer) UpdateCopyright(
	ctx context.Context,
	in *SeppoService.UpdateCopyrightRequest,
) (
	*SeppoService.UpdateCopyrightResponse,
	error,
) {
	res := &SeppoService.UpdateCopyrightResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	copyright, _ := models.FindCopyright(
		newDb,
		in.CopyrightId,
	)

	if copyright == nil {
		res.Success = true
		return res, nil
	}

	copyright.Name = in.Name
	copyright.Update(newDb)
	res.Success = true
	res.Copyright = generators.NewCopyright(copyright)

	return res, nil
}

func (s *SeppoServiceServer) RemoveCopyright(
	ctx context.Context,
	in *SeppoService.RemoveCopyrightRequest,
) (
	*SeppoService.RemoveCopyrightResponse,
	error,
) {
	res := &SeppoService.RemoveCopyrightResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	copyright, _ := models.FindCopyright(
		newDb,
		in.CopyrightId,
	)

	if copyright == nil {
		res.Success = false
		return res, nil
	}

	copyright.Delete(newDb)
	res.Success = true

	return res, nil
}

func (s *SeppoServiceServer) SearchCopyrights(
	ctx context.Context,
	in *SeppoService.SearchCopyrightsRequest,
) (
	*SeppoService.SearchCopyrightsResponse,
	error,
) {
	res := &SeppoService.SearchCopyrightsResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	var queryMods []qm.QueryMod

	queryMods = append(
		queryMods,
		qm.From("copyrights"),
	)

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

	if in.Offset > 0 {
		queryMods = append(
			queryMods,
			qm.Offset(int(in.Offset)),
		)
	}

	copyrights, _ := models.Copyrights(
		newDb,
		queryMods...,
	).All()

	for _, copyright := range copyrights {
		res.Copyrights = append(
			res.Copyrights,
			generators.NewCopyright(copyright),
		)
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchCopyrightById(
	ctx context.Context,
	in *SeppoService.FetchCopyrightByIdRequest,
) (
	*SeppoService.FetchCopyrightByIdResponse,
	error,
) {
	res := &SeppoService.FetchCopyrightByIdResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	copyrights, _ := models.Copyrights(
		newDb,
		qm.WhereIn("id in ?", in.CopyrightIds),
	).All()

	for _, copyrightID := range in.CopyrightIds {
		found := false
		for _, copyright := range copyrights {
			if copyrightID == copyright.ID {
				found = true
				res.Copyrights = append(
					res.Copyrights,
					generators.NewCopyright(copyright),
				)
				break
			}
		}
		if found == false {
			res.Copyrights = append(
				res.Copyrights,
				&SeppoService.Copyright{},
			)
		}
	}

	return res, nil
}
