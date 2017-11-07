package help

import (
	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/matias_service"
)

func NewVariationFromEwSong(
	ewSong *MatiasService.EwSong,
) *db.Variation {
	var newVariation db.Variation
	newVariation.VariationVersions = append(
		newVariation.VariationVersions,
		db.VariationVersion{
			Name:    ewSong.Title,
			Text:    ewSong.Text,
			Version: 1,
		},
	)
	return &newVariation
}
