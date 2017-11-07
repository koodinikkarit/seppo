package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
)

func (s *SeppoServiceServer) CreateSongDatabase(
	ctx context.Context,
	in *SeppoService.CreateSongDatabaseRequest,
) (
	*SeppoService.CreateSongDatabaseResponse,
	error,
) {
	res := &SeppoService.CreateSongDatabaseResponse{}
	tx := s.getDB().Begin()
	defer tx.Close()

	newSongdatabase := db.SongDatabase{
		Name: in.Name,
	}

	tx.Create(&newSongdatabase)

	tx.Commit()

	res.SongDatabase = NewSongDatabase(&newSongdatabase)

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
	defer tx.Close()

	var songDatabase db.SongDatabase
	tx.First(&songDatabase, in.SongDatabaseId)
	if songDatabase.ID > 0 {
		res.Success = true

		if in.Name != "" {
			songDatabase.Name = in.Name
		}

		for i := 0; i < len(in.AddTagIds); i++ {
			tx.Create(&db.SongDatabaseTag{
				TagID:          in.AddTagIds[i],
				SongDatabaseID: songDatabase.ID,
			})
		}
		if len(in.RemoveTagIds) > 0 {
			tx.Where("tag_id in (?)", in.RemoveTagIds).Delete(&db.SongDatabaseTag{})
		}

		tx.Save(&songDatabase)
		tx.Commit()
	} else {
		res.Success = false
	}

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
	tx := s.getDB().Begin()
	defer tx.Close()

	var songDatabase db.SongDatabase
	tx.First(&songDatabase, in.SongDatabaseId)

	if songDatabase.ID > 0 {
		res.Success = true
		tx.Delete(&songDatabase)
		tx.Commit()
	} else {
		res.Success = false
	}

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
	newDb := s.getDB()
	defer newDb.Close()

	songDatabases := []db.SongDatabase{}

	query := newDb.Table("song_databases")

	if in.VariationId > 0 {
		query = query.Joins("JOIN song_database_variations on song_database_variations.song_database_id = song_databases.id").
			Where("song_database_variations.variation_id = ?", in.VariationId)
	}

	query.Count(&res.MaxSongDatabases)

	if in.SearchWord != "" {
		query = query.Where("song_databases.name LIKE ?", "%"+in.SearchWord+"%")
	}

	query = query.Find(&songDatabases)

	for _, songDatabase := range songDatabases {
		res.SongDatabases = append(
			res.SongDatabases,
			NewSongDatabase(&songDatabase),
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
	newDb := s.getDB()
	defer newDb.Close()

	songDatabases := []db.SongDatabase{}
	newDb.Where("id in (?)", in.SongDatabaseIds).
		Find(&songDatabases)

	for _, songDatabaseId := range in.SongDatabaseIds {
		found := false
		for _, songDatabase := range songDatabases {
			if songDatabase.ID == songDatabaseId {
				found = true
				res.SongDatabases = append(
					res.SongDatabases,
					NewSongDatabase(&songDatabase),
				)
			}
		}
		if found == false {
			res.SongDatabases = append(
				res.SongDatabases,
				&SeppoService.SongDatabase{
					Id: 0,
				},
			)
		}
	}

	return res, nil
}
