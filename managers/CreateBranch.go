package managers

import (
	"database/sql"

	"github.com/koodinikkarit/seppo/models"
)

func CreateBranchAndVariation(
	tx *sql.Tx,
	sourceVariationVersionId uint64,
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
	newBranch := models.Branch{
		SourceVariationVersionID:      sourceVariationVersionId,
		DestinationVariationVersionID: newVariationVersion.ID,
	}
	newBranch.Insert(tx)
	return &newVariation, &newVariationVersion
}
