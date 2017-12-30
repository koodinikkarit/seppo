package models

import "time"

type SongDatabaseTag struct {
	ID             uint32
	TagID          uint32
	SongDatabaseID uint32
	CreatedAt      *time.Time
	DeletedAt      *time.Time
}
