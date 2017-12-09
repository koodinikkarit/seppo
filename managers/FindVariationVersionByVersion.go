package managers

import "github.com/koodinikkarit/seppo/db"

func FindVariationVersionByVersion(
	variationVersions []*db.VariationVersion,
	versionNumber uint32,
) *db.VariationVersion {
	for _, variationVersion := range variationVersions {
		if variationVersion.Version == versionNumber {
			return variationVersion
		}
	}
	return nil
}
