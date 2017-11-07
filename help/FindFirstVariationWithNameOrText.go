package help

import (
	"github.com/koodinikkarit/seppo/db"
)

func FindFirstVariationWithNameOrText(
	variations []db.Variation,
	name string,
	text string,
) (
	*db.Variation,
	*db.VariationVersion,
) {
	var variationVersion *db.VariationVersion
	var variation *db.Variation
	for _, v := range variations {
		variationVersion = FindFirstVariationVersionWithNameOrText(
			v.VariationVersions,
			name,
			text,
		)
		if variationVersion != nil {
			variation = &v
			break
		}
	}
	return variation, variationVersion
}
