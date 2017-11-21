package managers

import (
	"database/sql"

	"github.com/koodinikkarit/seppo/matias_service"
	"github.com/koodinikkarit/seppo/models"
)

func NewEwSongFromVariation(
	tx *sql.Tx,
	variation *models.Variation,
	variationVersion *models.VariationVersion,
) *MatiasService.EwSong {
	ewSong := MatiasService.EwSong{
		Title:       variationVersion.Name,
		Text:        variationVersion.Text,
		VariationId: uint32(variation.ID),
	}

	author, _ := variation.Author(tx).One()
	if author != nil {
		ewSong.Author = author.Name
	}
	copyright, _ := variation.Copyright(tx).One()
	if copyright != nil {
		ewSong.Copyright = copyright.Name
	}
	return &ewSong
}
