package services

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/koodinikkarit/seppo/managers"
	"github.com/koodinikkarit/seppo/matias_service"
	"github.com/koodinikkarit/seppo/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

type MatiasServiceServer struct {
	getDB func() *sql.DB
}

func StartMatiasService(
	port string,
	getDB func() *sql.DB,
) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	MatiasService.RegisterMatiasServer(s, &MatiasServiceServer{
		getDB: getDB,
	})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *MatiasServiceServer) RequestMatiasKey(
	ctx context.Context,
	in *MatiasService.RequestMatiasKeyRequest,
) (
	*MatiasService.RequestMatiasKeyResponse,
	error,
) {
	res := &MatiasService.RequestMatiasKeyResponse{}
	//newDb := s.getDB()

	// randString, _ := GenerateRandomString(10)

	// newMatiasClient := db.MatiasClient{
	// 	ClientKey: randString,
	// }

	// newDb.Create(&newMatiasClient)

	// res.Key = newMatiasClient.ClientKey

	return res, nil
}

func (s *MatiasServiceServer) InsertEwSongIds(
	ctx context.Context,
	in *MatiasService.InsertEwSongIdsRequest,
) (
	*MatiasService.InsertEwSongIdsResponse,
	error,
) {
	res := &MatiasService.InsertEwSongIdsResponse{}
	newDb := s.getDB()
	defer newDb.Close()
	tx, _ := newDb.Begin()

	ewDatabase, _ := models.EwDatabases(
		tx,
		qm.Where("ew_databases.ew_database_key = ?", in.EwDatabaseKey),
	).One()

	if ewDatabase == nil {
		res.EwDatabaseFound = false
		return res, nil
	}

	var variationIds []uint32
	for _, variationIdEwSongId := range in.VariationIdEwSongIds {
		variationIds = append(
			variationIds,
			variationIdEwSongId.VariationId,
		)
	}

	variations, _ := models.Variations(
		tx,
		qm.Where("id in ?", variationIds),
		qm.Load("VariationVersions"),
	).All()

	for _, variationIdEwSongId := range in.VariationIdEwSongIds {
		for _, variation := range variations {
			if uint64(variationIdEwSongId.VariationId) != variation.ID {
				continue
			}
			variationVersions, _ := variation.VariationVersions(tx).All()
			newestVariationVersion := managers.FindNewestVariationVersion(
				variationVersions,
			)
			if newestVariationVersion == nil {
				break
			}
			ewDatabaseLink := models.EwDatabaseLink{
				EwDatabaseID:     ewDatabase.ID,
				EwDatabaseSongID: uint64(variationIdEwSongId.EwSongId),
				VariationID:      uint64(variationIdEwSongId.VariationId),
				Version:          newestVariationVersion.Version,
			}
			ewDatabaseLink.Insert(tx)
			break
		}
	}

	var ewSongIDs []uint32
	for _, link := range in.NewSongIds {
		ewSongIDs = append(ewSongIDs, link.OldEwSongId)
	}

	ewDatabaseLinks, _ := models.EwDatabaseLinks(
		tx,
		qm.Where("ew_database_song_id in ?", ewSongIDs),
	).All()

	for _, ewDatabaseLink := range ewDatabaseLinks {
		for _, newSongId := range in.NewSongIds {
			if ewDatabaseLink.EwDatabaseSongID != uint64(newSongId.OldEwSongId) {
				continue
			}
			ewDatabaseLink.EwDatabaseSongID = uint64(newSongId.NewEwSongId)
			ewDatabaseLink.Update(tx, "ew_database_song_id")
		}
	}

	tx.Commit()

	return res, nil
}

func (s *MatiasServiceServer) RequestEwChanges(
	in *MatiasService.RequestEwDatabaseChangesRequest,
	stream MatiasService.Matias_RequestEwChangesServer,
) error {
	return nil
}
