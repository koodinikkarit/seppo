package help

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
)

func NewVariationVersionFromVariationVersion(
	tx *gorm.DB,
	variationVersion *db.VariationVersion,
) (
	*db.Variation,
	*db.VariationVersion,
) {
	newVariation := &db.Variation{}
	tx.Create(&newVariation)
	newVariationVersion := &db.VariationVersion{
		VariationID: newVariation.ID,
		Name:        variationVersion.Name,
		Text:        variationVersion.Text,
		Version:     1,
		Newest:      true,
	}
	tx.Create(&newVariationVersion)

	return newVariation, newVariationVersion
}
