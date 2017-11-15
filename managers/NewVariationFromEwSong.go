package managers

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/matias_service"
)

func NewVariationFromEwSong(
	tx *gorm.DB,
	ewSong *MatiasService.EwSong,
) db.Variation {
	newVariation := db.Variation{}
	newVariation.VariationVersions = append(
		newVariation.VariationVersions,
		db.VariationVersion{
			Name:    ewSong.Title,
			Text:    ewSong.Text,
			Version: 1,
		},
	)
	if ewSong.Author != "" {
		newAuthor := db.CreateAuthorByName(
			tx,
			ewSong.Author,
		)
		newVariation.AuthorID = &newAuthor.ID
	}

	if ewSong.Copyright != "" {
		newCopyright := db.CreateCopyrightByName(
			tx,
			ewSong.Copyright,
		)
		newVariation.Copyright = newCopyright
	}

	tx.Create(&newVariation)
	return newVariation
}
