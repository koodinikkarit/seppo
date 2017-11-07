package help

import (
	"github.com/koodinikkarit/seppo/db"
)

func FindFirstVariationVersionWithNameOrText(
	variations []db.VariationVersion,
	name string,
	text string,
) *db.VariationVersion {
	var variationVersion *db.VariationVersion

	for _, v := range variations {
		if v.Name == name && v.Text == text {
			variationVersion = &v
			break
		}
	}

	return variationVersion
}
