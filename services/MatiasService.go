package services

import (
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
	getDB func() *gorm.DB
}

func StartMatiasService(
	port string,
	getDB func() *gorm.DB,
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
	newDb := s.getDB()

	randString, _ := GenerateRandomString(10)

	newMatiasClient := db.MatiasClient{
		ClientKey: randString,
	}

	newDb.Create(&newMatiasClient)

	res.Key = newMatiasClient.ClientKey

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
	tx := s.getDB().Begin()

	var ewDatabase db.EwDatabase
	tx.Where("ew_databases.ew_database_key = ?", in.EwDatabaseKey).
		First(&ewDatabase)

	if ewDatabase.ID == 0 {
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

	variations := []db.Variation{}
	tx.Where("id in (?)", variationIds).
		Preload("VariationVersions").
		Find(&variations)

	var newEwDatabaseLinks []db.EwDatabaseLink

	for _, variationIdEwSongId := range in.VariationIdEwSongIds {
		for _, variation := range variations {
			newestVersion := variation.FindNewestVersion()
			if newestVersion.ID > 0 {
				if variationIdEwSongId.VariationId == variation.ID {
					newEwDatabaseLinks = append(
						newEwDatabaseLinks,
						db.EwDatabaseLink{
							EwDatabaseID:     ewDatabase.ID,
							EwDatabaseSongID: variationIdEwSongId.EwSongId,
							VariationID:      variationIdEwSongId.VariationId,
							Version:          newestVersion.Version,
						},
					)
				}
			}
		}
	}

	//tx.Model(&ewDatabaseLink).UpdateColumn("version", newestVariationVersion.Version)

	managers.BatchAddVariationsToEwDatabase(
		tx,
		newEwDatabaseLinks,
	)

	var ewSongIDs []uint32
	ewDatabaseLinks := []db.EwDatabaseLink{}

	for _, link := range in.NewSongIds {
		ewSongIDs = append(ewSongIDs, link.OldEwSongId)
	}

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

// variations := []db.Variation{}
// variationVersions := []db.VariationVersion{}

// var ewDatabase db.EwDatabase
// tx.Where("ew_databases.ew_database_key = ?", in.EwDatabaseKey).
// 	Find(&ewDatabase)

// if ewDatabase.ID > 0 {

// 	songDatabaseVariationVersions := []db.VariationVersion{}

// 	tx.Table("song_database_variations").
// 		Joins("left join song_database_variations on variation_versions.id = song_database_variations.variation_version_id").
// 		Joins("left join tag_variations on variation_versions.id = tag_variations.variation_version_id").
// 		Joins("left join song_database_tags on tag_variations.tag_id = song_database_tags.tag_id").
// 		Where("song_database_variations.song_database_id = ? or song_database_tags.song_database_id", ewDatabase.SongDatabaseID, ewDatabase.SongDatabaseID).
// 		Find(&songDatabaseVariationVersions)

// 	variationIds := []db.VariationVersion{}
// 	for i := 0; i < len(songDatabaseVariationVersions); i++ {
// 		variationIds = append(
// 			variationIds,
// 			songDatabaseVariationVersions[i],
// 		)
// 	}

// 	variations := []db.Variation{}
// 	newestVariationVersions := []db.VariationVersion{}

// 	tx.Where("variations.id in (?)", variationIds).
// 		Find(&variations)

// 	newestVariationVersionIds := []db.VariationVersion{}

// 	for i := 0; i < len(variations); i++ {
// 		if variations[i].VariationVersionID != nil {
// 			newestVariationVersionIds = append(
// 				newestVariationVersionIds,
// 				variations[i].VariationVersionID,
// 			)
// 		}
// 	}

// 	tx.Where("variation_versions.id in (?)", newestVariationVersionIds).
// 		Find(&newestVariationVersions)

// 	var ewDatabaseLinks []db.EwDatabaseLink{}
// 	tx.Where("ew_database_id = ?", ewDatabase.ID).
// 		Find(&ewDatabaseLinks)

// 	for i := 0; i < len(in.EwSongs); i++ {
// 		foundEwDatabaseLink := false
// 		for j := 0; j < len(ewDatabaseLinks); j++ {
// 			if ewDatabaseLinks[j].EwDatabaseSongID == in.EwSongs[i] {
// 				foundEwDatabaseLink = true
// 				foundVariationVersion := false
// 				for x := 0; x < len(songDatabaseVariationVersions); x++ {
// 					if ewDatabaseLinks[j].VariationVersionID == songDatabaseVariationVersions[x] {
// 						foundVariationVersion = true
// 						if songDatabaseVariationVersions[x].Newest == true {
// 							if songDatabaseVariationVersions[x].DisabledAt != nil {

// 							} else {
// 								if songDatabaseVariationVersions[x].Name != in.EwSongs[i].Title ||
// 									songDatabaseVariationVersions[x].Text != in.EwSongs[i].Text {

// 								}
// 							}
// 						} else {

// 						}
// 					}
// 				}
// 			}
// 		}
// 	}

// }
