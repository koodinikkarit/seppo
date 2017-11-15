package db

type SrEwDatabaseLink struct {
	ID               uint32
	SrID             uint32
	EwDatabaseID     uint32
	EwDatabaseSongID uint32
	VariationID      uint32
	Version          uint32
	Author           string
	Copyright        string
	Operation        bool
}
