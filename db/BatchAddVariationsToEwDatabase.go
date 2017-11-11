package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

func BatchAddVariationsToEwDatabase(
	tx *gorm.DB,
	ewDatabaseLinks []*EwDatabaseLink,
) {
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
	}
}
