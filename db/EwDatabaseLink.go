package SeppoDB

type EwDatabaseLink struct {
	ID               uint32
	EwDatabaseID     uint32
	EwDatabaseSongID uint32
	VariationID      uint32
	Version          uint64

	EwDatabase EwDatabase
	Variation  Variation
}
