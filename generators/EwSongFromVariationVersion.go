package generators

import (
	"github.com/koodinikkarit/seppo/matias_service"
	"github.com/koodinikkarit/seppo/models"
)

func NewEwSongFromVariationVersion(
	id uint32,
	variationVersion *models.VariationVersion,
) *MatiasService.EwSong {
	return &MatiasService.EwSong{
		Id:          id,
		Title:       variationVersion.Name,
		Text:        variationVersion.Text,
		VariationId: uint32(variationVersion.VariationID),
	}
}
