package help

import (
	"github.com/koodinikkarit/seppo/db"
)

func FindNewestVariationVersionFromSlice(
	variationVersions []db.VariationVersion,
) db.VariationVersion {
	newest := variationVersions[0]

	for i := 1; i < len(variationVersions); i++ {
		if variationVersions[i].Version > newest.Version {
			newest = variationVersions[i]
		}
	}

	return newest
}
