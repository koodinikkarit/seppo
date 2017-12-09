package managers

import (
	"github.com/koodinikkarit/seppo/db"
)

func FindNewestVariationVersion(
	variationVersions []db.VariationVersion,
) db.VariationVersion {
	var newest db.VariationVersion
	for _, variationVersion := range variationVersions {
		if newest.ID == 0 ||
			variationVersion.Version > newest.Version {
			newest = variationVersion
		}
	}
	return newest
}
