package db

type SrEwConflict struct {
	ID                 uint32
	SrID               uint32
	VariationVersionID uint32
	EwDatabaseID       uint32
	EwSongID           uint32
	Name               string
	Text               string
	Resolved           int
}
