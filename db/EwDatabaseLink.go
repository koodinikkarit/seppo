package db

import "time"

type EwDatabaseLink struct {
	ID                 uint32
	EwDatabaseID       uint32
	EwDatabaseSongID   uint32
	VariationVersionID uint32
	Version            uint64
	CreatedAt          *time.Time
}
