package models

import "time"

type SongDatabase struct {
	ID        uint32
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time

	SongDatabaseTags []SongDatabaseTag
	Tags             []Tag       `gorm:"many2many:tag_variations"`
	Variations       []Variation `gorm:"many2many:song_database_variations"`
}