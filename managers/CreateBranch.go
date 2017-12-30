package managers

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/models"
)

func CreateBranchAndVariation(
	tx *gorm.DB,
	sourceVariationVersionId uint32,
	name string,
	text string,
) (
	*models.Variation,
	*models.VariationVersion,
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
	newBranch := models.Branch{
		SourceVariationVersionID:      sourceVariationVersionId,
		DestinationVariationVersionID: newVariationVersion.ID,
	}
	tx.Create(&newBranch)
	return &newVariation, &newVariationVersion
}
