package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/generators"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
)

func (s *SeppoServiceServer) CreateCopyright(
	ctx context.Context,
	in *SeppoService.CreateCopyrightRequest,
) (
	*SeppoService.CreateCopyrightResponse,
	error,
) {
	res := &SeppoService.CreateCopyrightResponse{}
	newDB := s.getGormDB()
	defer newDB.Close()

	copyright := db.Copyright{
		Name: in.Name,
	}

	newDB.Create(&copyright)

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
	newDB := s.getGormDB()
	defer newDB.Close()

	var copyright db.Copyright
	newDB.First(&copyright, in.CopyrightId)

	if copyright.ID == 0 {
		res.Success = true
		return res, nil
	}

	if in.Name != "" {
		copyright.Name = in.Name
	}
	newDB.Save(&copyright)
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
	newDB := s.getGormDB()
	defer newDB.Close()

	var copyright db.Copyright
	newDB.First(&copyright, in.CopyrightId)

	if copyright.ID == 0 {
		res.Success = false
		return res, nil
	}

	newDB.Delete(&copyright)
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
	newDB := s.getGormDB()
	defer newDB.Close()

	query := newDB.Table("copyrights")

	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	} else {
		query = query.Limit(10000)
	}

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}

	var copyrights []db.Copyright
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
	newDB := s.getGormDB()
	defer newDB.Close()

	var copyrights []db.Copyright
	newDB.Where("id in (?)", in.CopyrightIds).
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
