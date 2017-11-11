package db

import (
	"github.com/koodinikkarit/seppo/matias_service"
)

func NewVariationFromEwSong(
	ewSong *MatiasService.EwSong,
) *Variation {
	var newVariation Variation
	newVariation.VariationVersions = append(
		newVariation.VariationVersions,
		VariationVersion{
			Name:    ewSong.Title,
			Text:    ewSong.Text,
			Version: 1,
		},
	)
	return &newVariation
}
