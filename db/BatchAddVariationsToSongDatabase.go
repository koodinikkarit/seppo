package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

func BatchAddVariationsToSongDatabase(
	tx *gorm.DB,
	songDatabaseVariations []*SongDatabaseVariation,
) {
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
	}
}
