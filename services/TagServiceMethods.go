package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/generators"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
)

func (s SeppoServiceServer) CreateTag(
	ctx context.Context,
	in *SeppoService.CreateTagRequest,
) (
	*SeppoService.CreateTagResponse,
	error,
) {
	res := &SeppoService.CreateTagResponse{}
	newDB := s.getGormDB()
	defer newDB.Close()

	tag := db.Tag{
		Name: in.Name,
	}

	newDB.Create(&tag)

	res.Tag = &SeppoService.Tag{
		Id:   tag.ID,
		Name: tag.Name,
	}

	return res, nil
}

func (s SeppoServiceServer) UpdateTag(
	ctx context.Context,
	in *SeppoService.UpdateTagRequest,
) (
	*SeppoService.UpdateTagResponse,
	error,
) {
	res := &SeppoService.UpdateTagResponse{}
	newDb := s.getGormDB()
	defer newDb.Close()

	tag := db.Tag{}

	newDb.First(&tag, in.TagId)

	if tag.ID == 0 {
		res.Success = false
		return res, nil
	}

	if in.Name != "" {
		tag.Name = in.Name
	}

	newDb.Save(&tag)
	res.Tag = &SeppoService.Tag{
		Id:   tag.ID,
		Name: tag.Name,
	}
	res.Success = true

	return res, nil
}

func (s SeppoServiceServer) RemoveTag(
	ctx context.Context,
	in *SeppoService.RemoveTagRequest,
) (
	*SeppoService.RemoveTagResponse,
	error,
) {
	res := &SeppoService.RemoveTagResponse{}
	newDb := s.getGormDB()
	defer newDb.Close()

	tag := db.Tag{}

	newDb.Select("id").First(&tag, in.TagId)

	if tag.ID == 0 {
		res.Success = false
		return res, nil
	}

	newDb.Delete(&tag)
	res.Success = true

	return res, nil
}

func (s *SeppoServiceServer) SearchTags(
	ctx context.Context,
	in *SeppoService.SearchTagsRequest,
) (
	*SeppoService.SearchTagsResponse,
	error,
) {
	res := &SeppoService.SearchTagsResponse{}
	newDB := s.getGormDB()
	defer newDB.Close()

	query := newDB.Select("tags.id, tags.name").Table("tags")

	if in.SongDatabaseId > 0 {
		query = query.Joins("left join song_database_tags on song_database_tags.tag_id = tags.id").
			Where("song_database_tags.song_database_id = ?", in.SongDatabaseId)
	}

	if in.VariationId > 0 {
		query = query.Joins("left join tag_variations on tags.id = tag_variations.tag_id").
			Where("tag_variations.variation_id = ?", in.VariationId)
	}

	if in.SearchWord != "" {
		query = query.Where("tags.name LIKE ?", "%"+in.SearchWord+"%")
	}

	query.Count(&res.MaxTags)

	query = query.Limit(5000)

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}

	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	}

	tags := []db.Tag{}

	query.Find(&tags)

	for _, tag := range tags {
		res.Tags = append(
			res.Tags,
			generators.NewTag(&tag),
		)
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchTagById(
	ctx context.Context,
	in *SeppoService.FetchTagByIdRequest,
) (
	*SeppoService.FetchTagByIdResponse,
	error,
) {
	res := &SeppoService.FetchTagByIdResponse{}
	newDB := s.getGormDB()
	defer newDB.Close()

	var tags []db.Tag
	newDB.Where("id in (?)", in.TagIds).
		Find(&tags)

	for _, tagID := range in.TagIds {
		found := false
		for _, tag := range tags {
			if tagID != tag.ID {
				continue
			}
			found = true
			res.Tags = append(
				res.Tags,
				generators.NewTag(&tag),
			)
			break
		}
		if found == false {
			res.Tags = append(res.Tags, &SeppoService.Tag{})
		}
	}

	return res, nil
}
