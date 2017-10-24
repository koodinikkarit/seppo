package db

import "time"

type SongDatabase struct {
	ID        uint32
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
