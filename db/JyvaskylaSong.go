package db

import "time"

type JyvaskylaSong struct {
	ID             uint32
	AddedAt        uint64
	AddedBy        string
	AdditionalInfo string
	ArrangementBy  string
	ComposedBy     string
	Copyright      string
	Deleted        bool
	GlobalID       uint32
	LyricsBy       string
	Modified       *time.Time
	Name           string
	OrigName       string
	Song           string
	SongBookID     uint32
	TranslatedBy   string
	Year           string

	Variation Variation
}
