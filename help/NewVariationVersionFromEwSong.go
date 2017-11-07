package help

import (
	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/matias_service"
)

func NewVariationVersionFromEwSong(
	ewSong *MatiasService.EwSong,
	variationID uint32,
	version uint32,
) *db.VariationVersion {
	newVariationVersion := &db.VariationVersion{
		VariationID: variationID,
		Name:        ewSong.Title,
		Text:        ewSong.Text,
		Version:     version,
	}

	return newVariationVersion
}
