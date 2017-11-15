package managers

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
)

func BatchCreateSrAddSongDatabaseVariations(
	tx *gorm.DB,
	srAddSongDatabaseVariations []db.SrAddSongDatabaseVariation,
) {
	if len(srAddSongDatabaseVariations) > 0 {
		sqlStr := "INSERT INTO `sr_add_song_database_variations` (`sr_id`, `variation_id`, `song_database_id`) VALUES "
		vals := []interface{}{}

		for _, srAddSongDatabaseVariation := range srAddSongDatabaseVariations {
			sqlStr += "(?, ?, ?), "
			vals = append(
				vals,
				srAddSongDatabaseVariation.SrID,
				srAddSongDatabaseVariation.VariationID,
				srAddSongDatabaseVariation.SongDatabaseID,
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
	}
}

func BatchCreateSrEwConflicts(
	tx *gorm.DB,
	srEwConflicts []db.SrEwConflict,
) {
	if len(srEwConflicts) > 0 {
		sqlStr := "INSERT INTO `tag_variations` (`sr_id`, `variation_version_id`, `ew_database_id`, `ew_song_id`, `name`, `text`, `resolved`) VALUES "
		vals := []interface{}{}

		for _, srEwConclict := range srEwConflicts {
			sqlStr += "(?, ?, ?, ?, ?, ?, ?), "
			vals = append(
				vals,
				srEwConclict.SrID,
				srEwConclict.VariationVersionID,
				srEwConclict.EwDatabaseID,
				srEwConclict.EwSongID,
				srEwConclict.Name,
				srEwConclict.Text,
				srEwConclict.Resolved,
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
	}
}

func BatchCreateSrEwDatabaseLinks(
	tx *gorm.DB,
	srEwDatabaseLinks []db.SrEwDatabaseLink,
) {
	if len(srEwDatabaseLinks) > 0 {
		sqlStr := "INSERT INTO `sr_ew_database_links` (`sr_id`, `ew_database_id`, `ew_database_song_id`, `variation_id`, `version`, `author`, `copyright`, `operation`) VALUES "
		vals := []interface{}{}

		for _, srEwDatabaseLink := range srEwDatabaseLinks {
			sqlStr += "(?, ?, ?, ?, ?, ?, ?, ?), "
			vals = append(
				vals,
				srEwDatabaseLink.SrID,
				srEwDatabaseLink.EwDatabaseID,
				srEwDatabaseLink.EwDatabaseSongID,
				srEwDatabaseLink.VariationID,
				srEwDatabaseLink.Version,
				srEwDatabaseLink.Author,
				srEwDatabaseLink.Copyright,
				srEwDatabaseLink.Operation,
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
	}
}

func BatchCreateSrEwSongs(
	tx *gorm.DB,
	srEwSongs []db.SrEwSong,
) {
	if len(srEwSongs) > 0 {
		sqlStr := "INSERT INTO `sr_ew_songs` (`sr_id`, `variation_version_id`, `operation`) VALUES "
		vals := []interface{}{}

		for _, srEwSong := range srEwSongs {
			sqlStr += "(?, ?, ?), "
			vals = append(
				vals,
				srEwSong.SrID,
				srEwSong.VariationVersionID,
				srEwSong.Operation,
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
	}
}

func BatchCreateSrNewAuthors(
	tx *gorm.DB,
	srNewAuthors []db.SrNewAuthors,
) {
	if len(srNewAuthors) > 0 {
		sqlStr := "INSERT INTO `sr_new_authors` (`sr_id`, `author_id`) VALUES "
		vals := []interface{}{}

		for _, srNewAuthor := range srNewAuthors {
			sqlStr += "(?, ?), "
			vals = append(
				vals,
				srNewAuthor.SrID,
				srNewAuthor.AuthorID,
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
	}
}

func BatchCreateSrNewBranches(
	tx *gorm.DB,
	srNewBranches []db.SrNewBranch,
) {
	if len(srNewBranches) > 0 {
		sqlStr := "INSERT INTO `sr_new_branches` (`sr_id`, `branch_id`) VALUES "
		vals := []interface{}{}

		for _, srNewBranch := range srNewBranches {
			sqlStr += "(?, ?, ?), "
			vals = append(
				vals,
				srNewBranch.SrID,
				srNewBranch.BranchID,
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
	}
}

func BatchCreateSrNewCopyrights(
	tx *gorm.DB,
	srNewCopyrights []db.SrcNewCopyright,
) {
	if len(srNewCopyrights) > 0 {
		sqlStr := "INSERT INTO `sr_new_copyrights` (`sr_id`, `author_id`) VALUES "
		vals := []interface{}{}

		for _, srNewCopyright := range srNewCopyrights {
			sqlStr += "(?, ?), "
			vals = append(
				vals,
				srNewCopyright.SrID,
				srNewCopyright.CopyrightID,
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
	}
}

func BatchCreateSrNewVariations(
	tx *gorm.DB,
	srNewVariations []db.SrNewVariation,
) {
	if len(srNewVariations) > 0 {
		sqlStr := "INSERT INTO `sr_new_variations` (`sr_id`, `variation_id`) VALUES "
		vals := []interface{}{}

		for _, srNewVariation := range srNewVariations {
			sqlStr += "(?, ?), "
			vals = append(
				vals,
				srNewVariation.SrID,
				srNewVariation.VariationID,
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
	}
}

func BatchCreateSrNewVariationVersions(
	tx *gorm.DB,
	srNewVariationVersions []db.SrNewVariationVersion,
) {
	if len(srNewVariationVersions) > 0 {
		sqlStr := "INSERT INTO `sr_new_variation_versions` (`sr_id`, `variation_version_id`) VALUES "
		vals := []interface{}{}

		for _, srNewVariationVersion := range srNewVariationVersions {
			sqlStr += "(?, ?), "
			vals = append(
				vals,
				srNewVariationVersion.SrID,
				srNewVariationVersion.VariationVersionID,
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
	}
}

func BatchCreateSrPassivatedVariationVersions(
	tx *gorm.DB,
	srPassivatedVariationVersions []db.SrPassivatedVariationVersion,
) {
	if len(srPassivatedVariationVersions) > 0 {
		sqlStr := "INSERT INTO `sr_passivated_variation_versions` (`sr_id`, `variation_version_id`) VALUES "
		vals := []interface{}{}

		for _, srPassivatedVariationVersion := range srPassivatedVariationVersions {
			sqlStr += "(?, ?), "
			vals = append(
				vals,
				srPassivatedVariationVersion.SrID,
				srPassivatedVariationVersion.VariationVersionID,
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
	}
}

func BatchCreateSrRemoveSongDatabaseVariations(
	tx *gorm.DB,
	srRemoveSongDatabaseVariations []db.SrRemoveSongDatabaseVariation,
) {
	if len(srRemoveSongDatabaseVariations) > 0 {
		sqlStr := "INSERT INTO `sr_remove_song_database_variations` (`sr_id`, `variation_id`, `song_database_id`) VALUES "
		vals := []interface{}{}

		for _, srRemoveSongDatabaseVariation := range srRemoveSongDatabaseVariations {
			sqlStr += "(?, ?, ?), "
			vals = append(
				vals,
				srRemoveSongDatabaseVariation.SrID,
				srRemoveSongDatabaseVariation.VariationID,
				srRemoveSongDatabaseVariation.SongDatabaseID,
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
	}
}
