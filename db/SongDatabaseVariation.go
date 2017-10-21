package db

import "time"

type SongDatabaseVariation struct {
	ID                 uint32
	SongDatabaseID     uint32
	VariationVersionID uint32
	CreatedAt          *time.Time
}
