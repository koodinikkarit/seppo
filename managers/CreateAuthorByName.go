package managers

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
)

func CreateAuthorByName(
	tx *gorm.DB,
	name string,
) db.Author {
	var author db.Author
	tx.Where("name = ?", name).
		First(&author)
	if author.ID > 0 {
		return author
	}

	author.Name = name

	tx.Create(&author)
	return author
}
