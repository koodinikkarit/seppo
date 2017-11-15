package managers

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
)

func NewVariationFromVariationVersion(
	tx *gorm.DB,
	variationVersion db.VariationVersion,
) db.Variation {
	newVariation := db.Variation{}
	// newVariation.VariationVersions = append(
	// 	newVariation.VariationVersions,
	// 	db.VariationVersion{

	// 	}
	// )

	return newVariation
}
