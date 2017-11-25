package services

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/managers"
	"github.com/koodinikkarit/seppo/matias_service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

type MatiasServiceServer struct {
	getDB     func() *sql.DB
	getGormDB func() *gorm.DB
}

func StartMatiasService(
	port string,
	getDB func() *sql.DB,
	getGormDB func() *gorm.DB,
) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	MatiasService.RegisterMatiasServer(s, &MatiasServiceServer{
		getDB:     getDB,
		getGormDB: getGormDB,
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
	tx := s.getGormDB().Begin()
	defer tx.Close()

	var ewDatabase *db.EwDatabase
	tx.Where("ew_databases.ew_database_key = ?", in.EwDatabaseKey).
		Find(&ewDatabase)

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

	var variations []db.Variation
	tx.Where("id in (?)", variationIds).
		Preload("VariationVersions").
		Find(&variations)

	var newEwDatabaseLinks []db.EwDatabaseLink
	for _, variationIdEwSongId := range in.VariationIdEwSongIds {
		for _, variation := range variations {
			if variationIdEwSongId.VariationId != variation.ID {
				continue
			}
			newestVariationVersion := variation.FindNewestVersion()
			if newestVariationVersion.ID == 0 {
				break
			}
			ewDatabaseLink := db.EwDatabaseLink{
				EwDatabaseID:     ewDatabase.ID,
				EwDatabaseSongID: variationIdEwSongId.EwSongId,
				VariationID:      variationIdEwSongId.VariationId,
				Version:          newestVariationVersion.Version,
			}

			newEwDatabaseLinks = append(
				newEwDatabaseLinks,
				ewDatabaseLink,
			)
			break
		}
	}

	managers.BatchAddVariationsToEwDatabase(
		tx,
		newEwDatabaseLinks,
	)

	var ewSongIDs []uint32
	for _, link := range in.NewSongIds {
		ewSongIDs = append(ewSongIDs, link.OldEwSongId)
	}
	var ewDatabaseLinks []db.EwDatabaseLink
	tx.Where("ew_database_song_id in (?)", ewSongIDs).
		Find(&ewDatabaseLinks)

	for _, ewDatabaseLink := range ewDatabaseLinks {
		for _, newSongId := range in.NewSongIds {
			if ewDatabaseLink.EwDatabaseSongID != newSongId.OldEwSongId {
				continue
			}
			tx.Model(&ewDatabaseLink).
				Update("ew_database_song_id", newSongId.NewEwSongId)
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
