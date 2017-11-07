package services

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
)

// func UpdateVariationVersion(
// 	variation
// )

func createNewVariationVersions(
	tx *gorm.DB,
	oldVariationVersion *db.VariationVersion,
	name string,
	text string,
) *db.VariationVersion {
	var sameVariationVersion *db.VariationVersion
	tx.Where("variation_versions.name = ?", name).
		Where("variation_versions.text = ?", text).
		First(&sameVariationVersion)

	if sameVariationVersion != nil {
		if sameVariationVersion.Newest == true ||
			sameVariationVersion.DisabledAt == nil {

			// MoveVariationVersionReferences(
			// 	tx,
			// 	oldVariationVersion.ID,
			// 	sameVariationVersion.ID,
			// )

			newMerge := db.Merge{
				VariationVersion1ID:           oldVariationVersion.ID,
				VariationVersion2ID:           sameVariationVersion.ID,
				DestinationVariationVersionID: sameVariationVersion.ID,
			}
			tx.Create(&newMerge)

			return sameVariationVersion
		}
		newMerge := db.Merge{
			VariationVersion1ID:           oldVariationVersion.ID,
			VariationVersion2ID:           sameVariationVersion.ID,
			DestinationVariationVersionID: sameVariationVersion.ID,
		}
		tx.Create(&newMerge)
		newVariation := db.Variation{}
		tx.Create(&newVariation)
		newVariationVersion := db.VariationVersion{
			VariationID: newVariation.ID,
			Name:        name,
			Text:        text,
			Version:     1,
			Newest:      true,
		}
		tx.Create(&newVariationVersion)
		tx.Save(&newVariation)
		newBranch := db.Branch{
			SourceVariationVersionID:      sameVariationVersion.ID,
			DestinationVariationVersionID: newVariationVersion.ID,
		}
		tx.Create(&newBranch)
		// MoveVariationVersionReferences(
		// 	tx,
		// 	oldVariationVersion.ID,
		// 	newVariationVersion.ID,
		// )
		return &newVariationVersion
	}
	oldVariationVersion.Newest = false
	tx.Save(&oldVariationVersion)
	newVariationVersion := db.VariationVersion{
		VariationID: oldVariationVersion.VariationID,
		Name:        name,
		Text:        text,
		Version:     oldVariationVersion.Version + 1,
		Newest:      true,
	}
	tx.Create(&newVariationVersion)
	return &newVariationVersion
}
