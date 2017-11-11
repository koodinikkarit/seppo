package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

func BatchAddTagsToVariation(
	tx *gorm.DB,
	tagVariations []TagVariation,
) {
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
	}
}
