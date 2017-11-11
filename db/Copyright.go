package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Copyright struct {
	ID        uint32
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func CreateCopyrightByName(
	tx *gorm.DB,
	name string,
) Copyright {
	var copyright Copyright
	tx.Where("name = ?", name).First(&copyright)
	if copyright.ID == 0 {
		copyright.Name = name
		tx.Create(&copyright)
	}
	return copyright
}
