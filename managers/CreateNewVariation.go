package managers

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
)

func CreateNewVariation(
	tx *gorm.DB,
	name string,
	text string,
) (
	db.Variation,
	db.VariationVersion,
) {
	newVariation := db.Variation{}
	newVariationVersion := db.VariationVersion{
		Name:    name,
		Text:    text,
		Version: 1,
	}
	newVariation.VariationVersions = append(
		newVariation.VariationVersions,
		newVariationVersion,
	)
	tx.Create(&newVariation)
	return newVariation, newVariationVersion
}
