package managers

import "github.com/koodinikkarit/seppo/models"

func FindVariationVersionByVersion(
	variationVersions []*models.VariationVersion,
	versionNumber uint32,
) *models.VariationVersion {
	for _, variationVersion := range variationVersions {
		if variationVersion.Version == versionNumber {
			return variationVersion
		}
	}
	return nil
}
