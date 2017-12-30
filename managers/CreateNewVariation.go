package managers

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/models"
)

func CreateNewVariation(
	tx *gorm.DB,
	name string,
	text string,
) (
	models.Variation,
	models.VariationVersion,
) {
	newVariation := models.Variation{}
	newVariationVersion := models.VariationVersion{
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
