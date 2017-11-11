package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
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
	newDb := s.getDB()
	defer newDb.Close()

	copyright := db.Copyright{
		Name: in.Name,
	}

	newDb.Create(&copyright)

	res.Copyright = NewCopyright(&copyright)

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

	var copyright db.Copyright
	newDb.First(&copyright, in.CopyrightId)
	if copyright.ID > 0 {
		if in.Name != "" {
			copyright.Name = in.Name
		}

		res.Success = true
		newDb.Save(&copyright)
	}

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

	var copyright db.Copyright
	newDb.First(&copyright, in.CopyrightId)
	if copyright.ID > 0 {
		newDb.Delete(&copyright)
		res.Success = true
	}

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

	query := newDb.Table("copyrights")

	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	} else {
		query = query.Limit(5000)
	}

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}

	var copyrights []db.Copyright
	query.Find(&copyrights)

	for _, copyright := range copyrights {
		res.Copyrights = append(
			res.Copyrights,
			NewCopyright(&copyright),
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

	var copyrights []db.Copyright
	newDb.Where("id in (?)", in.CopyrightIds).
		Find(&copyrights)

	for _, copyrightID := range in.CopyrightIds {
		found := false
		for _, copyright := range copyrights {
			if copyrightID == copyright.ID {
				found = true
				res.Copyrights = append(
					res.Copyrights,
					NewCopyright(&copyright),
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
