package managers

import (
	"database/sql"

	"github.com/koodinikkarit/seppo/models"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

func CreateAuthorByName(
	tx *sql.Tx,
	name string,
) *models.Author {
	author, _ := models.Authors(
		tx,
		Where("authors.name = ?", name),
	).One()

	if author != nil {
		return author
	}

	newAuthor := models.Author{
		Name: name,
	}

	newAuthor.Insert(tx)
	return &newAuthor
}
