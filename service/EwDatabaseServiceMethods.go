package service

import (
	"golang.org/x/net/context"

	SeppoService "github.com/koodinikkarit/go-clientlibs/seppo"
	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/matias"
	"github.com/koodinikkarit/seppo/models"
)

func (s *SeppoServiceServer) CreateEwDatabase(
	ctx context.Context,
	in *SeppoService.CreateEwDatabaseRequest,
) (
	*SeppoService.CreateEwDatabaseResponse,
	error,
) {
	res := &SeppoService.CreateEwDatabaseResponse{}
	db := s.getDB()

	ewDatabase := models.EwDatabase{
		Name:           in.Name,
		SongDatabaseID: in.SongDatabaseId,
		FilesystemPath: in.FilesystemPath,
		MatiasClientID: in.MatiasClientId,
	}

	db.Create(&ewDatabase)

	s.pubSub.Pub(matias.CreatedEwDatabaseEvent{
		EwDatabase: ewDatabase,
	}, "createdEwDatabase")

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
	tx := s.getDB().Begin()

	var ewDatabase models.EwDatabase
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

	s.pubSub.Pub(matias.UpdatedEwDatabaseEvent{
		EwDatabase: ewDatabase,
	}, "updatedEwDatabase")

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
	db := s.getDB()

	var ewDatabase models.EwDatabase
	db.First(&ewDatabase, in.EwDatabaseId)

	if ewDatabase.ID == 0 {
		res.Success = false
		return res, nil
	}

	db.Delete(&ewDatabase)
	res.Success = true

	s.pubSub.Pub(matias.RemovedEwDatabaseEvent{
		EwDatabaseId: in.EwDatabaseId,
	}, "updatedEwDatabase")

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
	db := s.getDB()

	query := db.Table("ew_databases")

	if in.SongDatabaseId > 0 {
		query = query.Where("song_database_id = ?", in.SongDatabaseId)
	}

	if in.MatiasClientId > 0 {
		query = query.Where("matias_client_id = ?", in.MatiasClientId)
	}

	query.Count(&res.MaxEwDatabases)

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}

	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	}

	ewDatabases := []models.EwDatabase{}
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
	db := s.getDB()

	ewDatabases := []models.EwDatabase{}
	db.Where("id in (?)", in.EwDatabaseIds).Find(&ewDatabases)
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
