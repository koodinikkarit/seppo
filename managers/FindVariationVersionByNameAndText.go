package managers

import "github.com/koodinikkarit/seppo/db"

func FindVariationVersionByNameAndText(
	variationVersions []*db.VariationVersion,
	name string,
	text string,
) *db.VariationVersion {
	for _, variationVersion := range variationVersions {
		if name != variationVersion.Name ||
			text != variationVersion.Text {
			continue
		}
		return variationVersion
	}
	return nil
}
