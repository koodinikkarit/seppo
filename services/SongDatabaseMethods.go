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

func (s *SeppoServiceServer) CreateSongDatabase(
	ctx context.Context,
	in *SeppoService.CreateSongDatabaseRequest,
) (
	*SeppoService.CreateSongDatabaseResponse,
	error,
) {
	res := &SeppoService.CreateSongDatabaseResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	newSongDatabase := models.SongDatabase{
		Name: in.Name,
	}
	newSongDatabase.Insert(newDb)
	res.SongDatabase = generators.NewSongDatabase(&newSongDatabase)
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
	newDb := s.getDB()
	defer newDb.Close()
	tx, _ := newDb.Begin()

	songDatabase, _ := models.FindSongDatabase(
		tx,
		in.SongDatabaseId,
	)

	if songDatabase == nil {
		res.Success = false
		return res, nil
	}

	if in.Name != "" {
		songDatabase.Name = in.Name
	}
	songDatabase.Update(tx)

	for _, newTagID := range in.AddTagIds {
		songDatabaseTag := models.SongDatabaseTag{
			TagID: newTagID,
		}
		songDatabase.AddSongDatabaseTags(
			tx,
			true,
			&songDatabaseTag,
		)
	}

	if len(in.RemoveTagIds) > 0 {
		songDatabase.SongDatabaseTags(
			tx,
			qm.WhereIn("song_database_tags.tag_id = ?", in.RemoveTagIds),
		).DeleteAll()
	}

	tx.Commit()
	res.Success = true
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
	newDb := s.getDB()
	defer newDb.Close()

	songDatabase, _ := models.FindSongDatabase(
		newDb,
		in.SongDatabaseId,
	)

	if songDatabase == nil {
		res.Success = false
		return res, nil
	}

	songDatabase.DeletedAt = null.NewTime(time.Now(), true)
	songDatabase.Update(newDb)
	res.Success = true
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

	var queryMods []qm.QueryMod

	if in.VariationId > 0 {
		queryMods = append(
			queryMods,
			qm.InnerJoin("song_database_variations sdv on sdv.song_database_id = song_databases.id"),
			qm.Where("sdv.variation_id = ?", in.VariationId),
		)
	}

	c, _ := models.SongDatabases(
		newDb,
		queryMods...,
	).Count()
	res.MaxSongDatabases = uint64(c)

	if in.SearchWord != "" {
		queryMods = append(
			queryMods,
			qm.Where("song_databases.name LIKE ?", "%"+in.SearchWord+"%"),
		)
	}

	songDatabases, _ := models.SongDatabases(
		newDb,
		queryMods...,
	).All()

	for _, songDatabase := range songDatabases {
		res.SongDatabases = append(
			res.SongDatabases,
			generators.NewSongDatabase(songDatabase),
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

	songDatabases, _ := models.SongDatabases(
		newDb,
		qm.WhereIn("id in (?)", in.SongDatabaseIds),
	).All()

	for _, songDatabaseID := range in.SongDatabaseIds {
		found := false
		for _, songDatabase := range songDatabases {
			if songDatabase.ID != songDatabaseID {
				continue
			}
			found = true
			res.SongDatabases = append(
				res.SongDatabases,
				generators.NewSongDatabase(songDatabase),
			)
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
