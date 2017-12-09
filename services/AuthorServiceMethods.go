package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/generators"
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
	newDB := s.getGormDB()
	defer newDB.Close()

	author := db.Author{
		Name: in.Name,
	}
	newDB.Create(&author)

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
	newDB := s.getGormDB()
	defer newDB.Close()

	var author db.Author
	newDB.First(&author, in.AuthorId)

	if author.ID == 0 {
		res.Success = false
		return res, nil
	}
	if in.Name != "" {
		author.Name = in.Name
	}

	newDB.Save(&author)
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
	newDB := s.getGormDB()
	defer newDB.Close()

	var author db.Author
	newDB.First(&author, in.AuthorId)

	if author.ID == 0 {
		res.Success = false
		return res, nil
	}

	newDB.Delete(&author)
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
	newDB := s.getGormDB()
	defer newDB.Close()

	query := newDB.Table("authors")

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
	newDB := s.getGormDB()
	defer newDB.Close()

	var authors []db.Author
	newDB.Where("id in (?)", in.AuthorIds).
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
