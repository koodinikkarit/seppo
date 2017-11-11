package db

func CreateNewVariationAndVersion(
	name string,
	text string,
) Variation {
	newVariation := Variation{}
	newVariation.VariationVersions = append(
		newVariation.VariationVersions,
		VariationVersion{
			Name:    name,
			Text:    text,
			Version: 1,
		},
	)

	return newVariation
}
