package services

import (
	"database/sql"
	"time"

	null "gopkg.in/volatiletech/null.v6"

	"github.com/koodinikkarit/seppo/managers"
	"golang.org/x/net/context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/matias_service"
	"github.com/koodinikkarit/seppo/models"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

func insertVariationLinks(
	tx *sql.Tx,
	synchronizationRaport *models.SynchronizationRaport,
	ewDatabase *models.EwDatabase,
	variation *models.Variation,
	ewSong *MatiasService.EwSong,
) {
	newSongDatabaseVariation := models.SongDatabaseVariation{
		SongDatabaseID: ewDatabase.SongDatabaseID,
		VariationID:    variation.ID,
	}
	newSongDatabaseVariation.Insert(tx)
	synchronizationRaport.AddSRSRAddSongDatabaseVariations(
		tx,
		true,
		generators.NewSrSongDatabaseVariation(
			&newSongDatabaseVariation,
		),
	)
	ewDatabaseLink := models.EwDatabaseLink{
		EwDatabaseID:     ewDatabase.ID,
		EwDatabaseSongID: uint64(ewSong.Id),
		VariationID:      variation.ID,
		Version:          1,
		Author:           ewSong.Author,
		Copyright:        ewSong.Copyright,
	}
	ewDatabaseLink.Insert(tx)
	synchronizationRaport.AddSRSREwDatabaseLinks(
		tx,
		true,
		generators.NewSrEwDatabaseLinkFromEwDatabaseLink(
			&ewDatabaseLink,
			1,
		),
	)
}

func removeVariationLinks(
	tx *sql.Tx,
	synchronizationRaport *models.SynchronizationRaport,
	ewDatabase *models.EwDatabase,
	variation *models.Variation,
) {
	models.SongDatabaseVariations(
		tx,
		Where("song_database_id = ?", ewDatabase.SongDatabaseID),
		Where("variation_id = ?", variation.ID),
	).DeleteAll()
	srSongDatabaseVariation := models.SRRemoveSongDatabaseVariation{
		VariationID:    variation.ID,
		SongDatabaseID: ewDatabase.SongDatabaseID,
	}
	srSongDatabaseVariation.Insert(tx)
	ewDatabaseLink, _ := models.EwDatabaseLinks(
		tx,
		Where("ew_database_id = ?", ewDatabase.ID),
		Where("variation_id = ?", variation.ID),
	).One()
	srEwDatabaseLink := models.SREwDatabaseLink{
		EwDatabaseID:     ewDatabase.ID,
		EwDatabaseSongID: ewDatabaseLink.EwDatabaseSongID,
		VariationID:      ewDatabaseLink.VariationID,
		Version:          ewDatabaseLink.Version,
		Author:           ewDatabaseLink.Author,
		Copyright:        ewDatabaseLink.Copyright,
		Operation:        0,
	}
	ewDatabaseLink.Delete(tx)
	synchronizationRaport.AddSRSREwDatabaseLinks(
		tx,
		true,
		&srEwDatabaseLink,
	)
}

func createBranch(
	tx *sql.Tx,
	synchronizationRaport *models.SynchronizationRaport,
	ewSong *MatiasService.EwSong,
	newestVariationVersion *models.VariationVersion,
	ewDatabase *models.EwDatabase,
) {
	newVariation, newVariationVersion := managers.NewVariationFromEwSong(
		tx,
		ewSong,
	)
	synchronizationRaport.AddSRSRNewVariations(
		tx,
		true,
		generators.NewSrVariationFromVariation(newVariation),
	)
	synchronizationRaport.AddSRSRNewVariationVersions(
		tx,
		true,
		generators.NewSrVariationVersionFromVariationVersion(
			newVariationVersion,
		),
	)
	newBranch := models.Branch{
		SourceVariationVersionID:      newestVariationVersion.ID,
		DestinationVariationVersionID: newVariationVersion.ID,
	}
	newBranch.Insert(tx)
	synchronizationRaport.AddSRSRNewBranches(
		tx,
		true,
		generators.NewSrBranchFromBranch(
			&newBranch,
		),
	)
	insertVariationLinks(
		tx,
		synchronizationRaport,
		ewDatabase,
		newVariation,
		ewSong,
	)
}

func ewSyncRemoveSongsFromEwDatabase(
	tx *sql.Tx,
	synchronizationRaport *models.SynchronizationRaport,
	ewDatabase *models.EwDatabase,
	res *MatiasService.SyncEwDatabaseResponse,
	ewDatabaseLink *models.EwDatabaseLink,
	ewSong *MatiasService.EwSong,
) {
	if ewDatabase.RemoveSongsFromEwDatabase == null.NewInt8(1, true) {
		res.RemoveEwSongIds = append(
			res.RemoveEwSongIds,
			ewSong.Id,
		)
		// srEwSong := *models.SREwSong{

		// }
		srEwDatabaseLink := models.SREwDatabaseLink{
			EwDatabaseID:     ewDatabaseLink.EwDatabaseID,
			EwDatabaseSongID: ewDatabaseLink.EwDatabaseSongID,
			VariationID:      ewDatabaseLink.VariationID,
			Version:          ewDatabaseLink.Version,
			Author:           ewDatabaseLink.Author,
			Copyright:        ewDatabaseLink.Copyright,
			Operation:        0,
		}
		srEwDatabaseLink.Insert(tx)
		synchronizationRaport.AddSRSREwDatabaseLinks(
			tx,
			true,
			&srEwDatabaseLink,
		)
	}
}

func ewSyncNewVariationVersionFromEwSong(
	tx *sql.Tx,
	synchronizationRaport *models.SynchronizationRaport,
	variation *models.Variation,
	ewSong *MatiasService.EwSong,
	version uint,
) {
	models.VariationVersions(
		tx,
		Where("variation_versions.variation_id = ?", variation.ID),
	).UpdateAll(
		models.M{
			"disabled_at": time.Now(),
		},
	)
	newVariationVersion := models.VariationVersion{
		Name:    ewSong.Title,
		Text:    ewSong.Text,
		Version: version,
	}
	variation.AddVariationVersions(
		tx,
		true,
		&newVariationVersion,
	)
	synchronizationRaport.AddSRSRPassivatedVariationVersions(
		tx,
		true,
		generators.NewSrPassivatedVariationVersionFromVariationVersion(
			&newVariationVersion,
		),
	)
}

func setEwSyncResponseNewEwSong(
	tx *sql.Tx,
	synchronizationRaport *models.SynchronizationRaport,
	res *MatiasService.SyncEwDatabaseResponse,
	variationVersion *models.VariationVersion,
	ewDatabaseLink *models.EwDatabaseLink,
	author string,
	copyright string,
) {
	res.EwSongs = append(
		res.EwSongs,
		generators.NewEwSongFromVariationVersion(
			uint32(ewDatabaseLink.EwDatabaseSongID),
			variationVersion,
		),
	)
	updateEwDatabaseLinkVersion(
		tx,
		synchronizationRaport,
		ewDatabaseLink,
		variationVersion.Version,
	)
}

func updateEwDatabaseLinkVersion(
	tx *sql.Tx,
	synchronizationRaport *models.SynchronizationRaport,
	ewDatabaseLink *models.EwDatabaseLink,
	version uint,
) {
	srUpdateEwDatabaseLinkVersion := models.SRUpdatedEwDatabaseLinkVersion{
		OrigVersion:    ewDatabaseLink.Version,
		ChangedVersion: version,
		VariationID:    ewDatabaseLink.VariationID,
	}
	synchronizationRaport.AddSRSRUpdatedEwDatabaseLinkVersions(
		tx,
		true,
		&srUpdateEwDatabaseLinkVersion,
	)
	ewDatabaseLink.Version = version
	ewDatabaseLink.Update(
		tx,
		"version",
	)
}

func ewSyncUpdateAuthor(
	tx *sql.Tx,
	ewDatabaseLink *models.EwDatabaseLink,
	variation *models.Variation,
	newEwSong *MatiasService.EwSong,
	authorName string,
) {
	author, _ := variation.Author(tx).One()
	if authorName == "" {
		if author != nil {
			newEwSong.Author = author.Name
		}
		return
	}

	if author == nil {
		newAuthor := managers.CreateAuthorByName(
			tx,
			authorName,
		)
		variation.AuthorID = null.NewUint64(newAuthor.ID, true)
		variation.Update(tx, "author_id")
		if ewDatabaseLink.Author != authorName {
			ewDatabaseLink.Author = authorName
			ewDatabaseLink.Update(tx, "author")
		}
		return
	}

	if ewDatabaseLink.Author == authorName {
		if author.Name == authorName {
			return
		}

		newEwSong.Author = author.Name
		ewDatabaseLink.Author = author.Name
		ewDatabaseLink.Update(tx, "author")
		return
	}

	if author.Name == ewDatabaseLink.Author {
		newEwSong.Author = author.Name
		ewDatabaseLink.Author = author.Name
		ewDatabaseLink.Update(tx, "author")
	} else {
		ewDatabaseLink.Author = authorName
		ewDatabaseLink.Update(tx, "author")
		author.Name = authorName
		author.Update(tx, "name")
	}
}

func ewSyncUpdateCopyright(
	tx *sql.Tx,
	ewDatabaseLink *models.EwDatabaseLink,
	variation *models.Variation,
	newEwSong *MatiasService.EwSong,
	copyrightName string,
) {
	copyright, _ := variation.Copyright(tx).One()

	if copyrightName == "" {
		if copyright != nil {
			newEwSong.Copyright = copyright.Name
		}
		return
	}

	if copyright == nil {
		newCopyright := managers.CreateCopyrightByName(
			tx,
			copyrightName,
		)
		variation.CopyrightID = null.NewUint64(newCopyright.ID, true)
		variation.Update(tx, "copyright_id")
		if ewDatabaseLink.Copyright != copyrightName {
			ewDatabaseLink.Copyright = copyrightName
			ewDatabaseLink.Update(tx, "copyright")
		}
		return
	}

	if ewDatabaseLink.Copyright == copyrightName {
		if copyright.Name == copyrightName {
			return
		}

		newEwSong.Copyright = copyright.Name
		ewDatabaseLink.Copyright = copyright.Name
		ewDatabaseLink.Update(tx, "copyright")
		return
	}

	if copyright.Name == ewDatabaseLink.Copyright {
		newEwSong.Copyright = copyright.Name
		ewDatabaseLink.Copyright = copyright.Name
		ewDatabaseLink.Update(tx, "copyright")
	} else {
		ewDatabaseLink.Copyright = copyrightName
		ewDatabaseLink.Update(tx, "copyright")
		copyright.Name = copyrightName
		copyright.Update(tx, "name")
	}
}

// Process one ewSync ewSong
func ewSyncProcessEwSong(
	tx *sql.Tx,
	res *MatiasService.SyncEwDatabaseResponse,
	ewDatabase *models.EwDatabase,
	synchronizationRaport *models.SynchronizationRaport,
	ewSong *MatiasService.EwSong,
) {
	ewDatabaseLink, _ := ewDatabase.EwDatabaseLinks(
		tx,
		Load("Variation"),
		Load("Variation.VariationVersions"),
		Where(
			"ew_database_links.ew_database_song_id = ?",
			ewSong.Id,
		),
	).One()

	if ewDatabaseLink != nil {
		newEwSong := &MatiasService.EwSong{
			Id: ewSong.Id,
		}
		variation, _ := ewDatabaseLink.Variation(tx).One()
		if variation == nil {
			ewSyncRemoveSongsFromEwDatabase(
				tx,
				synchronizationRaport,
				ewDatabase,
				res,
				ewDatabaseLink,
				ewSong,
			)
			return
		}

		ewSyncUpdateAuthor(
			tx,
			ewDatabaseLink,
			variation,
			newEwSong,
			ewSong.Author,
		)

		ewSyncUpdateCopyright(
			tx,
			ewDatabaseLink,
			variation,
			newEwSong,
			ewSong.Copyright,
		)

		variationVersions, _ := variation.VariationVersions(tx).All()
		variationVersion := managers.FindVariationVersionByVersion(
			variationVersions,
			ewDatabaseLink.Version,
		)
		if variationVersion == nil {
			ewSyncRemoveSongsFromEwDatabase(
				tx,
				synchronizationRaport,
				ewDatabase,
				res,
				ewDatabaseLink,
				ewSong,
			)
			return
		}

		newestVariationVersion := managers.FindNewestVariationVersion(
			variationVersions,
		)

		if newestVariationVersion.ID == variationVersion.ID {
			if newestVariationVersion.DisabledAt.Valid == true {
				ewSyncRemoveSongsFromEwDatabase(
					tx,
					synchronizationRaport,
					ewDatabase,
					res,
					ewDatabaseLink,
					ewSong,
				)
				return
			}

			if ewSong.Title == newestVariationVersion.Name &&
				ewSong.Text == newestVariationVersion.Text {
				return
			}

			ewSyncNewVariationVersionFromEwSong(
				tx,
				synchronizationRaport,
				variation,
				ewSong,
				newestVariationVersion.Version+1,
			)
			return
		}

		if newestVariationVersion.DisabledAt.Valid == true {
			ewSyncRemoveSongsFromEwDatabase(
				tx,
				synchronizationRaport,
				ewDatabase,
				res,
				ewDatabaseLink,
				ewSong,
			)
			return
		}

		if ewSong.Title == variationVersion.Name &&
			ewSong.Text == variationVersion.Text {
			setEwSyncResponseNewEwSong(
				tx,
				synchronizationRaport,
				res,
				newestVariationVersion,
				ewDatabaseLink,
				"",
				"",
			)
			return
		}

		switch ewDatabase.VariationVersionConflictAction {
		case 1: // Use ew version
			ewSyncNewVariationVersionFromEwSong(
				tx,
				synchronizationRaport,
				variation,
				ewSong,
				newestVariationVersion.Version+1,
			)
		case 2: // Use database version
			setEwSyncResponseNewEwSong(
				tx,
				synchronizationRaport,
				res,
				newestVariationVersion,
				ewDatabaseLink,
				"",
				"",
			)
		case 3: // Report conflict
			srNewConflict := models.SREwConflict{
				VariationVersionID: newestVariationVersion.ID,
				EwDatabaseID:       ewDatabase.ID,
				EwSongID:           ewDatabaseLink.EwDatabaseSongID,
				Name:               null.NewString(ewSong.Title, true),
				Text:               null.NewString(ewSong.Text, true),
			}
			synchronizationRaport.AddSRSREwConflicts(
				tx,
				true,
				&srNewConflict,
			)
		case 4: // Create branch
			createBranch(
				tx,
				synchronizationRaport,
				ewSong,
				newestVariationVersion,
				ewDatabase,
			)
		default:
			setEwSyncResponseNewEwSong(
				tx,
				synchronizationRaport,
				res,
				newestVariationVersion,
				ewDatabaseLink,
				"",
				"",
			)
		}

		if newEwSong.Title != "" ||
			newEwSong.Text != "" ||
			newEwSong.Author != "" ||
			newEwSong.Copyright != "" {

			res.EwSongs = append(
				res.EwSongs,
				newEwSong,
			)
		}

		return
	}

	sameVariationVersion, _ := models.VariationVersions(
		tx,
		Load("Variation"),
		Load("Variation.Author"),
		Load("Variation.Copyright"),
		Load("Variation.VariationVersions"),
		Where("variation_versions.name = ?", ewSong.Title),
		Where("variation_versions.text = ?", ewSong.Text),
	).One()

	// Check if there is not variationversion
	// with same name and text then create new
	if sameVariationVersion == nil {
		newVariation, newVariationVersion := managers.NewVariationFromEwSong(
			tx,
			ewSong,
		)
		synchronizationRaport.AddSRSRNewVariations(
			tx,
			true,
			generators.NewSrVariationFromVariation(newVariation),
		)
		synchronizationRaport.AddSRSRNewVariationVersions(
			tx,
			true,
			generators.NewSrVariationVersionFromVariationVersion(
				newVariationVersion,
			),
		)
		insertVariationLinks(
			tx,
			synchronizationRaport,
			ewDatabase,
			newVariation,
			ewSong,
		)
		return
	}

	// First find newest variationversion and then check if
	// found is same else create new branch from it.
	variation, _ := sameVariationVersion.Variation(tx).One()
	variationVersions, _ := variation.VariationVersions(tx).All()
	newestVariationVersion := managers.FindNewestVariationVersion(variationVersions)

	if sameVariationVersion.ID == newestVariationVersion.ID {
		// Tarkistetaan on uusin variationversion disabloitu
		// jos on tarvitsee siitä luoda uusi branch.
		if sameVariationVersion.DisabledAt.Valid == true {
			createBranch(
				tx,
				synchronizationRaport,
				ewSong,
				newestVariationVersion,
				ewDatabase,
			)
			return
		}
		if found, _ := models.SongDatabaseTags(
			tx,
			InnerJoin("tag_variations tv on tv.tag_id = song_database_tags.tag_id"),
			Where("song_database_tags.song_database_id = ?", ewDatabase.SongDatabaseID),
			Where("tag_variations.variation_id = ?", variation.ID),
		).Exists(); found == false {
			insertVariationLinks(
				tx,
				synchronizationRaport,
				ewDatabase,
				variation,
				ewSong,
			)
			return
		}
		return
	}

	createBranch(
		tx,
		synchronizationRaport,
		ewSong,
		newestVariationVersion,
		ewDatabase,
	)
}

func ewSyncCreateEwSong(
	tx *sql.Tx,
	res *MatiasService.SyncEwDatabaseResponse,
	synchronizationRaport *models.SynchronizationRaport,
	variation *models.Variation,
	variationVersion *models.VariationVersion,
) {
	ewSong := MatiasService.EwSong{
		Title:       variationVersion.Name,
		Text:        variationVersion.Text,
		VariationId: uint32(variation.ID),
	}

	author := variation.Author(tx).OneP()
	if author != nil {
		ewSong.Author = author.Name
	}
	copyright := variation.Copyright(tx).OneP()
	if copyright != nil {
		ewSong.Copyright = copyright.Name
	}
	res.EwSongs = append(
		res.EwSongs,
		&ewSong,
	)
}

func ewSyncProcessSongDatabaseVariation(
	tx *sql.Tx,
	res *MatiasService.SyncEwDatabaseResponse,
	ewSongs map[uint32]*MatiasService.EwSong,
	ewDatabase *models.EwDatabase,
	songDatabaseVariation *models.SongDatabaseVariation,
	synchronizationRaport *models.SynchronizationRaport,
) {
	variation := songDatabaseVariation.Variation(tx).OneP()
	variationVersions := variation.VariationVersions(tx).AllP()
	newestVariationVersion := managers.FindNewestVariationVersion(
		variationVersions,
	)

	ewDatabaseLink, _ := ewDatabase.EwDatabaseLinks(
		tx,
		Load("Variation"),
		Load("Variation.VariationVersions"),
		Where(
			"ew_database_links.variation_id = ?",
			songDatabaseVariation.VariationID,
		),
	).One()

	if ewDatabaseLink == nil {
		ewSong := managers.FindEwSongWithNameText(
			ewSongs,
			newestVariationVersion.Name,
			newestVariationVersion.Text,
		)

		if ewSong != nil {
			return
		}

		ewSyncCreateEwSong(
			tx,
			res,
			synchronizationRaport,
			variation,
			newestVariationVersion,
		)

		return
	}

	if ewSongs[uint32(ewDatabaseLink.EwDatabaseSongID)] != nil {
		return
	}

	ewSyncRemoveSongsFromEwDatabase(
		tx,
		synchronizationRaport,
		ewDatabase,
		res,
		ewDatabaseLink,
		ewSongs[uint32(ewDatabaseLink.EwDatabaseSongID)],
	)

}

func (s *MatiasServiceServer) SyncEwDatabase(
	ctx context.Context,
	in *MatiasService.SyncEwDatabaseRequest,
) (
	*MatiasService.SyncEwDatabaseResponse,
	error,
) {
	res := &MatiasService.SyncEwDatabaseResponse{}

	db, _ := sql.Open(
		"mysql",
		"root:asdf321@tcp(localhost:3306)/seppo2?parseTime=True&loc=Local",
	)
	defer db.Close()
	tx, _ := db.Begin()

	ewSongs := make(map[uint32]*MatiasService.EwSong)

	synchronizationRaport := models.SynchronizationRaport{}
	synchronizationRaport.Insert(db)

	ewDatabase, _ := models.EwDatabases(
		tx,
		Load("SongDatabase"),
		Load("SongDatabase.SongDatabaseVariations"),
		Load("SongDatabase.SongDatabaseVariations.Variation"),
		Load("SongDatabase.SongDatabaseVariations.Variation.VariationVersions"),
		Load("EwDatabaseLinks"),
		WhereIn("ew_databases.ew_database_key = ?", in.EwDatabaseKey),
	).One()

	if ewDatabase.ID == 0 {
		return res, nil
	}

	ewSongs, dublicates := managers.RemoveDuplicatesFromEwSongs(in.EwSongs)
	res.RemoveEwSongIds = append(
		res.RemoveEwSongIds,
		dublicates...,
	)

	for _, ewSong := range ewSongs {
		ewSyncProcessEwSong(
			tx,
			res,
			ewDatabase,
			&synchronizationRaport,
			ewSong,
		)
	}

	songDatabase, _ := ewDatabase.SongDatabase(tx).One()
	if songDatabase == nil {
		return res, nil
	}

	songDatabaseVariations := songDatabase.SongDatabaseVariations(tx).AllP()
	for _, songDatabaseVariation := range songDatabaseVariations {
		ewSyncProcessSongDatabaseVariation(
			tx,
			res,
			ewSongs,
			ewDatabase,
			songDatabaseVariation,
			&synchronizationRaport,
		)
	}

	tx.Commit()

	return res, nil
}

// func prosessNoSameVariation(
// 	tx *gorm.DB,
// ) {
// 	// No variation with same name and text
// 	newVariation := managers.NewVariationFromEwSong(
// 		tx,
// 		ewSong,
// 	)
// 	// Add variation and version to list for raport generation
// 	srNewVariations = append(
// 		srNewVariations,
// 		db.SrNewVariation{
// 			SrID:        synchronizationRaport.ID,
// 			VariationID: newVariation.ID,
// 		},
// 	)
// 	srNewVariationVersions = append(
// 		srNewVariationVersions,
// 		db.SrNewVariationVersion{
// 			SrID:               synchronizationRaport.ID,
// 			VariationVersionID: newVariation.VariationVersions[0].ID,
// 		},
// 	)

// 	// These are added to slice for batch insert
// 	addVariationsToSongDatabase = append(
// 		addVariationsToSongDatabase,
// 		db.SongDatabaseVariation{
// 			SongDatabaseID: ewDatabase.SongDatabaseID,
// 			VariationID:    newVariation.ID,
// 		},
// 	)
// 	addVariationsToEwDatabase = append(
// 		addVariationsToEwDatabase,
// 		db.EwDatabaseLink{
// 			EwDatabaseID:     ewDatabase.ID,
// 			EwDatabaseSongID: ewSong.Id,
// 			VariationID:      newVariation.ID,
// 			Version:          1,
// 		},
// 	)
// }

// func ewSyncProcessNoDatabaseLink(
// 	tx *gorm.DB,
// ) {
// 	var sameVariation db.Variation
// 	tx.Table("variations").
// 		Joins("left join variation_versions on variations.id = variation_versions.variation_id").
// 		Where("variation_versions.name = ?", ewSong.Title).
// 		Where("variation_versions.text = ?", ewSong.Text).
// 		Preload("VariationVersions").
// 		Preload("SongDatabases").
// 		Preload("TagVariations").
// 		First(&sameVariation)

// 	if sameVariation.ID == 0 {
// 		prosessNoSameVariation(
// 			tx,
// 		)
// 	} else {
// 		newestVariationVersion := sameVariation.FindNewestVersion()
// 		sameVariationVersion := sameVariation.FindVariationVersionByNameAndText(
// 			ewSong.Title,
// 			ewSong.Text,
// 		)
// 		if sameVariationVersion.Version == newestVariationVersion.Version {
// 			if sameVariationVersion.DisabledAt == nil {
// 				if sameVariation.FindSongDatabaseByID(ewDatabase.SongDatabaseID) == nil {
// 					foundTag := false
// 					for _, tagVariation := range sameVariation.TagVariations {
// 						if ewDatabase.SongDatabase.HasSongDatabaseTag(tagVariation.TagID) == true {
// 							foundTag = true
// 							break
// 						}
// 					}
// 					if foundTag == false {

// 						if ewSong.Author != "" &&
// 							sameVariation.AuthorID == nil {

// 							newAuthor := db.CreateAuthorByName(
// 								tx,
// 								ewSong.Author,
// 							)
// 							tx.Model(&sameVariation).
// 								Update("author_id", newAuthor.ID)
// 						}

// 						if ewSong.Copyright != "" &&
// 							sameVariation.CopyrightID == nil {

// 							newCopyright := db.CreateCopyrightByName(
// 								tx,
// 								ewSong.Copyright,
// 							)
// 							tx.Model(&sameVariation).
// 								Update("copyright_id", newCopyright.ID)
// 						}

// 						addVariationsToSongDatabase = append(
// 							addVariationsToSongDatabase,
// 							db.SongDatabaseVariation{
// 								SongDatabaseID: ewDatabase.SongDatabaseID,
// 								VariationID:    sameVariation.ID,
// 							},
// 						)
// 						addVariationsToEwDatabase = append(
// 							addVariationsToEwDatabase,
// 							db.EwDatabaseLink{
// 								EwDatabaseID:     ewDatabase.ID,
// 								EwDatabaseSongID: ewSong.Id,
// 								VariationID:      sameVariation.ID,
// 								Version:          sameVariationVersion.Version,
// 							},
// 						)
// 					}
// 				}
// 			} else {
// 				newVariation := managers.NewVariationFromEwSong(
// 					tx,
// 					ewSong,
// 				)

// 				srNewVariations = append(
// 					srNewVariations,
// 					db.SrNewVariation{
// 						SrID:        synchronizationRaport.ID,
// 						VariationID: newVariation.ID,
// 					},
// 				)
// 				srNewVariationVersions = append(
// 					srNewVariationVersions,
// 					db.SrNewVariationVersion{
// 						SrID:               synchronizationRaport.ID,
// 						VariationVersionID: newVariation.VariationVersions[0].ID,
// 					},
// 				)

// 				createBranches = append(
// 					createBranches,
// 					db.Branch{
// 						SourceVariationVersionID:      sameVariationVersion.ID,
// 						DestinationVariationVersionID: newVariation.VariationVersions[0].ID,
// 					},
// 				)

// 				addVariationsToSongDatabase = append(
// 					addVariationsToSongDatabase,
// 					db.SongDatabaseVariation{
// 						SongDatabaseID: ewDatabase.SongDatabaseID,
// 						VariationID:    newVariation.ID,
// 					},
// 				)
// 				addVariationsToEwDatabase = append(
// 					addVariationsToEwDatabase,
// 					db.EwDatabaseLink{
// 						EwDatabaseID:     ewDatabase.ID,
// 						EwDatabaseSongID: ewSong.Id,
// 						VariationID:      newVariation.ID,
// 						Version:          1,
// 					},
// 				)
// 			}
// 		} else {
// 			newVariation := managers.NewVariationFromEwSong(
// 				tx,
// 				ewSong,
// 			)

// 			srNewVariations = append(
// 				srNewVariations,
// 				db.SrNewVariation{
// 					SrID:        synchronizationRaport.ID,
// 					VariationID: newVariation.ID,
// 				},
// 			)
// 			srNewVariationVersions = append(
// 				srNewVariationVersions,
// 				db.SrNewVariationVersion{
// 					SrID:               synchronizationRaport.ID,
// 					VariationVersionID: newVariation.VariationVersions[0].ID,
// 				},
// 			)

// 			createBranches = append(
// 				createBranches,
// 				db.Branch{
// 					SourceVariationVersionID:      sameVariationVersion.ID,
// 					DestinationVariationVersionID: newVariation.VariationVersions[0].ID,
// 				},
// 			)

// 			addVariationsToSongDatabase = append(
// 				addVariationsToSongDatabase,
// 				db.SongDatabaseVariation{
// 					SongDatabaseID: ewDatabase.SongDatabaseID,
// 					VariationID:    newVariation.ID,
// 				},
// 			)
// 			addVariationsToEwDatabase = append(
// 				addVariationsToEwDatabase,
// 				db.EwDatabaseLink{
// 					EwDatabaseID:     ewDatabase.ID,
// 					EwDatabaseSongID: ewSong.Id,
// 					VariationID:      newVariation.ID,
// 					Version:          1,
// 				},
// 			)
// 		}
// 	}
// }

// func ewSyncProcessUpdateAuthorState(
// 	tx *gorm.DB,
// 	ewSong *MatiasService.EwSong,
// ) {
// 	if ewDatabaseLink.Variation.Author.ID > 0 {
// 		if ewDatabaseLink.Author == ewSong.Author {
// 			if ewDatabaseLink.Variation.Author.Name != ewSong.Author {
// 				newEwSong := res.CreateOrGetEwSong(ewSong.Id)
// 				newEwSong.Author = ewDatabaseLink.Variation.Author.Name
// 				tx.Model(&ewDatabaseLink).
// 					Update("author", ewDatabaseLink.Variation.Author.Name)
// 			}
// 		} else {
// 			if ewDatabaseLink.Variation.Author.Name != ewDatabaseLink.Author {
// 				newEwSong := res.CreateOrGetEwSong(ewSong.Id)
// 				newEwSong.Author = ewDatabaseLink.Variation.Author.Name
// 				tx.Model(&ewDatabaseLink).
// 					Update("author", ewDatabaseLink.Variation.Author.Name)
// 			} else {
// 				tx.Model(&ewDatabaseLink).
// 					Update("author", ewSong.Author)
// 				tx.Model(&ewDatabaseLink.Variation.Author).
// 					Update("name", ewSong.Author)
// 			}
// 		}
// 	} else {
// 		newAuthor := db.CreateAuthorByName(
// 			tx,
// 			ewSong.Author,
// 		)
// 		tx.Model(&ewDatabaseLink.Variation).
// 			Update("author_id", newAuthor.ID)

// 		if ewDatabaseLink.Author != ewSong.Author {
// 			tx.Model(&ewDatabaseLink).
// 				Update("author", ewSong.Author)
// 		}
// 	}
// }

// func ewSyncProcessUpdateCopyrightState(
// 	tx *gorm.DB,
// 	ewSong *MatiasService.EwSong,
// ) {
// 	if ewDatabaseLink.Variation.Copyright.ID > 0 {
// 		if ewDatabaseLink.Copyright == ewSong.Copyright {
// 			if ewDatabaseLink.Variation.Copyright.Name != ewSong.Copyright {
// 				newEwSong := res.CreateOrGetEwSong(ewSong.Id)
// 				newEwSong.Copyright = ewDatabaseLink.Variation.Copyright.Name
// 				tx.Model(&ewDatabaseLink).
// 					Update("copyright", ewDatabaseLink.Variation.Copyright.Name)
// 			}
// 		} else {
// 			if ewDatabaseLink.Variation.Copyright.Name != ewDatabaseLink.Copyright {
// 				newEwSong := res.CreateOrGetEwSong(ewSong.Id)
// 				newEwSong.Copyright = ewDatabaseLink.Variation.Copyright.Name
// 				tx.Model(&ewDatabaseLink).
// 					Update("copyright", ewDatabaseLink.Variation.Copyright.Name)
// 			} else {
// 				tx.Model(&ewDatabaseLink).
// 					Update("copyright", ewSong.Copyright)
// 				tx.Model(&ewDatabaseLink.Variation.Copyright).
// 					Update("name", ewSong.Copyright)
// 			}
// 		}
// 	} else {
// 		newCopyright := db.CreateCopyrightByName(
// 			tx,
// 			ewSong.Copyright,
// 		)

// 		tx.Model(&ewDatabaseLink.Variation).
// 			Update("copyright_id", newCopyright.ID)

// 		if ewDatabaseLink.Copyright != ewSong.Copyright {
// 			tx.Model(&ewDatabaseLink).
// 				Update("copyright", ewSong.Copyright)
// 		}
// 	}
// }

// func ewSyncProcessHasDatabaseLink(
// 	tx *gorm.DB,
// 	ewSong *MatiasService.EwSong,
// ) {
// 	if ewDatabaseLink.Variation.ID == 0 {
// 		return
// 	}
// 	if ewSong.Author != "" {
// 		ewSyncProcessUpdateAuthorState(
// 			tx,
// 			ewSong,
// 		)
// 	} else {
// 		if ewDatabaseLink.Variation.Author.ID > 0 {
// 			newEwSong := res.CreateOrGetEwSong(ewSong.Id)
// 			newEwSong.Author = ewDatabaseLink.Variation.Author.Name
// 		}
// 	}

// 	if ewSong.Copyright != "" {
// 		ewSyncProcessUpdateCopyrightState(
// 			tx,
// 			ewSong,
// 		)
// 	} else {
// 		if ewDatabaseLink.Variation.Copyright.ID > 0 {
// 			newEwSong := res.CreateOrGetEwSong(ewSong.Id)
// 			newEwSong.Copyright = ewDatabaseLink.Variation.Copyright.Name
// 		}
// 	}

// 	newestVariationVersion := ewDatabaseLink.Variation.FindNewestVersion()
// 	if newestVariationVersion.ID == 0 {
// 		return
// 	}

// 	if newestVariationVersion.Version == ewDatabaseLink.Version {
// 		if newestVariationVersion.DisabledAt == nil &&
// 			ewDatabase.SongDatabase.
// 				HasVariation(newestVariationVersion.VariationID) == true {
// 			if ewSong.Title != newestVariationVersion.Name ||
// 				ewSong.Text != newestVariationVersion.Text {

// 				newVariationVersion := help.NewVariationVersionFromEwSong(
// 					ewSong,
// 					newestVariationVersion.VariationID,
// 					newestVariationVersion.Version+1,
// 				)
// 				tx.Model(&newestVariationVersion).
// 					Update("disabled_at", time.Now())

// 				srNewPassivatedVariationVersions = append(
// 					srNewPassivatedVariationVersions,
// 					db.SrPassivatedVariationVersion{
// 						SrID:               synchronizationRaport.ID,
// 						VariationVersionID: newestVariationVersion.ID,
// 					},
// 				)

// 				createNewVariationVersions = append(
// 					createNewVariationVersions,
// 					*newVariationVersion,
// 				)
// 			}
// 		} else {
// 			if ewDatabase.RemoveSongsFromEwDatabase == true {
// 				res.RemoveEwSongIds = append(
// 					res.RemoveEwSongIds,
// 					ewDatabaseLink.EwDatabaseSongID,
// 				)
// 				removeEwDatabaseLinks = append(
// 					removeEwDatabaseLinks,
// 					ewDatabaseLink.ID,
// 				)

// 				srEwSongs = append(
// 					srEwSongs,
// 					db.SrEwSong{
// 						SrID:               synchronizationRaport.ID,
// 						VariationVersionID: newestVariationVersion.ID,
// 						Operation:          false,
// 					},
// 				)
// 				srEwDatabaseLinks = append(
// 					srEwDatabaseLinks,
// 					db.SrEwDatabaseLink{
// 						SrID:             synchronizationRaport.ID,
// 						EwDatabaseID:     ewDatabaseLink.EwDatabaseID,
// 						EwDatabaseSongID: ewDatabaseLink.EwDatabaseSongID,
// 						VariationID:      ewDatabaseLink.VariationID,
// 						Version:          ewDatabaseLink.Version,
// 						Author:           ewDatabaseLink.Author,
// 						Copyright:        ewDatabaseLink.Copyright,
// 						Operation:        false,
// 					},
// 				)
// 			}
// 		}
// 	} else {
// 		if newestVariationVersion.DisabledAt == nil {
// 			ewLinkVariationVersion := ewDatabaseLink.Variation.
// 				FindVersionWithVersionNumber(ewDatabaseLink.Version)

// 			if ewSong.Title != ewLinkVariationVersion.Name ||
// 				ewSong.Text != ewLinkVariationVersion.Text {

// 				switch ewDatabase.VariationVersionConflictAction {
// 				case 1: // Use ew version
// 					newVariationVersion := help.NewVariationVersionFromEwSong(
// 						ewSong,
// 						newestVariationVersion.VariationID,
// 						newestVariationVersion.Version,
// 					)
// 					createNewVariationVersions = append(
// 						createNewVariationVersions,
// 						*newVariationVersion,
// 					)

// 				case 2: // Use database
// 					tx.Model(&ewDatabaseLink).
// 						Update(
// 							"version",
// 							newestVariationVersion.Version,
// 						)

// 					ewSong := res.CreateOrGetEwSong(ewSong.Id)

// 					ewSong.Title = newestVariationVersion.Name
// 					ewSong.Text = newestVariationVersion.Text
// 				case 3: // Report conflict
// 					srEwConflicts = append(
// 						srEwConflicts,
// 						db.SrEwConflict{
// 							VariationVersionID: newestVariationVersion.ID,
// 							EwDatabaseID:       ewDatabase.ID,
// 							EwSongID:           ewSong.Id,
// 							Name:               ewSong.Title,
// 							Text:               ewSong.Text,
// 							Resolved:           0,
// 						},
// 					)
// 				case 4: // Create branch
// 					newVariation, newVariationVersion := help.NewVariationVersionFromVariationVersion(
// 						tx,
// 						&newestVariationVersion,
// 					)

// 					if ewSong.Author != "" {
// 						newAuthor := db.CreateAuthorByName(
// 							tx,
// 							ewSong.Author,
// 						)
// 						tx.Model(&newVariation).
// 							Update("author_id", newAuthor.ID)
// 					}

// 					if ewSong.Copyright != "" {
// 						newCopyright := db.CreateCopyrightByName(
// 							tx,
// 							ewSong.Copyright,
// 						)
// 						tx.Model(&newCopyright).
// 							Update("copyright_id", newCopyright.ID)
// 					}

// 					createBranches = append(
// 						createBranches,
// 						db.Branch{
// 							SourceVariationVersionID:      newestVariationVersion.ID,
// 							DestinationVariationVersionID: newVariationVersion.ID,
// 						},
// 					)

// 					addVariationsToSongDatabase = append(
// 						addVariationsToSongDatabase,
// 						db.SongDatabaseVariation{
// 							SongDatabaseID: ewDatabase.SongDatabaseID,
// 							VariationID:    newVariation.ID,
// 						},
// 					)
// 					addVariationsToEwDatabase = append(
// 						addVariationsToEwDatabase,
// 						db.EwDatabaseLink{
// 							EwDatabaseID:     ewDatabase.ID,
// 							EwDatabaseSongID: ewSong.Id,
// 							VariationID:      newVariation.ID,
// 							Version:          newVariationVersion.Version,
// 						},
// 					)
// 				default:
// 					tx.Model(&ewDatabaseLink).
// 						Update(
// 							"version",
// 							newestVariationVersion.Version,
// 						)

// 					ewSong := res.CreateOrGetEwSong(ewSong.Id)

// 					ewSong.Title = newestVariationVersion.Name
// 					ewSong.Text = newestVariationVersion.Text
// 				}
// 			} else {
// 				tx.Model(&ewDatabaseLink).
// 					Update(
// 						"version",
// 						newestVariationVersion.Version,
// 					)

// 				ewSong := res.CreateOrGetEwSong(ewSong.Id)

// 				ewSong.Title = newestVariationVersion.Name
// 				ewSong.Text = newestVariationVersion.Text
// 			}
// 		} else {
// 			if ewDatabase.RemoveSongsFromEwDatabase == true {
// 				res.RemoveEwSongIds = append(
// 					res.RemoveEwSongIds,
// 					ewDatabaseLink.EwDatabaseSongID,
// 				)
// 				removeEwDatabaseLinks = append(
// 					removeEwDatabaseLinks,
// 					ewDatabaseLink.ID,
// 				)
// 			}
// 		}
// 	}
// }

// func ewSyncProcessEwSong(
// 	tx *gorm.DB,
// 	ewSong *MatiasService.EwSong,
// ) {
// 	ewDatabaseLink := ewDatabase.FindEwDatabaseLinkByEwSongID(ewSong.Id)
// 	if ewDatabaseLink == nil {
// 		prosessEwDatabaseNoDatabaseLink(
// 			tx,
// 		)
// 		return
// 	}
// 	ewSyncProcessHasDatabaseLink(
// 		tx,
// 	)
// }

// func ewSyncProcessEwSongs(
// 	tx *gorm.DB,
// 	ewDatabase *db.EwDatabase,
// 	ewSongs map[uint32]*MatiasService.EwSong,
// ) {
// 	for _, ewSong := range ewSongs {
// 		ewSyncProcessEwSongs(
// 			tx,
// 			ewSong,
// 		)
// 	}
// }

// func ewSyncRaportGeneration(
// 	tx *gorm.DB,
// 	srAddSongDatabaseVariations []db.SrAddSongDatabaseVariation,
// 	srEwConflicts []db.SrEwConflict,
// 	srEwDatabaseLinks []db.SrEwDatabaseLink,
// 	srEwSongs []db.SrEwSong,
// 	srNewAuthors []db.SrNewAuthors,
// 	srNewBranches []db.SrNewBranch,
// 	srNewCopyrights []db.SrcNewCopyright,
// 	srNewVariations []db.SrNewVariation,
// 	srNewVariationVersions []db.SrNewVariationVersion,
// 	srNewPassivatedVariationVersions []db.SrPassivatedVariationVersion,
// 	srRemoveSongDatabaseVariations []db.SrRemoveSongDatabaseVariation,
// ) {
// 	managers.BatchCreateSrAddSongDatabaseVariations(
// 		tx,
// 		srAddSongDatabaseVariations,
// 	)
// 	managers.BatchCreateSrEwConflicts(
// 		tx,
// 		srEwConflicts,
// 	)
// 	managers.BatchCreateSrEwDatabaseLinks(
// 		tx,
// 		srEwDatabaseLinks,
// 	)
// 	managers.BatchCreateSrEwSongs(
// 		tx,
// 		srEwSongs,
// 	)
// 	managers.BatchCreateSrNewAuthors(
// 		tx,
// 		srNewAuthors,
// 	)
// 	managers.BatchCreateSrNewBranches(
// 		tx,
// 		srNewBranches,
// 	)
// 	managers.BatchCreateSrNewCopyrights(
// 		tx,
// 		srNewCopyrights,
// 	)
// 	managers.BatchCreateSrNewVariations(
// 		tx,
// 		srNewVariations,
// 	)
// 	managers.BatchCreateSrNewVariationVersions(
// 		tx,
// 		srNewVariationVersions,
// 	)
// 	managers.BatchCreateSrRemoveSongDatabaseVariations(
// 		tx,
// 		srRemoveSongDatabaseVariations,
// 	)
// }

// timeNow := time.Now()
// logFile, err := os.OpenFile(
// 	"logs/" + timeNow.Format("2006-01-02T15-04-05Z07-00")+".txt",
// 	os.O_RDWR|os.O_CREATE|os.O_APPEND,
// 	0666,
// )
// if err != nil {
// 	fmt.Println("error opening file: %v", err)
// }

// w := bufio.NewWriter(logFile)
// tx.SetLogger(log.New(w, "\r\n", 0))

// var synchronizationRaport db.SynchronizationRaport
// tx.Create(&synchronizationRaport)
// startDate := time.Now()
// startTime := startDate.UnixNano() / int64(time.Millisecond)

// synchronizationRaport.StartedAt = &startDate
// synchronizationRaport.DatabaseKey = in.EwDatabaseKey

// db, err := sql.Open(
// 	"mysql",
// 	"root:asdf321@tcp(localhost:3306)/seppo2?parseTime=True&loc=Local",
// )

// models.EwDatabases(db)

// ewDatabases, _ := models.EwDatabases(db).All()

// for _, ewDatabase := range ewDatabases {
// 	ewDatabase.DEl
// }

// models.(ew)

// var ewDatabase db.EwDatabase
// tx.Where("ew_databases.ew_database_key = ?", in.EwDatabaseKey).
// 	Preload("SongDatabase").
// 	Preload("SongDatabase.Variations").
// 	Preload("SongDatabase.Variations.Author").
// 	Preload("SongDatabase.Variations.Copyright").
// 	Preload("SongDatabase.Variations.VariationVersions").
// 	Preload("SongDatabase.SongDatabaseTags").
// 	Preload("EwDatabaseLinks").
// 	Preload("EwDatabaseLinks.Variation").
// 	Preload("EwDatabaseLinks.Variation.Author").
// 	Preload("EwDatabaseLinks.Variation.Copyright").
// 	Preload("EwDatabaseLinks.Variation.VariationVersions").
// 	Find(&ewDatabase)

// if ewDatabase.ID == 0 {
// 	res.EwDatabaseFound = false
// 	synchronizationRaport.DatabaseFound = false
// 	endDate := time.Now()
// 	endTime := endDate.UnixNano() / int64(time.Millisecond)
// 	synchronizationRaport.DurationMS = endTime - startTime
// 	synchronizationRaport.FinishedAt = &endDate
// 	tx.Save(&synchronizationRaport)
// 	tx.Commit()
// 	return res, nil
// }

// synchronizationRaport.DatabaseFound = true
// synchronizationRaport.DatabaseID = &ewDatabase.ID

// // These slices are used for batch insert
// var addVariationsToSongDatabase []db.SongDatabaseVariation
// var addVariationsToEwDatabase []db.EwDatabaseLink
// var createBranches []db.Branch
// //var srEwConflicts []db.SrEwConflict
// var createNewVariationVersions []db.VariationVersion

// // These slices are used for batch remove
// var removeEwDatabaseLinks []uint32
// var removeSongDatabaseVariationLinks []uint32

// // These are used for synchronization raport generation
// var srAddSongDatabaseVariations []db.SrAddSongDatabaseVariation
// var srEwConflicts []db.SrEwConflict
// var srEwDatabaseLinks []db.SrEwDatabaseLink
// var srEwSongs []db.SrEwSong
// var srNewAuthors []db.SrNewAuthors
// var srNewBranches []db.SrNewBranch
// var srNewCopyrights []db.SrcNewCopyright
// var srNewVariations []db.SrNewVariation
// var srNewVariationVersions []db.SrNewVariationVersion
// var srNewPassivatedVariationVersions []db.SrPassivatedVariationVersion
// var srRemoveSongDatabaseVariations []db.SrRemoveSongDatabaseVariation

// ewSongs, removeEwSongIds := managers.RemoveDuplicatesFromEwSong(in.EwSongs)
// res.RemoveEwSongIds = append(res.RemoveEwSongIds, removeEwSongIds...)

// ewSyncProcessEwSongs(
// 	tx,
// 	ewSongs,
// )

// var excludeVariationIds []uint32

// for _, variation := range ewDatabase.SongDatabase.Variations {
// 	excludeVariationIds = append(
// 		excludeVariationIds,
// 		variation.ID,
// 	)

// 	if is, link := ewDatabase.HasVariation(variation.ID); is == true {
// 		if ewSongs[link.EwDatabaseSongID] == nil {
// 			if ewDatabase.RemoveSongsFromSongDatabase == true {
// 				removeSongDatabaseVariationLinks = append(
// 					removeSongDatabaseVariationLinks,
// 					variation.ID,
// 				)
// 				removeEwDatabaseLinks = append(
// 					removeEwDatabaseLinks,
// 					link.ID,
// 				)
// 				srRemoveSongDatabaseVariations = append(
// 					srRemoveSongDatabaseVariations,
// 					db.SrRemoveSongDatabaseVariation{
// 						SrID:           synchronizationRaport.ID,
// 						VariationID:    variation.ID,
// 						SongDatabaseID: ewDatabase.SongDatabaseID,
// 					},
// 				)
// 				srEwDatabaseLinks = append(
// 					srEwDatabaseLinks,
// 					db.SrEwDatabaseLink{
// 						SrID:             synchronizationRaport.ID,
// 						EwDatabaseID:     ewDatabase.ID,
// 						EwDatabaseSongID: link.EwDatabaseSongID,
// 						VariationID:      variation.ID,
// 						Version:          link.Version,
// 						Author:           link.Author,
// 						Copyright:        link.Copyright,
// 						Operation:        false,
// 					},
// 				)
// 			} else {
// 				newestVersion := variation.FindNewestVersion()
// 				if newestVersion.ID > 0 {
// 					removeEwDatabaseLinks = append(
// 						removeEwDatabaseLinks,
// 						link.ID,
// 					)

// 					newEwSong := &MatiasService.EwSong{
// 						VariationId: variation.ID,
// 						Title:       newestVersion.Name,
// 						Text:        newestVersion.Text,
// 					}

// 					if variation.Author.ID > 0 {
// 						newEwSong.Author = variation.Author.Name
// 					}

// 					if variation.Copyright.ID > 0 {
// 						newEwSong.Copyright = variation.Copyright.Name
// 					}

// 					res.EwSongs = append(
// 						res.EwSongs,
// 						newEwSong,
// 					)
// 				}
// 			}
// 		}
// 	} else {
// 		newestVersion := variation.FindNewestVersion()
// 		if managers.HasEwSongByNameAndText(
// 			ewSongs,
// 			newestVersion.Name,
// 			newestVersion.Text,
// 		) == false {
// 			if newestVersion.ID > 0 {
// 				newEwSong := &MatiasService.EwSong{
// 					VariationId: variation.ID,
// 					Title:       newestVersion.Name,
// 					Text:        newestVersion.Text,
// 				}

// 				if variation.Author.ID > 0 {
// 					newEwSong.Author = variation.Author.Name
// 				}

// 				if variation.Copyright.ID > 0 {
// 					newEwSong.Copyright = variation.Copyright.Name
// 				}

// 				res.EwSongs = append(
// 					res.EwSongs,
// 					newEwSong,
// 				)
// 			}
// 		}
// 	}
// }

// var songDatabaseTagVariations []db.Variation
// q := tx.Table("variations").
// 	Preload("Author").
// 	Preload("Copyright").
// 	Preload("VariationVersions").
// 	Joins("left join tag_variations on tag_variations.variation_id = variations.id").
// 	Joins("left join song_database_tags on song_database_tags.tag_id = tag_variations.tag_id").
// 	Where("song_database_tags.song_database_id = ?", ewDatabase.SongDatabaseID).
// 	Select("distinct variations.*")
// if len(excludeVariationIds) > 0 {
// 	q = q.Where("variations.id not in(?)", excludeVariationIds)
// }
// q.Find(&songDatabaseTagVariations)

// for _, variation := range songDatabaseTagVariations {
// 	if is, link := ewDatabase.HasVariation(variation.ID); is == true {
// 		if ewSongs[link.EwDatabaseSongID] == nil {
// 			newestVersion := variation.FindNewestVersion()
// 			if newestVersion.ID > 0 {
// 				removeEwDatabaseLinks = append(
// 					removeEwDatabaseLinks,
// 					link.ID,
// 				)
// 				newEwSong := &MatiasService.EwSong{
// 					VariationId: variation.ID,
// 					Title:       newestVersion.Name,
// 					Text:        newestVersion.Text,
// 				}

// 				if variation.Author.ID > 0 {
// 					newEwSong.Author = variation.Author.Name
// 				}

// 				if variation.Copyright.ID > 0 {
// 					newEwSong.Copyright = variation.Copyright.Name
// 				}

// 				res.EwSongs = append(
// 					res.EwSongs,
// 					newEwSong,
// 				)
// 			}
// 		}
// 	} else {
// 		newestVersion := variation.FindNewestVersion()
// 		if managers.HasEwSongByNameAndText(
// 			ewSongs,
// 			newestVersion.Name,
// 			newestVersion.Text,
// 		) == false {
// 			if newestVersion.ID > 0 {
// 				newEwSong := &MatiasService.EwSong{
// 					VariationId: variation.ID,
// 					Title:       newestVersion.Name,
// 					Text:        newestVersion.Text,
// 				}

// 				if variation.Author.ID > 0 {
// 					newEwSong.Author = variation.Author.Name
// 				}

// 				if variation.Copyright.ID > 0 {
// 					newEwSong.Copyright = variation.Copyright.Name
// 				}

// 				res.EwSongs = append(
// 					res.EwSongs,
// 					newEwSong,
// 				)
// 			}
// 		}
// 	}
// }

// var affectedRows int64

// affectedRows = managers.BatchCreateVariationVersions(
// 	tx,
// 	createNewVariationVersions,
// )
// var newVariationVersions []db.VariationVersion
// tx.Limit(affectedRows).Find(&newVariationVersions)
// for _, newVariationVersion := range newVariationVersions {
// 	srNewVariationVersions = append(
// 		srNewVariationVersions,
// 		db.SrNewVariationVersion{
// 			SrID:               synchronizationRaport.ID,
// 			VariationVersionID: newVariationVersion.ID,
// 		},
// 	)
// }

// affectedRows = managers.BatchAddVariationsToSongDatabase(
// 	tx,
// 	addVariationsToSongDatabase,
// )
// var newSongDatabaseVariations []db.SongDatabaseVariation
// tx.Limit(affectedRows).Find(&newSongDatabaseVariations)
// for _, songDatabaseVariation := range newSongDatabaseVariations {
// 	srAddSongDatabaseVariations = append(
// 		srAddSongDatabaseVariations,
// 		db.SrAddSongDatabaseVariation{
// 			SrID:           synchronizationRaport.ID,
// 			VariationID:    songDatabaseVariation.VariationID,
// 			SongDatabaseID: songDatabaseVariation.SongDatabaseID,
// 		},
// 	)
// }

// affectedRows = managers.BatchAddVariationsToEwDatabase(
// 	tx,
// 	addVariationsToEwDatabase,
// )

// // var addedVariationsToEwDatabase []db.EwDatabaseLink
// // tx.Limit(affectedRows).First(&addedVariationsToEwDatabase)
// // for _, ewDatabaseLink := range addedVariationsToEwDatabase {
// // 	synchronizationRaport.NewEwDatabaseLinks = append(
// // 		synchronizationRaport.NewEwDatabaseLinks,
// // 		db.SrEwDatabaseLink{
// // 			EwDatabaseID:     ewDatabaseLink.EwDatabaseID,
// // 			EwDatabaseSongID: ewDatabaseLink.EwDatabaseSongID,
// // 			VariationID:      ewDatabaseLink.VariationID,
// // 			Version:          ewDatabaseLink.Version,
// // 			Author:           ewDatabaseLink.Author,
// // 			Copyright:        ewDatabaseLink.Copyright,
// // 			Operation:        true,
// // 		},
// // 	)
// // }

// affectedRows = managers.BatchCreateBranches(
// 	tx,
// 	createBranches,
// )
// var newBranches []db.SrNewBranch
// tx.Limit(affectedRows).First(&newBranches)
// for _, newBranch := range newBranches {
// 	srNewBranches = append(
// 		srNewBranches,
// 		db.SrNewBranch{
// 			SrID:     synchronizationRaport.ID,
// 			BranchID: newBranch.ID,
// 		},
// 	)
// }

// if len(removeEwDatabaseLinks) > 0 {
// 	tx.Where("id in (?)", removeEwDatabaseLinks).Delete(&db.EwDatabaseLink{})
// }

// if len(removeSongDatabaseVariationLinks) > 0 {
// 	tx.Where("variation_id in (?)", removeSongDatabaseVariationLinks).
// 		Where("song_database_id = ?", ewDatabase.SongDatabaseID).
// 		Delete(&db.SongDatabaseVariation{})
// }

// endDate := time.Now()
// endTime := endDate.UnixNano() / int64(time.Millisecond)
// synchronizationRaport.DurationMS = endTime - startTime
// synchronizationRaport.FinishedAt = &endDate

// // Raport generation
// ewSyncRaportGeneration(
// 	tx,
// 	srAddSongDatabaseVariations,
// 	srEwConflicts,
// 	srEwDatabaseLinks,
// 	srEwSongs,
// 	srNewAuthors,
// 	srNewBranches,
// 	srNewCopyrights,
// 	srNewVariations,
// 	srNewVariationVersions,
// 	srNewPassivatedVariationVersions,
// 	srRemoveSongDatabaseVariations,
// )

// tx.Save(&synchronizationRaport)

// tx.Commit()
