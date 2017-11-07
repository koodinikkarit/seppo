package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

func BatchCreateBranches(
	tx *gorm.DB,
	branches []*Branch,
) {
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
	}
}
