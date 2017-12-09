package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/managers"
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
	newDB := s.getGormDB()
	defer newDB.Close()

	newSongdatabase := db.SongDatabase{
		Name: in.Name,
	}

	newDB.Create(&newSongdatabase)

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
	tx := s.getGormDB().Begin()
	defer tx.Close()

	var songDatabase db.SongDatabase
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
		var songDatabaseTags []db.SongDatabaseTag
		for _, tagID := range in.AddTagIds {
			songDatabaseTags = append(
				songDatabaseTags,
				db.SongDatabaseTag{
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
		tx.Where("tag_id in (?)", in.RemoveTagIds).Delete(&db.SongDatabaseTag{})
	}

	res.SongDatabase = generators.NewSongDatabase(&songDatabase)

	tx.Commit()

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
	newDB := s.getGormDB()
	defer newDB.Close()

	var songDatabase db.SongDatabase
	newDB.First(&songDatabase, in.SongDatabaseId)

	if songDatabase.ID == 0 {
		res.Success = false
		return res, nil
	}

	res.Success = true
	newDB.Delete(&songDatabase)

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
	newDB := s.getGormDB()
	defer newDB.Close()

	songDatabases := []db.SongDatabase{}

	query := newDB.Table("song_databases")

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
	newDB := s.getGormDB()
	defer newDB.Close()

	var songDatabases []db.SongDatabase
	newDB.Where("id in (?)", in.SongDatabaseIds).
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
