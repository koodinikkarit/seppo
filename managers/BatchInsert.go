package managers

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
)

func BatchAddTagsToVariation(
	tx *gorm.DB,
	tagVariations []db.TagVariation,
) int64 {
	if len(tagVariations) > 0 {
		sqlStr := "INSERT INTO `tag_variations` (`tag_id`, `variation_id`, `created_at`) VALUES "
		vals := []interface{}{}

		for _, tagVariation := range tagVariations {
			sqlStr += "(?, ?, ?), "
			vals = append(
				vals,
				tagVariation.TagID,
				tagVariation.VariationID,
				time.Now(),
			)
		}

		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
		return tx.RowsAffected
	}
	return 0
}

func BatchAddVariationsToEwDatabase(
	tx *gorm.DB,
	ewDatabaseLinks []db.EwDatabaseLink,
) int64 {
	if len(ewDatabaseLinks) > 0 {
		sqlStr := "INSERT INTO `ew_database_links` (`ew_database_id`, `ew_database_song_id`, `variation_id`, `version`, `created_at`) VALUES "
		vals := []interface{}{}

		for _, ewDatabaseLink := range ewDatabaseLinks {
			sqlStr += "(?, ?, ?, ?, ?), "
			vals = append(
				vals,
				ewDatabaseLink.EwDatabaseID,
				ewDatabaseLink.EwDatabaseSongID,
				ewDatabaseLink.VariationID,
				ewDatabaseLink.Version,
				time.Now(),
			)
		}

		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
		return tx.RowsAffected
	}
	return 0
}

func BatchAddVariationsToSongDatabase(
	tx *gorm.DB,
	songDatabaseVariations []db.SongDatabaseVariation,
) int64 {
	if len(songDatabaseVariations) > 0 {
		sqlStr := "INSERT INTO `song_database_variations` (`song_database_id`, `variation_id`, `created_at`) VALUES "
		vals := []interface{}{}

		for _, songDatabaseVariation := range songDatabaseVariations {
			sqlStr += "(?, ?, ?), "
			vals = append(
				vals,
				songDatabaseVariation.SongDatabaseID,
				songDatabaseVariation.VariationID,
				time.Now(),
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
		return tx.RowsAffected
	}
	return 0
}

func BatchCreateBranches(
	tx *gorm.DB,
	branches []db.Branch,
) int64 {
	if len(branches) > 0 {
		sqlStr := "INSERT INTO `branches` (`source_variation_version_id`, `destination_variation_version_id`, `created_at`) VALUES "
		vals := []interface{}{}

		for _, branch := range branches {
			sqlStr += "(?, ?, ?), "
			vals = append(
				vals,
				branch.SourceVariationVersionID,
				branch.DestinationVariationVersionID,
				time.Now(),
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
		return tx.RowsAffected
	}
	return 0
}

func BatchCreateVariationVersions(
	tx *gorm.DB,
	variationVersions []db.VariationVersion,
) int64 {
	if len(variationVersions) > 0 {
		sqlStr := "INSERT INTO `variation_versions` (`variation_id`,`name`,`text`,`version`,`newest`,`created_at`) VALUES "
		vals := []interface{}{}

		for _, variationVersion := range variationVersions {
			sqlStr += "(?, ?, ?, ?, ?, ?), "
			vals = append(
				vals,
				variationVersion.VariationID,
				variationVersion.Name,
				variationVersion.Text,
				variationVersion.Version,
				variationVersion.Newest,
				time.Now(),
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
		return tx.RowsAffected
	}
	return 0
}

func BatchCreateScheduleVariations(
	tx *gorm.DB,
	scheduleVariations []db.ScheduleVariation,
) int64 {
	if len(scheduleVariations) > 0 {
		sqlStr := "INSERT INTO `schedule_variations` (`schedule_id`,`variation_id`,`order_number`,`created_at`) VALUES "
		vals := []interface{}{}

		for _, scheduleVariation := range scheduleVariations {
			sqlStr += "(?, ?, ?, ?), "
			vals = append(
				vals,
				scheduleVariation.ScheduleID,
				scheduleVariation.VariationID,
				scheduleVariation.OrderNumber,
				scheduleVariation.CreatedAt,
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
		return tx.RowsAffected
	}
	return 0
}

func BatchCreateSongDatabaseTags(
	tx *gorm.DB,
	songDatabaseTags []db.SongDatabaseTag,
) int64 {
	if len(songDatabaseTags) > 0 {
		sqlStr := "INSERT INTO song_database_tags (tag_id, song_database_id) VALUES "
		vals := []interface{}{}
		for _, songDatabaseTag := range songDatabaseTags {
			sqlStr += "(?, ?), "
			vals = append(
				vals,
				songDatabaseTag.TagID,
				songDatabaseTag.SongDatabaseID,
			)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-2]
		tx.Exec(sqlStr, vals...)
		return tx.RowsAffected
	}
	return 0
}
