package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/generators"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
)

func (s *SeppoServiceServer) CreateEwDatabase(
	ctx context.Context,
	in *SeppoService.CreateEwDatabaseRequest,
) (
	*SeppoService.CreateEwDatabaseResponse,
	error,
) {
	res := &SeppoService.CreateEwDatabaseResponse{}
	newDb := s.getGormDB()
	defer newDb.Close()

	randString, _ := GenerateRandomString(10)

	ewDatabase := db.EwDatabase{
		Name:           in.Name,
		SongDatabaseID: in.SongDatabaseId,
		EwDatabaseKey:  randString,
	}

	newDb.Create(&ewDatabase)

	res.EwDatabase = generators.NewEwDatabase(&ewDatabase)

	return res, nil
}

func (s *SeppoServiceServer) UpdateEwDatabase(
	ctx context.Context,
	in *SeppoService.UpdateEwDatabaseRequest,
) (
	*SeppoService.UpdateEwDatabaseResponse,
	error,
) {
	res := &SeppoService.UpdateEwDatabaseResponse{}
	tx := s.getGormDB().Begin()
	defer tx.Close()

	var ewDatabase db.EwDatabase
	tx.First(&ewDatabase, in.EwDatabaseId)

	if ewDatabase.ID == 0 {
		res.Success = false
		return res, nil
	}

	if in.Name != "" {
		ewDatabase.Name = in.Name
	}

	if in.RemoveSongsFromEwDatabase > 0 {
		if in.RemoveSongsFromEwDatabase == 1 {
			ewDatabase.RemoveSongsFromEwDatabase = true
		} else {
			ewDatabase.RemoveSongsFromEwDatabase = false
		}
	}

	if in.RemoveSongsFromSongDatabase > 0 {
		if in.RemoveSongsFromSongDatabase == 1 {
			ewDatabase.RemoveSongsFromSongDatabase = true
		} else {
			ewDatabase.RemoveSongsFromSongDatabase = false
		}
	}

	if in.VariationVersionConflictAction > 0 {
		ewDatabase.VariationVersionConflictAction = in.VariationVersionConflictAction
	}

	tx.Save(&ewDatabase)
	res.EwDatabase = generators.NewEwDatabase(&ewDatabase)
	res.Success = true
	tx.Commit()

	return res, nil
}

func (s *SeppoServiceServer) RemoveEwDatabase(
	ctx context.Context,
	in *SeppoService.RemoveEwDatabaseRequest,
) (
	*SeppoService.RemoveEwDatabaseResponse,
	error,
) {
	res := &SeppoService.RemoveEwDatabaseResponse{}
	newDb := s.getGormDB()
	defer newDb.Close()

	var ewDatabase db.EwDatabase
	newDb.First(&ewDatabase, in.EwDatabaseId)

	if ewDatabase.ID == 0 {
		res.Success = false
		return res, nil
	}

	newDb.Delete(&ewDatabase)
	res.Success = true

	return res, nil
}

func (s *SeppoServiceServer) SearchEwDatabases(
	ctx context.Context,
	in *SeppoService.SearchEwDatabasesRequest,
) (
	*SeppoService.SearchEwDatabasesResponse,
	error,
) {
	res := &SeppoService.SearchEwDatabasesResponse{}
	newDb := s.getGormDB()
	defer newDb.Close()

	query := newDb.Table("ew_databases")

	if in.SongDatabaseId > 0 {
		query = query.Where("song_database_id = ?", in.SongDatabaseId)
	}

	query.Count(&res.MaxEwDatabases)

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}

	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	}

	ewDatabases := []db.EwDatabase{}
	query.Find(&ewDatabases)

	for i := 0; i < len(ewDatabases); i++ {
		res.EwDatabases = append(
			res.EwDatabases,
			generators.NewEwDatabase(&ewDatabases[i]),
		)
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchEwDatabaseById(
	ctx context.Context,
	in *SeppoService.FetchEwDatabaseByIdRequest,
) (
	*SeppoService.FetchEwDatabaseByIdResponse,
	error,
) {
	res := &SeppoService.FetchEwDatabaseByIdResponse{}
	newDb := s.getGormDB()
	defer newDb.Close()

	ewDatabases := []db.EwDatabase{}
	newDb.Where("id in (?)", in.EwDatabaseIds).Find(&ewDatabases)
	for _, ewDatabaseId := range in.EwDatabaseIds {
		found := false
		for _, ewDatabase := range ewDatabases {
			if ewDatabaseId != ewDatabase.ID {
				continue
			}
			found = true
			res.EwDatabases = append(
				res.EwDatabases,
				generators.NewEwDatabase(&ewDatabase),
			)
			break
		}
		if found == false {
			res.EwDatabases = append(
				res.EwDatabases,
				&SeppoService.EwDatabase{},
			)
		}
	}

	return res, nil
}
