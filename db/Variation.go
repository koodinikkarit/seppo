package db

import "time"

type Variation struct {
	ID         uint32
	SongID     *uint32
	LanguageID *uint32
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
}
