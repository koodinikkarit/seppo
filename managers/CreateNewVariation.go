package managers

import (
	"database/sql"

	"github.com/koodinikkarit/seppo/models"
)

func CreateNewVariation(
	tx *sql.Tx,
	name string,
	text string,
) (
	*models.Variation,
	*models.VariationVersion,
) {
	newVariation := models.Variation{}
	newVariation.Insert(tx)
	newVariationVersion := models.VariationVersion{
		Name:    name,
		Text:    text,
		Version: 1,
	}
	newVariation.AddVariationVersions(
		tx,
		true,
		&newVariationVersion,
	)
	return &newVariation, &newVariationVersion
}
