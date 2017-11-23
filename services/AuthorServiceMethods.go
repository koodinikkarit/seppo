package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/models"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (s *SeppoServiceServer) CreateAuthor(
	ctx context.Context,
	in *SeppoService.CreateAuthorRequest,
) (
	*SeppoService.CreateAuthorResponse,
	error,
) {
	res := &SeppoService.CreateAuthorResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	author := models.Author{
		Name: in.Name,
	}
	author.Insert(newDb)

	res.Author = generators.NewAuthor(&author)

	return res, nil
}

func (s *SeppoServiceServer) UpdateAuthor(
	ctx context.Context,
	in *SeppoService.UpdateAuthorRequest,
) (
	*SeppoService.UpdateAuthorResponse,
	error,
) {
	res := &SeppoService.UpdateAuthorResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	author, _ := models.FindAuthor(newDb, in.AuthorId)
	if author == nil {
		res.Success = false
		return res, nil
	}
	author.Name = in.Name
	author.Update(newDb)
	res.Success = true
	res.Author = generators.NewAuthor(author)

	return res, nil
}

func (s *SeppoServiceServer) RemoveAuthor(
	ctx context.Context,
	in *SeppoService.RemoveAuthorRequest,
) (
	*SeppoService.RemoveAuthorResponse,
	error,
) {
	res := &SeppoService.RemoveAuthorResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	author, _ := models.FindAuthor(newDb, in.AuthorId)
	if author == nil {
		res.Success = false
		return res, nil
	}
	author.Delete(newDb)
	res.Success = true

	return res, nil
}

func (s *SeppoServiceServer) SearchAuthors(
	ctx context.Context,
	in *SeppoService.SearchAuthorsRequest,
) (
	*SeppoService.SearchAuthorsResponse,
	error,
) {
	res := &SeppoService.SearchAuthorsResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	var queryes []qm.QueryMod

	queryes = append(
		queryes,
		qm.From("authors"),
	)

	if in.Limit > 0 {
		queryes = append(
			queryes,
			qm.Limit(int(in.Limit)),
		)
	} else {
		queryes = append(
			queryes,
			qm.Limit(10000),
		)
	}

	if in.Offset > 0 {
		queryes = append(
			queryes,
			qm.Offset(int(in.Offset)),
		)
	}

	authors, _ := models.Authors(
		newDb,
		queryes...,
	).All()

	for _, author := range authors {
		res.Authors = append(
			res.Authors,
			generators.NewAuthor(author),
		)
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchAuthorById(
	ctx context.Context,
	in *SeppoService.FetchAuthorByIdRequest,
) (
	*SeppoService.FetchAuthorByIdResponse,
	error,
) {
	res := &SeppoService.FetchAuthorByIdResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	authors, _ := models.Authors(
		newDb,
		qm.WhereIn("id in ?", in.AuthorIds),
	).All()

	for _, authorID := range in.AuthorIds {
		found := false
		for _, author := range authors {
			if authorID == author.ID {
				found = true
				res.Authors = append(
					res.Authors,
					generators.NewAuthor(author),
				)
				break
			}
		}
		if found == false {
			res.Authors = append(
				res.Authors,
				&SeppoService.Author{},
			)
		}
	}

	return res, nil
}
