package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/models"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (s *SeppoServiceServer) SearchEwDatabases(
	ctx context.Context,
	in *SeppoService.SearchEwDatabasesRequest,
) (
	*SeppoService.SearchEwDatabasesResponse,
	error,
) {
	res := &SeppoService.SearchEwDatabasesResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	var queryMods []qm.QueryMod

	queryMods = append(
		queryMods,
		qm.From("ew_databases"),
	)

	if in.SongDatabaseId > 0 {
		queryMods = append(
			queryMods,
			qm.Where("song_database_id = ?", in.SongDatabaseId),
		)
	}

	c, _ := models.EwDatabases(newDb, queryMods...).Count()
	res.MaxEwDatabases = uint64(c)

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

	ewDatabases, _ := models.EwDatabases(
		newDb,
		queryMods...,
	).All()

	for _, ewDatabase := range ewDatabases {
		res.EwDatabases = append(
			res.EwDatabases,
			generators.NewEwDatabase(ewDatabase),
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
	newDb := s.getDB()
	defer newDb.Close()

	ewDatabases, _ := models.EwDatabases(
		newDb,
		qm.WhereIn("id in ?", in.EwDatabaseIds),
	).All()

	for _, ewDatabaseId := range in.EwDatabaseIds {
		found := false
		for _, ewDatabase := range ewDatabases {
			if ewDatabaseId == ewDatabase.ID {
				found = true
				res.EwDatabases = append(
					res.EwDatabases,
					generators.NewEwDatabase(ewDatabase),
				)
			}
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

func (s *SeppoServiceServer) CreateEwDatabase(
	ctx context.Context,
	in *SeppoService.CreateEwDatabaseRequest,
) (
	*SeppoService.CreateEwDatabaseResponse,
	error,
) {
	res := &SeppoService.CreateEwDatabaseResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	randString, _ := GenerateRandomString(10)

	ewDatabase := models.EwDatabase{
		Name:           in.Name,
		SongDatabaseID: in.SongDatabaseId,
		EwDatabaseKey:  randString,
	}
	ewDatabase.Insert(newDb)

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
	newDb := s.getDB()
	defer newDb.Close()

	ewDatabase, _ := models.FindEwDatabase(
		newDb,
		in.EwDatabaseId,
	)

	if ewDatabase == nil {
		res.Success = true
		return res, nil
	}

	if in.RemoveSongsFromEwDatabase > 0 {
		if in.RemoveSongsFromEwDatabase == 1 {
			ewDatabase.RemoveSongsFromEwDatabase = 1
		} else {
			ewDatabase.RemoveSongsFromEwDatabase = 0
		}
	}

	if in.RemoveSongsFromSongDatabase > 0 {
		if in.RemoveSongsFromSongDatabase == 1 {
			ewDatabase.RemoveSongsFromSongDatabase = 1
		} else {
			ewDatabase.RemoveSongsFromSongDatabase = 0
		}
	}

	if in.VariationVersionConflictAction > 0 {
		ewDatabase.VariationVersionConflictAction = uint(in.VariationVersionConflictAction)
	}

	ewDatabase.Update(newDb)
	res.EwDatabase = generators.NewEwDatabase(ewDatabase)
	res.Success = true

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
	newDb := s.getDB()
	defer newDb.Close()

	ewDatabase, _ := models.FindEwDatabase(
		newDb,
		in.EwDatabaseId,
	)

	if ewDatabase == nil {
		res.Success = true
		return res, nil
	}

	ewDatabase.Delete(newDb)
	res.Success = true
	return res, nil
}
