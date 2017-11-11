package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Author struct {
	ID        uint32
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func CreateAuthorByName(
	tx *gorm.DB,
	name string,
) Author {
	var author Author
	tx.Where("name = ?", name).
		First(&author)

	if author.ID == 0 {
		author.Name = name
		tx.Create(&author)
	}
	return author
}
