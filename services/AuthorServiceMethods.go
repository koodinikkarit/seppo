package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
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

	author := db.Author{
		Name: in.Name,
	}

	newDb.Create(&author)

	res.Author = NewAuthor(&author)

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

	var author db.Author
	newDb.First(&author, in.AuthorId)
	if author.ID > 0 {
		if in.Name != "" {
			author.Name = in.Name
		}

		res.Success = true
		newDb.Save(&author)
	}

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

	var author db.Author
	newDb.First(&author, in.AuthorId)
	if author.ID > 0 {
		newDb.Delete(&author)
		res.Success = true
	}

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

	query := newDb.Table("authors")

	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	} else {
		query = query.Limit(5000)
	}

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}

	var authors []db.Author
	query.Find(&authors)

	for _, author := range authors {
		res.Authors = append(
			res.Authors,
			NewAuthor(&author),
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

	var authors []db.Author
	newDb.Where("id in (?)", in.AuthorIds).
		Find(&authors)

	for _, authorID := range in.AuthorIds {
		found := false
		for _, author := range authors {
			if authorID == author.ID {
				found = true
				res.Authors = append(
					res.Authors,
					NewAuthor(&author),
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
