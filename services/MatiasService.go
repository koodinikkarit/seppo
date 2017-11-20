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

func getEwDatabaseLink(
	newDb *gorm.DB,
	ewDatabaseLinks map[uint32]*db.EwDatabaseLink,
	ewDatabaseId uint32,
	ewSongId uint32,
) *db.EwDatabaseLink {
	ewDatabaseLink := ewDatabaseLinks[ewSongId]
	if ewDatabaseLink == nil {
		newDb.Where("ew_database_id = ?", ewDatabaseId).
			Where("ew_database_song_id = ?", ewSongId).First(&ewDatabaseLink)
		if ewDatabaseLink != nil {
			ewDatabaseLinks[ewSongId] = ewDatabaseLink
		}
	}
	return ewDatabaseLink
}

func getVariationVersion(
	newDb *gorm.DB,
	variationVersions map[uint32]*db.VariationVersion,
	variationVersionId uint32,
) *db.VariationVersion {
	variationVersion := variationVersions[variationVersionId]
	if variationVersion == nil {
		newDb.Where("variation_version_id = ?", variationVersionId).Find(&variationVersion)
		if variationVersion != nil {
			variationVersions[variationVersionId] = variationVersion
		}
	}
	return variationVersion
}

func getVariation(
	newDb *gorm.DB,
	variations map[uint32]*db.Variation,
	variationId uint32,
) *db.Variation {
	variation := variations[variationId]
	if variation == nil {
		newDb.Where("variations.id = ?", variationId).First(&variation)
		if variation != nil {
			variations[variationId] = variation
		}
	}
	return variation
}

func createBranchAndAddItToSongDatabase(
	tx *gorm.DB,
	ewDatabase *db.EwDatabase,
	ewSong *MatiasService.EwSong,
	sameVariationVersion *db.VariationVersion,
) {
	newVariation := db.Variation{}
	tx.Create(&newVariation)
	newVariationVersion := db.VariationVersion{
		VariationID: newVariation.ID,
		Name:        ewSong.Title,
		Text:        ewSong.Text,
		Version:     1,
		Newest:      true,
	}
	tx.Save(&newVariationVersion)
	newBranch := db.Branch{
		SourceVariationVersionID:      sameVariationVersion.ID,
		DestinationVariationVersionID: newVariationVersion.ID,
	}
	tx.Create(&newBranch)
	tx.Create(&db.SongDatabaseVariation{
		SongDatabaseID: ewDatabase.SongDatabaseID,
		VariationID:    newVariationVersion.ID,
	})
	tx.Create(&db.EwDatabaseLink{
		EwDatabaseID:     ewDatabase.ID,
		EwDatabaseSongID: ewSong.Id,
		VariationID:      newVariationVersion.ID,
	})
}

func (s *MatiasServiceServer) SyncEwSong(
	stream MatiasService.Matias_SyncEwSongServer,
) error {
	// tx := s.getDB().Begin()
	// defer tx.Close()

	// var ewDatabase *db.EwDatabase
	// variations := make(map[uint32]*db.Variation)
	// variationVersions := make(map[uint32]*db.VariationVersion)
	// ewDatabaseLinks := make(map[uint32]*db.EwDatabaseLink)

	// for {
	// 	in, err := stream.Recv()
	// 	if err == io.EOF {
	// 		return nil
	// 	}
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}

	// 	if ewDatabase == nil {
	// 		tx.Where("ew_database_key = ?", in.EwDatabaseKey).First(&ewDatabase)
	// 	}

	// 	if ewDatabase != nil {
	// 		ewDatabaseLink := getEwDatabaseLink(tx, ewDatabaseLinks, ewDatabase.ID, in.EwSong.Id)

	// 		if ewDatabaseLink != nil {
	// 			variationVersion := getVariationVersion(
	// 				tx,
	// 				variationVersions,
	// 				ewDatabaseLink.VariationVersionID,
	// 			)
	// 			if variationVersion != nil {
	// 				if variationVersion.Newest == true {
	// 					if variationVersion.DisabledAt != nil {
	// 						if ewDatabase.RemoveSongsFromEwDatabase == true {
	// 							// poistetaann
	// 						}
	// 					} else {
	// 						if variationVersion.Name != in.EwSong.Title ||
	// 							variationVersion.Text != in.EwSong.Text {
	// 							createNewVariationVersions(
	// 								tx,
	// 								variationVersion,
	// 								in.EwSong.Title,
	// 								in.EwSong.Text,
	// 							)
	// 						}
	// 					}
	// 				} else {
	// 					variation := getVariation(
	// 						tx,
	// 						variations,
	// 						variationVersion.VariationID,
	// 					)
	// 					newestVariationVersion := getVariationVersion(
	// 						tx,
	// 						variationVersions,
	// 						*variation.VariationVersionID,
	// 					)
	// 					if newestVariationVersion.DisabledAt != nil {
	// 						// poistetaan
	// 					} else {
	// 						if newestVariationVersion.Name != in.EwSong.Title ||
	// 							newestVariationVersion.Text != in.EwSong.Text {

	// 							switch ewDatabase.VariationVersionConflictAction {
	// 							case 1:
	// 							case 2:
	// 							case 3:
	// 							case 4:
	// 							}
	// 						} else {
	// 							ewDatabaseLink.VariationVersionID = newestVariationVersion.ID
	// 							tx.Save(&ewDatabaseLink)
	// 						}
	// 					}
	// 				}
	// 			}
	// 		} else {
	// 			var sameVariationVersion db.VariationVersion
	// 			tx.Where("name = ?", in.EwSong.Title).
	// 				Where("text = ?", in.EwSong.Text).
	// 				First(&sameVariationVersion)

	// 			if sameVariationVersion.ID > 0 {
	// 				if sameVariationVersion.Newest == true {
	// 					if sameVariationVersion.DisabledAt != nil {
	// 						createBranchAndAddItToSongDatabase(
	// 							tx,
	// 							ewDatabase,
	// 							in.EwSong,
	// 							&sameVariationVersion,
	// 						)
	// 					} else {

	// 					}
	// 				} else {
	// 					createBranchAndAddItToSongDatabase(
	// 						tx,
	// 						ewDatabase,
	// 						in.EwSong,
	// 						&sameVariationVersion,
	// 					)
	// 				}
	// 			} else {
	// 				newVariation := db.Variation{}
	// 				tx.Create(&newVariation)
	// 				newVariationVersion := db.VariationVersion{
	// 					VariationID: newVariation.ID,
	// 					Name:        in.EwSong.Title,
	// 					Text:        in.EwSong.Text,
	// 					Version:     1,
	// 					Newest:      true,
	// 				}
	// 				newVariation.VariationVersionID = &newVariationVersion.ID
	// 				tx.Save(&newVariationVersion)
	// 			}
	// 		}
	// 	}

	// }

	return nil
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
