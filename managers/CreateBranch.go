package managers

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
)

func CreateBranchAndVariation(
	tx *gorm.DB,
	sourceVariationVersionId uint32,
	name string,
	text string,
) (
	*db.Variation,
	*db.VariationVersion,
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
	newBranch := db.Branch{
		SourceVariationVersionID:      sourceVariationVersionId,
		DestinationVariationVersionID: newVariationVersion.ID,
	}
	tx.Create(&newBranch)
	return &newVariation, &newVariationVersion
}
