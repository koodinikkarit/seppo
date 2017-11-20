package managers

import (
	"database/sql"

	"github.com/koodinikkarit/seppo/matias_service"
	"github.com/koodinikkarit/seppo/models"
	null "gopkg.in/volatiletech/null.v6"
)

func NewVariationFromEwSong(
	tx *sql.Tx,
	ewSong *MatiasService.EwSong,
) (
	*models.Variation,
	*models.VariationVersion,
) {
	newVariation := &models.Variation{}
	if ewSong.Author != "" {
		newAuthor := CreateAuthorByName(
			tx,
			ewSong.Author,
		)
		newVariation.AuthorID = null.NewUint64(newAuthor.ID, true)
	}

	if ewSong.Copyright != "" {
		newCopyright := CreateCopyrightByName(
			tx,
			ewSong.Copyright,
		)
		newVariation.CopyrightID = null.NewUint64(newCopyright.ID, true)
	}
	newVariation.Insert(tx)
	newVariationVersion := &models.VariationVersion{
		Name: ewSong.Title,
		Text: ewSong.Text,
	}
	newVariation.AddVariationVersions(
		tx,
		true,
		newVariationVersion,
	)

	return newVariation, newVariationVersion
}
