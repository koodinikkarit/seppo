package service

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/models"
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
	db := s.getDB()

	author := models.Author{
		Name: in.Name,
	}
	db.Create(&author)

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
	db := s.getDB()

	var author models.Author
	db.First(&author, in.AuthorId)

	if author.ID == 0 {
		res.Success = false
		return res, nil
	}
	if in.Name != "" {
		author.Name = in.Name
	}

	db.Save(&author)
	res.Success = true
	res.Author = generators.NewAuthor(&author)

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
	db := s.getDB()

	var author models.Author
	db.First(&author, in.AuthorId)

	if author.ID == 0 {
		res.Success = false
		return res, nil
	}

	db.Delete(&author)
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
	db := s.getDB()

	query := db.Table("authors")

	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	} else {
		query = query.Limit(5000)
	}

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}

	var authors []models.Author
	query.Find(&authors)

	for _, author := range authors {
		res.Authors = append(
			res.Authors,
			generators.NewAuthor(&author),
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
	db := s.getDB()

	var authors []models.Author
	db.Where("id in (?)", in.AuthorIds).
		Find(&authors)

	for _, authorID := range in.AuthorIds {
		found := false
		for _, author := range authors {
			if authorID != author.ID {
				continue
			}
			found = true
			res.Authors = append(
				res.Authors,
				generators.NewAuthor(&author),
			)
			break
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
