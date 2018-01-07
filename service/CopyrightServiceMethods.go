package service

import (
	"golang.org/x/net/context"

	SeppoService "github.com/koodinikkarit/go-clientlibs/seppo"
	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/models"
)

func (s *SeppoServiceServer) CreateCopyright(
	ctx context.Context,
	in *SeppoService.CreateCopyrightRequest,
) (
	*SeppoService.CreateCopyrightResponse,
	error,
) {
	res := &SeppoService.CreateCopyrightResponse{}
	db := s.getDB()

	copyright := models.Copyright{
		Name: in.Name,
	}

	db.Create(&copyright)

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
	db := s.getDB()

	var copyright models.Copyright
	db.First(&copyright, in.CopyrightId)

	if copyright.ID == 0 {
		res.Success = true
		return res, nil
	}

	if in.Name != "" {
		copyright.Name = in.Name
	}
	db.Save(&copyright)
	res.Success = true
	res.Copyright = generators.NewCopyright(&copyright)

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
	db := s.getDB()

	var copyright models.Copyright
	db.First(&copyright, in.CopyrightId)

	if copyright.ID == 0 {
		res.Success = false
		return res, nil
	}

	db.Delete(&copyright)
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
	db := s.getDB()

	query := db.Table("copyrights")

	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	} else {
		query = query.Limit(10000)
	}

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}

	var copyrights []models.Copyright
	query.Find(&copyrights)

	for _, copyright := range copyrights {
		res.Copyrights = append(
			res.Copyrights,
			generators.NewCopyright(&copyright),
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
	db := s.getDB()

	var copyrights []models.Copyright
	db.Where("id in (?)", in.CopyrightIds).
		Find(&copyrights)

	for _, copyrightID := range in.CopyrightIds {
		found := false
		for _, copyright := range copyrights {
			if copyrightID != copyright.ID {
				continue
			}
			found = true
			res.Copyrights = append(
				res.Copyrights,
				generators.NewCopyright(&copyright),
			)
			break
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
