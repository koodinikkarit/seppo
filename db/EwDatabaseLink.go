package db

import "time"

type EwDatabaseLink struct {
	ID               uint32
	EwDatabaseID     uint32
	EwDatabaseSongID uint32
	VariationID      uint32
	Version          uint32
	Author           string
	Copyright        string
	CreatedAt        *time.Time
	UpdatedAt        *time.Time

	EwDatabase EwDatabase
	Variation  Variation
}
