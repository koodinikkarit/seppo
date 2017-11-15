package services

import (
	"log"
	"net"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/help"
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
		for i := 0; i < len(in.NewSongIds); i++ {
			if ewDatabaseLink.EwDatabaseSongID == in.NewSongIds[i].OldEwSongId {
				tx.Model(&ewDatabaseLink).
					Update("ew_database_song_id", ewDatabaseLink.EwDatabaseSongID)
			}
		}
	}

	tx.Commit()

	return res, nil
}

func hasEwSongByNameAndText(
	ewSongs map[uint32]*MatiasService.EwSong,
	name string,
	text string,
) bool {
	for _, ewSong := range ewSongs {
		if ewSong.Title == name &&
			ewSong.Text == text {
			return true
		}
	}
	return false
}

func removeDuplicatesFromEwSong(
	ewSongs []*MatiasService.EwSong,
) (
	map[uint32]*MatiasService.EwSong,
	[]uint32,
) {
	newEwSongs := make(map[uint32]*MatiasService.EwSong)
	var removeEwSongs []uint32
	excludeEwSongs := make(map[uint32]*MatiasService.EwSong)
	for i := 0; i < len(ewSongs); i++ {
		if excludeEwSongs[ewSongs[i].Id] == nil {
			newEwSongs[ewSongs[i].Id] = ewSongs[i]
			if i != len(ewSongs)-1 {
				for j := len(ewSongs) - 1; j > i; j-- {
					if ewSongs[j].Title == ewSongs[i].Title &&
						ewSongs[j].Text == ewSongs[i].Text {
						excludeEwSongs[ewSongs[j].Id] = ewSongs[j]
						removeEwSongs = append(
							removeEwSongs,
							ewSongs[j].Id,
						)
					}
				}
			}
		}
	}
	return newEwSongs, removeEwSongs
}

func (s *MatiasServiceServer) SyncEwDatabase(
	ctx context.Context,
	in *MatiasService.SyncEwDatabaseRequest,
) (
	*MatiasService.SyncEwDatabaseResponse,
	error,
) {
	res := &MatiasService.SyncEwDatabaseResponse{}
	tx := s.getDB().Begin()
	defer tx.Close()

	var synchronizationRaport db.SynchronizationRaport
	tx.Create(&synchronizationRaport)
	startDate := time.Now()
	startTime := startDate.UnixNano() / int64(time.Millisecond)

	synchronizationRaport.StartedAt = &startDate
	synchronizationRaport.DatabaseKey = in.EwDatabaseKey

	var ewDatabase db.EwDatabase
	tx.Where("ew_databases.ew_database_key = ?", in.EwDatabaseKey).
		Preload("SongDatabase").
		Preload("SongDatabase.Variations").
		Preload("SongDatabase.Variations.Author").
		Preload("SongDatabase.Variations.Copyright").
		Preload("SongDatabase.Variations.VariationVersions").
		Preload("SongDatabase.SongDatabaseTags").
		Preload("EwDatabaseLinks").
		Preload("EwDatabaseLinks.Variation").
		Preload("EwDatabaseLinks.Variation.Author").
		Preload("EwDatabaseLinks.Variation.Copyright").
		Preload("EwDatabaseLinks.Variation.VariationVersions").
		Find(&ewDatabase)

	if ewDatabase.ID == 0 {
		res.EwDatabaseFound = false
		synchronizationRaport.DatabaseFound = false
		endDate := time.Now()
		endTime := endDate.UnixNano() / int64(time.Millisecond)
		synchronizationRaport.DurationMS = endTime - startTime
		synchronizationRaport.FinishedAt = &endDate
		tx.Save(&synchronizationRaport)
		tx.Commit()
		return res, nil
	}

	synchronizationRaport.DatabaseFound = true
	synchronizationRaport.DatabaseID = &ewDatabase.ID

	// These slices are used for batch insert
	var addVariationsToSongDatabase []db.SongDatabaseVariation
	var addVariationsToEwDatabase []db.EwDatabaseLink
	var createBranches []db.Branch
	//var srEwConflicts []db.SrEwConflict
	var createNewVariationVersions []db.VariationVersion

	// These slices are used for batch remove
	var removeEwDatabaseLinks []uint32
	var removeSongDatabaseVariationLinks []uint32

	// These are used for synchronization raport generation
	var srAddSongDatabaseVariations []db.SrAddSongDatabaseVariation
	var srEwConflicts []db.SrEwConflict
	var srEwDatabaseLinks []db.SrEwDatabaseLink
	var srEwSongs []db.SrEwSong
	var srNewAuthors []db.SrNewAuthors
	var srNewBranches []db.SrNewBranch
	var srNewCopyrights []db.SrcNewCopyright
	var srNewVariations []db.SrNewVariation
	var srNewVariationVersions []db.SrNewVariationVersion
	var srNewPassivatedVariationVersions []db.SrPassivatedVariationVersion
	var srRemoveSongDatabaseVariations []db.SrRemoveSongDatabaseVariation

	ewSongs, removeEwSongIds := removeDuplicatesFromEwSong(in.EwSongs)
	res.RemoveEwSongIds = append(res.RemoveEwSongIds, removeEwSongIds...)

	for _, ewSong := range ewSongs {
		ewDatabaseLink := ewDatabase.FindEwDatabaseLinkByEwSongID(ewSong.Id)
		if ewDatabaseLink == nil {
			var sameVariation db.Variation

			tx.Table("variations").
				Joins("left join variation_versions on variations.id = variation_versions.variation_id").
				Where("variation_versions.name = ?", ewSong.Title).
				Where("variation_versions.text = ?", ewSong.Text).
				Preload("VariationVersions").
				Preload("SongDatabases").
				Preload("TagVariations").
				First(&sameVariation)

			if sameVariation.ID == 0 {
				// No variation with same name and text

				newVariation := managers.NewVariationFromEwSong(
					tx,
					ewSong,
				)
				// Add variation and version to list for raport generation
				srNewVariations = append(
					srNewVariations,
					db.SrNewVariation{
						SrID:        synchronizationRaport.ID,
						VariationID: newVariation.ID,
					},
				)
				srNewVariationVersions = append(
					srNewVariationVersions,
					db.SrNewVariationVersion{
						SrID:               synchronizationRaport.ID,
						VariationVersionID: newVariation.VariationVersions[0].ID,
					},
				)

				// These are added to slice for batch insert
				addVariationsToSongDatabase = append(
					addVariationsToSongDatabase,
					db.SongDatabaseVariation{
						SongDatabaseID: ewDatabase.SongDatabaseID,
						VariationID:    newVariation.ID,
					},
				)
				addVariationsToEwDatabase = append(
					addVariationsToEwDatabase,
					db.EwDatabaseLink{
						EwDatabaseID:     ewDatabase.ID,
						EwDatabaseSongID: ewSong.Id,
						VariationID:      newVariation.ID,
						Version:          1,
					},
				)
			} else {
				newestVariationVersion := sameVariation.FindNewestVersion()
				sameVariationVersion := sameVariation.FindVariationVersionByNameAndText(
					ewSong.Title,
					ewSong.Text,
				)
				if sameVariationVersion.Version == newestVariationVersion.Version {
					if sameVariationVersion.DisabledAt == nil {
						if sameVariation.FindSongDatabaseByID(ewDatabase.SongDatabaseID) == nil {
							foundTag := false
							for _, tagVariation := range sameVariation.TagVariations {
								if ewDatabase.SongDatabase.HasSongDatabaseTag(tagVariation.TagID) == true {
									foundTag = true
									break
								}
							}
							if foundTag == false {

								if ewSong.Author != "" &&
									sameVariation.AuthorID == nil {

									newAuthor := db.CreateAuthorByName(
										tx,
										ewSong.Author,
									)
									tx.Model(&sameVariation).
										Update("author_id", newAuthor.ID)
								}

								if ewSong.Copyright != "" &&
									sameVariation.CopyrightID == nil {

									newCopyright := db.CreateCopyrightByName(
										tx,
										ewSong.Copyright,
									)
									tx.Model(&sameVariation).
										Update("copyright_id", newCopyright.ID)
								}

								addVariationsToSongDatabase = append(
									addVariationsToSongDatabase,
									db.SongDatabaseVariation{
										SongDatabaseID: ewDatabase.SongDatabaseID,
										VariationID:    sameVariation.ID,
									},
								)
								addVariationsToEwDatabase = append(
									addVariationsToEwDatabase,
									db.EwDatabaseLink{
										EwDatabaseID:     ewDatabase.ID,
										EwDatabaseSongID: ewSong.Id,
										VariationID:      sameVariation.ID,
										Version:          sameVariationVersion.Version,
									},
								)
							}
						}
					} else {
						newVariation := managers.NewVariationFromEwSong(
							tx,
							ewSong,
						)

						srNewVariations = append(
							srNewVariations,
							db.SrNewVariation{
								SrID:        synchronizationRaport.ID,
								VariationID: newVariation.ID,
							},
						)
						srNewVariationVersions = append(
							srNewVariationVersions,
							db.SrNewVariationVersion{
								SrID:               synchronizationRaport.ID,
								VariationVersionID: newVariation.VariationVersions[0].ID,
							},
						)

						createBranches = append(
							createBranches,
							db.Branch{
								SourceVariationVersionID:      sameVariationVersion.ID,
								DestinationVariationVersionID: newVariation.VariationVersions[0].ID,
							},
						)

						addVariationsToSongDatabase = append(
							addVariationsToSongDatabase,
							db.SongDatabaseVariation{
								SongDatabaseID: ewDatabase.SongDatabaseID,
								VariationID:    newVariation.ID,
							},
						)
						addVariationsToEwDatabase = append(
							addVariationsToEwDatabase,
							db.EwDatabaseLink{
								EwDatabaseID:     ewDatabase.ID,
								EwDatabaseSongID: ewSong.Id,
								VariationID:      newVariation.ID,
								Version:          1,
							},
						)
					}
				} else {
					newVariation := managers.NewVariationFromEwSong(
						tx,
						ewSong,
					)

					srNewVariations = append(
						srNewVariations,
						db.SrNewVariation{
							SrID:        synchronizationRaport.ID,
							VariationID: newVariation.ID,
						},
					)
					srNewVariationVersions = append(
						srNewVariationVersions,
						db.SrNewVariationVersion{
							SrID:               synchronizationRaport.ID,
							VariationVersionID: newVariation.VariationVersions[0].ID,
						},
					)

					createBranches = append(
						createBranches,
						db.Branch{
							SourceVariationVersionID:      sameVariationVersion.ID,
							DestinationVariationVersionID: newVariation.VariationVersions[0].ID,
						},
					)

					addVariationsToSongDatabase = append(
						addVariationsToSongDatabase,
						db.SongDatabaseVariation{
							SongDatabaseID: ewDatabase.SongDatabaseID,
							VariationID:    newVariation.ID,
						},
					)
					addVariationsToEwDatabase = append(
						addVariationsToEwDatabase,
						db.EwDatabaseLink{
							EwDatabaseID:     ewDatabase.ID,
							EwDatabaseSongID: ewSong.Id,
							VariationID:      newVariation.ID,
							Version:          1,
						},
					)
				}
			}
		} else {
			if ewDatabaseLink.Variation.ID > 0 {
				if ewSong.Author != "" {
					if ewDatabaseLink.Variation.Author.ID > 0 {
						if ewDatabaseLink.Author == ewSong.Author {
							if ewDatabaseLink.Variation.Author.Name != ewSong.Author {
								newEwSong := res.CreateOrGetEwSong(ewSong.Id)
								newEwSong.Author = ewDatabaseLink.Variation.Author.Name
								tx.Model(&ewDatabaseLink).
									Update("author", ewDatabaseLink.Variation.Author.Name)
							}
						} else {
							if ewDatabaseLink.Variation.Author.Name != ewDatabaseLink.Author {
								newEwSong := res.CreateOrGetEwSong(ewSong.Id)
								newEwSong.Author = ewDatabaseLink.Variation.Author.Name
								tx.Model(&ewDatabaseLink).
									Update("author", ewDatabaseLink.Variation.Author.Name)
							} else {
								tx.Model(&ewDatabaseLink).
									Update("author", ewSong.Author)
								tx.Model(&ewDatabaseLink.Variation.Author).
									Update("name", ewSong.Author)
							}
						}
					} else {
						newAuthor := db.CreateAuthorByName(
							tx,
							ewSong.Author,
						)
						tx.Model(&ewDatabaseLink.Variation).
							Update("author_id", newAuthor.ID)

						if ewDatabaseLink.Author != ewSong.Author {
							tx.Model(&ewDatabaseLink).
								Update("author", ewSong.Author)
						}
					}
				} else {
					if ewDatabaseLink.Variation.Author.ID > 0 {
						newEwSong := res.CreateOrGetEwSong(ewSong.Id)
						newEwSong.Author = ewDatabaseLink.Variation.Author.Name
					}
				}

				if ewSong.Copyright != "" {
					if ewDatabaseLink.Variation.Copyright.ID > 0 {
						if ewDatabaseLink.Copyright == ewSong.Copyright {
							if ewDatabaseLink.Variation.Copyright.Name != ewSong.Copyright {
								newEwSong := res.CreateOrGetEwSong(ewSong.Id)
								newEwSong.Copyright = ewDatabaseLink.Variation.Copyright.Name
								tx.Model(&ewDatabaseLink).
									Update("copyright", ewDatabaseLink.Variation.Copyright.Name)
							}
						} else {
							if ewDatabaseLink.Variation.Copyright.Name != ewDatabaseLink.Copyright {
								newEwSong := res.CreateOrGetEwSong(ewSong.Id)
								newEwSong.Copyright = ewDatabaseLink.Variation.Copyright.Name
								tx.Model(&ewDatabaseLink).
									Update("copyright", ewDatabaseLink.Variation.Copyright.Name)
							} else {
								tx.Model(&ewDatabaseLink).
									Update("copyright", ewSong.Copyright)
								tx.Model(&ewDatabaseLink.Variation.Copyright).
									Update("name", ewSong.Copyright)
							}
						}
					} else {
						newCopyright := db.CreateCopyrightByName(
							tx,
							ewSong.Copyright,
						)

						tx.Model(&ewDatabaseLink.Variation).
							Update("copyright_id", newCopyright.ID)

						if ewDatabaseLink.Copyright != ewSong.Copyright {
							tx.Model(&ewDatabaseLink).
								Update("copyright", ewSong.Copyright)
						}
					}
				} else {
					if ewDatabaseLink.Variation.Copyright.ID > 0 {
						newEwSong := res.CreateOrGetEwSong(ewSong.Id)
						newEwSong.Copyright = ewDatabaseLink.Variation.Copyright.Name
					}
				}

				newestVariationVersion := ewDatabaseLink.Variation.FindNewestVersion()
				if newestVariationVersion.ID > 0 {
					if newestVariationVersion.Version == ewDatabaseLink.Version {
						if newestVariationVersion.DisabledAt == nil &&
							ewDatabase.SongDatabase.
								HasVariation(newestVariationVersion.VariationID) == true {
							if ewSong.Title != newestVariationVersion.Name ||
								ewSong.Text != newestVariationVersion.Text {

								newVariationVersion := help.NewVariationVersionFromEwSong(
									ewSong,
									newestVariationVersion.VariationID,
									newestVariationVersion.Version+1,
								)
								tx.Model(&newestVariationVersion).
									Update("disabled_at", time.Now())

								srNewPassivatedVariationVersions = append(
									srNewPassivatedVariationVersions,
									db.SrPassivatedVariationVersion{
										SrID:               synchronizationRaport.ID,
										VariationVersionID: newestVariationVersion.ID,
									},
								)

								createNewVariationVersions = append(
									createNewVariationVersions,
									*newVariationVersion,
								)
							}
						} else {
							if ewDatabase.RemoveSongsFromEwDatabase == true {
								res.RemoveEwSongIds = append(
									res.RemoveEwSongIds,
									ewDatabaseLink.EwDatabaseSongID,
								)
								removeEwDatabaseLinks = append(
									removeEwDatabaseLinks,
									ewDatabaseLink.ID,
								)

								srEwSongs = append(
									srEwSongs,
									db.SrEwSong{
										SrID:               synchronizationRaport.ID,
										VariationVersionID: newestVariationVersion.ID,
										Operation:          false,
									},
								)
								srEwDatabaseLinks = append(
									srEwDatabaseLinks,
									db.SrEwDatabaseLink{
										SrID:             synchronizationRaport.ID,
										EwDatabaseID:     ewDatabaseLink.EwDatabaseID,
										EwDatabaseSongID: ewDatabaseLink.EwDatabaseSongID,
										VariationID:      ewDatabaseLink.VariationID,
										Version:          ewDatabaseLink.Version,
										Author:           ewDatabaseLink.Author,
										Copyright:        ewDatabaseLink.Copyright,
										Operation:        false,
									},
								)
							}
						}
					} else {
						if newestVariationVersion.DisabledAt == nil {
							ewLinkVariationVersion := ewDatabaseLink.Variation.
								FindVersionWithVersionNumber(ewDatabaseLink.Version)

							if ewSong.Title != ewLinkVariationVersion.Name ||
								ewSong.Text != ewLinkVariationVersion.Text {

								switch ewDatabase.VariationVersionConflictAction {
								case 1: // Use ew version
									newVariationVersion := help.NewVariationVersionFromEwSong(
										ewSong,
										newestVariationVersion.VariationID,
										newestVariationVersion.Version,
									)
									createNewVariationVersions = append(
										createNewVariationVersions,
										*newVariationVersion,
									)

								case 2: // Use database
									tx.Model(&ewDatabaseLink).
										Update(
											"version",
											newestVariationVersion.Version,
										)

									ewSong := res.CreateOrGetEwSong(ewSong.Id)

									ewSong.Title = newestVariationVersion.Name
									ewSong.Text = newestVariationVersion.Text
								case 3: // Report conflict
									srEwConflicts = append(
										srEwConflicts,
										db.SrEwConflict{
											VariationVersionID: newestVariationVersion.ID,
											EwDatabaseID:       ewDatabase.ID,
											EwSongID:           ewSong.Id,
											Name:               ewSong.Title,
											Text:               ewSong.Text,
											Resolved:           0,
										},
									)
								case 4: // Create branch
									newVariation, newVariationVersion := help.NewVariationVersionFromVariationVersion(
										tx,
										&newestVariationVersion,
									)

									if ewSong.Author != "" {
										newAuthor := db.CreateAuthorByName(
											tx,
											ewSong.Author,
										)
										tx.Model(&newVariation).
											Update("author_id", newAuthor.ID)
									}

									if ewSong.Copyright != "" {
										newCopyright := db.CreateCopyrightByName(
											tx,
											ewSong.Copyright,
										)
										tx.Model(&newCopyright).
											Update("copyright_id", newCopyright.ID)
									}

									createBranches = append(
										createBranches,
										db.Branch{
											SourceVariationVersionID:      newestVariationVersion.ID,
											DestinationVariationVersionID: newVariationVersion.ID,
										},
									)

									addVariationsToSongDatabase = append(
										addVariationsToSongDatabase,
										db.SongDatabaseVariation{
											SongDatabaseID: ewDatabase.SongDatabaseID,
											VariationID:    newVariation.ID,
										},
									)
									addVariationsToEwDatabase = append(
										addVariationsToEwDatabase,
										db.EwDatabaseLink{
											EwDatabaseID:     ewDatabase.ID,
											EwDatabaseSongID: ewSong.Id,
											VariationID:      newVariation.ID,
											Version:          newVariationVersion.Version,
										},
									)
								default:
									tx.Model(&ewDatabaseLink).
										Update(
											"version",
											newestVariationVersion.Version,
										)

									ewSong := res.CreateOrGetEwSong(ewSong.Id)

									ewSong.Title = newestVariationVersion.Name
									ewSong.Text = newestVariationVersion.Text
								}
							} else {
								tx.Model(&ewDatabaseLink).
									Update(
										"version",
										newestVariationVersion.Version,
									)

								ewSong := res.CreateOrGetEwSong(ewSong.Id)

								ewSong.Title = newestVariationVersion.Name
								ewSong.Text = newestVariationVersion.Text
							}
						} else {
							if ewDatabase.RemoveSongsFromEwDatabase == true {
								res.RemoveEwSongIds = append(
									res.RemoveEwSongIds,
									ewDatabaseLink.EwDatabaseSongID,
								)
								removeEwDatabaseLinks = append(
									removeEwDatabaseLinks,
									ewDatabaseLink.ID,
								)
							}
						}
					}
				}
			}
		}

	}

	var excludeVariationIds []uint32

	for _, variation := range ewDatabase.SongDatabase.Variations {
		excludeVariationIds = append(
			excludeVariationIds,
			variation.ID,
		)

		if is, link := ewDatabase.HasVariation(variation.ID); is == true {
			if ewSongs[link.EwDatabaseSongID] == nil {
				if ewDatabase.RemoveSongsFromSongDatabase == true {
					removeSongDatabaseVariationLinks = append(
						removeSongDatabaseVariationLinks,
						variation.ID,
					)
					removeEwDatabaseLinks = append(
						removeEwDatabaseLinks,
						link.ID,
					)
					srRemoveSongDatabaseVariations = append(
						srRemoveSongDatabaseVariations,
						db.SrRemoveSongDatabaseVariation{
							SrID:           synchronizationRaport.ID,
							VariationID:    variation.ID,
							SongDatabaseID: ewDatabase.SongDatabaseID,
						},
					)
					srEwDatabaseLinks = append(
						srEwDatabaseLinks,
						db.SrEwDatabaseLink{
							SrID:             synchronizationRaport.ID,
							EwDatabaseID:     ewDatabase.ID,
							EwDatabaseSongID: link.EwDatabaseSongID,
							VariationID:      variation.ID,
							Version:          link.Version,
							Author:           link.Author,
							Copyright:        link.Copyright,
							Operation:        false,
						},
					)
				} else {
					newestVersion := variation.FindNewestVersion()
					if newestVersion.ID > 0 {
						removeEwDatabaseLinks = append(
							removeEwDatabaseLinks,
							link.ID,
						)

						newEwSong := &MatiasService.EwSong{
							VariationId: variation.ID,
							Title:       newestVersion.Name,
							Text:        newestVersion.Text,
						}

						if variation.Author.ID > 0 {
							newEwSong.Author = variation.Author.Name
						}

						if variation.Copyright.ID > 0 {
							newEwSong.Copyright = variation.Copyright.Name
						}

						res.EwSongs = append(
							res.EwSongs,
							newEwSong,
						)
					}
				}
			}
		} else {
			newestVersion := variation.FindNewestVersion()
			if hasEwSongByNameAndText(
				ewSongs,
				newestVersion.Name,
				newestVersion.Text,
			) == false {
				if newestVersion.ID > 0 {
					newEwSong := &MatiasService.EwSong{
						VariationId: variation.ID,
						Title:       newestVersion.Name,
						Text:        newestVersion.Text,
					}

					if variation.Author.ID > 0 {
						newEwSong.Author = variation.Author.Name
					}

					if variation.Copyright.ID > 0 {
						newEwSong.Copyright = variation.Copyright.Name
					}

					res.EwSongs = append(
						res.EwSongs,
						newEwSong,
					)
				}
			}
		}
	}

	var songDatabaseTagVariations []db.Variation
	q := tx.Table("variations").
		Preload("Author").
		Preload("Copyright").
		Preload("VariationVersions").
		Joins("left join tag_variations on tag_variations.variation_id = variations.id").
		Joins("left join song_database_tags on song_database_tags.tag_id = tag_variations.tag_id").
		Where("song_database_tags.song_database_id = ?", ewDatabase.SongDatabaseID).
		Select("distinct variations.*")
	if len(excludeVariationIds) > 0 {
		q = q.Where("variations.id not in(?)", excludeVariationIds)
	}
	q.Find(&songDatabaseTagVariations)

	for _, variation := range songDatabaseTagVariations {
		if is, link := ewDatabase.HasVariation(variation.ID); is == true {
			if ewSongs[link.EwDatabaseSongID] == nil {
				newestVersion := variation.FindNewestVersion()
				if newestVersion.ID > 0 {
					removeEwDatabaseLinks = append(
						removeEwDatabaseLinks,
						link.ID,
					)
					newEwSong := &MatiasService.EwSong{
						VariationId: variation.ID,
						Title:       newestVersion.Name,
						Text:        newestVersion.Text,
					}

					if variation.Author.ID > 0 {
						newEwSong.Author = variation.Author.Name
					}

					if variation.Copyright.ID > 0 {
						newEwSong.Copyright = variation.Copyright.Name
					}

					res.EwSongs = append(
						res.EwSongs,
						newEwSong,
					)
				}
			}
		} else {
			newestVersion := variation.FindNewestVersion()
			if hasEwSongByNameAndText(
				ewSongs,
				newestVersion.Name,
				newestVersion.Text,
			) == false {
				if newestVersion.ID > 0 {
					newEwSong := &MatiasService.EwSong{
						VariationId: variation.ID,
						Title:       newestVersion.Name,
						Text:        newestVersion.Text,
					}

					if variation.Author.ID > 0 {
						newEwSong.Author = variation.Author.Name
					}

					if variation.Copyright.ID > 0 {
						newEwSong.Copyright = variation.Copyright.Name
					}

					res.EwSongs = append(
						res.EwSongs,
						newEwSong,
					)
				}
			}
		}
	}

	var affectedRows int64

	affectedRows = managers.BatchCreateVariationVersions(
		tx,
		createNewVariationVersions,
	)
	var newVariationVersions []db.VariationVersion
	tx.Limit(affectedRows).Find(&newVariationVersions)
	for _, newVariationVersion := range newVariationVersions {
		srNewVariationVersions = append(
			srNewVariationVersions,
			db.SrNewVariationVersion{
				SrID:               synchronizationRaport.ID,
				VariationVersionID: newVariationVersion.ID,
			},
		)
	}

	affectedRows = managers.BatchAddVariationsToSongDatabase(
		tx,
		addVariationsToSongDatabase,
	)
	var newSongDatabaseVariations []db.SongDatabaseVariation
	tx.Limit(affectedRows).Find(&newSongDatabaseVariations)
	for _, songDatabaseVariation := range newSongDatabaseVariations {
		srAddSongDatabaseVariations = append(
			srAddSongDatabaseVariations,
			db.SrAddSongDatabaseVariation{
				SrID:           synchronizationRaport.ID,
				VariationID:    songDatabaseVariation.VariationID,
				SongDatabaseID: songDatabaseVariation.SongDatabaseID,
			},
		)
	}

	affectedRows = managers.BatchAddVariationsToEwDatabase(
		tx,
		addVariationsToEwDatabase,
	)

	// var addedVariationsToEwDatabase []db.EwDatabaseLink
	// tx.Limit(affectedRows).First(&addedVariationsToEwDatabase)
	// for _, ewDatabaseLink := range addedVariationsToEwDatabase {
	// 	synchronizationRaport.NewEwDatabaseLinks = append(
	// 		synchronizationRaport.NewEwDatabaseLinks,
	// 		db.SrEwDatabaseLink{
	// 			EwDatabaseID:     ewDatabaseLink.EwDatabaseID,
	// 			EwDatabaseSongID: ewDatabaseLink.EwDatabaseSongID,
	// 			VariationID:      ewDatabaseLink.VariationID,
	// 			Version:          ewDatabaseLink.Version,
	// 			Author:           ewDatabaseLink.Author,
	// 			Copyright:        ewDatabaseLink.Copyright,
	// 			Operation:        true,
	// 		},
	// 	)
	// }

	affectedRows = managers.BatchCreateBranches(
		tx,
		createBranches,
	)
	var newBranches []db.SrNewBranch
	tx.Limit(affectedRows).First(&newBranches)
	for _, newBranch := range newBranches {
		srNewBranches = append(
			srNewBranches,
			db.SrNewBranch{
				SrID:     synchronizationRaport.ID,
				BranchID: newBranch.ID,
			},
		)
	}

	if len(removeEwDatabaseLinks) > 0 {
		tx.Where("id in (?)", removeEwDatabaseLinks).Delete(&db.EwDatabaseLink{})
	}

	if len(removeSongDatabaseVariationLinks) > 0 {
		tx.Where("variation_id in (?)", removeSongDatabaseVariationLinks).
			Where("song_database_id = ?", ewDatabase.SongDatabaseID).
			Delete(&db.SongDatabaseVariation{})
	}

	endDate := time.Now()
	endTime := endDate.UnixNano() / int64(time.Millisecond)
	synchronizationRaport.DurationMS = endTime - startTime
	synchronizationRaport.FinishedAt = &endDate

	// Raport generation
	managers.BatchCreateSrAddSongDatabaseVariations(
		tx,
		srAddSongDatabaseVariations,
	)
	managers.BatchCreateSrEwConflicts(
		tx,
		srEwConflicts,
	)
	managers.BatchCreateSrEwDatabaseLinks(
		tx,
		srEwDatabaseLinks,
	)
	managers.BatchCreateSrEwSongs(
		tx,
		srEwSongs,
	)
	managers.BatchCreateSrNewAuthors(
		tx,
		srNewAuthors,
	)
	managers.BatchCreateSrNewBranches(
		tx,
		srNewBranches,
	)
	managers.BatchCreateSrNewCopyrights(
		tx,
		srNewCopyrights,
	)
	managers.BatchCreateSrNewVariations(
		tx,
		srNewVariations,
	)
	managers.BatchCreateSrNewVariationVersions(
		tx,
		srNewVariationVersions,
	)

	managers.BatchCreateSrRemoveSongDatabaseVariations(
		tx,
		srRemoveSongDatabaseVariations,
	)

	tx.Save(&synchronizationRaport)

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
