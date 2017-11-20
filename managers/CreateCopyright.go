package managers

import (
	"database/sql"

	"github.com/koodinikkarit/seppo/models"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

func CreateCopyrightByName(
	tx *sql.Tx,
	name string,
) *models.Copyright {
	copyright, _ := models.Copyrights(
		tx,
		Where("copyrights.name = ?", name),
	).One()

	if copyright != nil {
		return copyright
	}

	newCopyright := models.Copyright{
		Name: name,
	}

	newCopyright.Insert(tx)
	return &newCopyright
}
