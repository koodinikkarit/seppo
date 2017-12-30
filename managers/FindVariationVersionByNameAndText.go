package managers

import "github.com/koodinikkarit/seppo/models"

func FindVariationVersionByNameAndText(
	variationVersions []*models.VariationVersion,
	name string,
	text string,
) *models.VariationVersion {
	for _, variationVersion := range variationVersions {
		if name != variationVersion.Name ||
			text != variationVersion.Text {
			continue
		}
		return variationVersion
	}
	return nil
}
