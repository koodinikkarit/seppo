package service

import (
	"golang.org/x/net/context"

	SeppoService "github.com/koodinikkarit/go-clientlibs/seppo"
	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/managers"
	"github.com/koodinikkarit/seppo/matias"
	"github.com/koodinikkarit/seppo/models"
)

func (s *SeppoServiceServer) CreateSongDatabase(
	ctx context.Context,
	in *SeppoService.CreateSongDatabaseRequest,
) (
	*SeppoService.CreateSongDatabaseResponse,
	error,
) {
	res := &SeppoService.CreateSongDatabaseResponse{}
	db := s.getDB()

	newSongdatabase := models.SongDatabase{
		Name: in.Name,
	}

	db.Create(&newSongdatabase)

	s.pubSub.Pub(matias.CreatedSongDatabaseEvent{
		SongDatabase: newSongdatabase,
	}, "createdSongDatabase")

	res.SongDatabase = generators.NewSongDatabase(&newSongdatabase)

	return res, nil
}

func (s *SeppoServiceServer) UpdateSongDatabase(
	ctx context.Context,
	in *SeppoService.UpdateSongDatabaseRequest,
) (
	*SeppoService.UpdateSongDatabaseResponse,
	error,
) {
	res := &SeppoService.UpdateSongDatabaseResponse{}
	tx := s.getDB().Begin()

	var songDatabase models.SongDatabase
	tx.First(&songDatabase, in.SongDatabaseId)

	if songDatabase.ID == 0 {
		res.Success = false
		return res, nil
	}

	res.Success = true

	if in.Name != "" {
		songDatabase.Name = in.Name
	}
	tx.Save(&songDatabase)

	if len(in.AddTagIds) > 0 {
		var songDatabaseTags []models.SongDatabaseTag
		for _, tagID := range in.AddTagIds {
			songDatabaseTags = append(
				songDatabaseTags,
				models.SongDatabaseTag{
					TagID:          tagID,
					SongDatabaseID: songDatabase.ID,
				},
			)
		}
		managers.BatchCreateSongDatabaseTags(
			tx,
			songDatabaseTags,
		)
	}

	if len(in.RemoveTagIds) > 0 {
		tx.Where("tag_id in (?)", in.RemoveTagIds).Delete(&models.SongDatabaseTag{})
	}

	res.SongDatabase = generators.NewSongDatabase(&songDatabase)

	tx.Commit()

	s.pubSub.Pub(matias.UpdateSongDatabaseEvent{
		SongDatabase: songDatabase,
	}, "updatedSongDatabase")

	return res, nil
}

func (s *SeppoServiceServer) RemoveSongDatabase(
	ctx context.Context,
	in *SeppoService.RemoveSongDatabaseRequest,
) (
	*SeppoService.RemoveSongDatabaseResponse,
	error,
) {
	res := &SeppoService.RemoveSongDatabaseResponse{}
	db := s.getDB()

	var songDatabase models.SongDatabase
	db.First(&songDatabase, in.SongDatabaseId)

	if songDatabase.ID == 0 {
		res.Success = false
		return res, nil
	}

	res.Success = true
	db.Delete(&songDatabase)

	s.pubSub.Pub(matias.RemovedSongDatabaseEvent{
		SongDatabaseID: in.SongDatabaseId,
	})

	return res, nil
}

func (s *SeppoServiceServer) SearchSongDatabases(
	ctx context.Context,
	in *SeppoService.SearchSongDatabasesRequest,
) (
	*SeppoService.SearchSongDatabasesResponse,
	error,
) {
	res := &SeppoService.SearchSongDatabasesResponse{}
	db := s.getDB()

	songDatabases := []models.SongDatabase{}

	query := db.Table("song_databases")

	if in.VariationId > 0 {
		query = query.Joins("JOIN song_database_variations on song_database_variations.song_database_id = song_databases.id").
			Where("song_database_variations.variation_id = ?", in.VariationId).
			Where("song_database_variations.deleted_at is null")
	}

	query.Count(&res.MaxSongDatabases)

	if in.SearchWord != "" {
		query = query.Where("song_databases.name LIKE ?", "%"+in.SearchWord+"%")
	}

	query = query.Find(&songDatabases)

	for _, songDatabase := range songDatabases {
		res.SongDatabases = append(
			res.SongDatabases,
			generators.NewSongDatabase(&songDatabase),
		)
	}
	return res, nil
}

func (s *SeppoServiceServer) FetchSongDatabaseById(
	ctx context.Context,
	in *SeppoService.FetchSongDatabaseByIdRequest,
) (
	*SeppoService.FetchSongDatabaseByIdResponse,
	error,
) {
	res := &SeppoService.FetchSongDatabaseByIdResponse{}
	db := s.getDB()

	var songDatabases []models.SongDatabase
	db.Where("id in (?)", in.SongDatabaseIds).
		Find(&songDatabases)

	for _, songDatabaseID := range in.SongDatabaseIds {
		found := false
		for _, songDatabase := range songDatabases {
			if songDatabase.ID != songDatabaseID {
				continue
			}
			found = true
			res.SongDatabases = append(
				res.SongDatabases,
				generators.NewSongDatabase(&songDatabase),
			)
			break
		}
		if found == false {
			res.SongDatabases = append(
				res.SongDatabases,
				&SeppoService.SongDatabase{},
			)
		}
	}

	return res, nil
}
