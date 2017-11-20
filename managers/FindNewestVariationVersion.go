package managers

import "github.com/koodinikkarit/seppo/models"

func FindNewestVariationVersion(
	variationVersions []*models.VariationVersion,
) *models.VariationVersion {
	var newest *models.VariationVersion
	for _, variationVersion := range variationVersions {
		if newest == nil ||
			variationVersion.Version > newest.Version {
			newest = variationVersion
		}
	}
	return newest
}
