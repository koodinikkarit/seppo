package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

func BatchCreateVariationVersions(
	tx *gorm.DB,
	variationVersions []*VariationVersion,
) {
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
	}
}
