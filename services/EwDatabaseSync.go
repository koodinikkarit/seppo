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
	if ewDatabase.RemoveSongsFromEwDatabase == 1 {
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
	ewSong := managers.NewEwSongFromVariation(
		tx,
		variation,
		variationVersion,
	)
	res.EwSongs = append(
		res.EwSongs,
		ewSong,
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

func ewSyncProcessSongDatabaseTagVariation(
	tx *sql.Tx,
	res *MatiasService.SyncEwDatabaseResponse,
	ewSongs map[uint32]*MatiasService.EwSong,
	ewDatabase *models.EwDatabase,
	variation *models.Variation,
	synchronizationRaport *models.SynchronizationRaport,
) {
	variationVersions, _ := variation.VariationVersions(tx).All()
	newestVariationVersion := managers.FindNewestVariationVersion(
		variationVersions,
	)

	ewDatabaseLink, _ := ewDatabase.EwDatabaseLinks(
		tx,
		Where("ew_Database_links.variation_id", ewDatabase.SongDatabaseID),
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

	res.EwSongs = append(
		res.EwSongs,
		managers.NewEwSongFromVariation(
			tx,
			variation,
			newestVariationVersion,
		),
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

	songDatabaseVariations, _ := songDatabase.SongDatabaseVariations(tx).All()
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

	songDatabaseTagVariations, _ := models.Variations(
		tx,
		Load("VariationVersions"),
		InnerJoin("tag_variations tg on tg.variation_id = variations.id"),
		InnerJoin("song_database_tags sdt on sdt.tag_id = tg.tag_id"),
		Where("sdt.song_database_id = ?", ewDatabase.SongDatabaseID),
	).All()

	for _, variation := range songDatabaseTagVariations {
		ewSyncProcessSongDatabaseTagVariation(
			tx,
			res,
			ewSongs,
			ewDatabase,
			variation,
			&synchronizationRaport,
		)
	}

	tx.Commit()

	return res, nil
}
