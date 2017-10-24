package db

import "time"

type EwDatabase struct {
	ID             uint32
	Name           string
	SongDatabaseID uint32
	Key            string
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
}
