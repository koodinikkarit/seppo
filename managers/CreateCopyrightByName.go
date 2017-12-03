package managers

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
)

func CreateCopyrightByName(
	tx *gorm.DB,
	name string,
) db.Copyright {
	var copyright db.Copyright
	tx.Where("name = ?", name).
		First(&copyright)

	if copyright.ID > 0 {
		return copyright
	}

	copyright.Name = name
	tx.Create(&copyright)
	return copyright
}
