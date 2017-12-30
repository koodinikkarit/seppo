package models

import "time"

type EwDatabase struct {
	ID                             uint32
	Name                           string
	SongDatabaseID                 uint32
	FilesystemPath                 string
	MatiasClientID                 uint32
	RemoveSongsFromEwDatabase      bool
	RemoveSongsFromSongDatabase    bool
	VariationVersionConflictAction uint32
	CreatedAt                      *time.Time
	UpdatedAt                      *time.Time
	DeletedAt                      *time.Time

	SongDatabase SongDatabase
	Variations   []Variation `gorm:"many2many:ew_database_links"`
}
