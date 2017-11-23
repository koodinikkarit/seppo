package services

import (
	"time"

	null "gopkg.in/volatiletech/null.v6"

	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/models"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (s SeppoServiceServer) CreateTag(
	ctx context.Context,
	in *SeppoService.CreateTagRequest,
) (
	*SeppoService.CreateTagResponse,
	error,
) {
	res := &SeppoService.CreateTagResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	tag := models.Tag{
		Name: in.Name,
	}
	tag.Insert(newDb)

	res.Tag = generators.NewTag(&tag)

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
	newDb := s.getDB()
	defer newDb.Close()

	tag, _ := models.FindTag(
		newDb,
		in.TagId,
	)

	if tag == nil {
		res.Success = false
		return res, nil
	}

	if in.Name != "" {
		tag.Name = in.Name
	}
	tag.Update(newDb)
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
	newDb := s.getDB()
	defer newDb.Close()

	tag, _ := models.FindTag(
		newDb,
		in.TagId,
	)

	if tag == nil {
		res.Success = false
		return res, nil
	}

	tag.DeletedAt = null.NewTime(time.Now(), true)
	tag.Update(newDb)
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
	newDb := s.getDB()
	defer newDb.Close()

	var queryMods []qm.QueryMod

	if in.SongDatabaseId > 0 {
		queryMods = append(
			queryMods,
			qm.InnerJoin("song_database_tags sdt on sdt.tag_id = tags.id"),
			qm.Where("sdt.song_database_id = ?", in.SongDatabaseId),
		)
	}

	if in.VariationId > 0 {
		queryMods = append(
			queryMods,
			qm.InnerJoin("tag_variations tv on tags.id = tv.tag_id"),
			qm.Where("tv.variation_id = ?", in.VariationId),
		)
	}

	c, _ := models.Tags(
		newDb,
		queryMods...,
	).Count()
	res.MaxTags = uint64(c)

	if in.SearchWord != "" {
		queryMods = append(
			queryMods,
			qm.Where("tags.name LIKE ?", "%"+in.SearchWord+"%"),
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
			qm.Limit(int(in.Offset)),
		)
	}

	tags, _ := models.Tags(
		newDb,
		queryMods...,
	).All()

	for _, tag := range tags {
		res.Tags = append(
			res.Tags,
			generators.NewTag(tag),
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
	newDb := s.getDB()
	defer newDb.Close()

	tags, _ := models.Tags(
		newDb,
		qm.WhereIn("id in ?", in.TagIds),
	).All()

	for _, tagID := range in.TagIds {
		found := false
		for _, tag := range tags {
			if tagID != tag.ID {
				continue
			}
			found = true
			res.Tags = append(
				res.Tags,
				generators.NewTag(tag),
			)
		}
		if found == false {
			res.Tags = append(res.Tags, &SeppoService.Tag{})
		}
	}

	return res, nil
}
