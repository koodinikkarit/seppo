package db

import "time"

type SongDatabaseVariation struct {
	ID             uint32
	SongDatabaseID uint32
	VariationID    uint32
	CreatedAt      *time.Time

	SongDatabase SongDatabase
	Variation    Variation
}
