package db

import "time"

type EwDatabase struct {
	ID                             uint32
	Name                           string
	SongDatabaseID                 uint32
	EwDatabaseKey                  string
	RemoveSongsFromEwDatabase      bool
	RemoveSongsFromSongDatabase    bool
	VariationVersionConflictAction uint32
	CreatedAt                      *time.Time
	UpdatedAt                      *time.Time
	DeletedAt                      *time.Time

	SongDatabase    SongDatabase
	EwDatabaseLinks []EwDatabaseLink
	Variations      []Variation `gorm:"many2many:ew_database_links"`
}

func (ew *EwDatabase) HasVariation(
	variationID uint32,
) (
	bool,
	*EwDatabaseLink,
) {
	for i := 0; i < len(ew.EwDatabaseLinks); i++ {
		if ew.EwDatabaseLinks[i].VariationID == variationID {
			return true, &ew.EwDatabaseLinks[i]
		}
	}
	return false, nil
}

func (ew *EwDatabase) FindEwDatabaseLinkByEwSongID(
	ewSongID uint32,
) *EwDatabaseLink {
	for _, ewDatabaseLink := range ew.EwDatabaseLinks {
		if ewDatabaseLink.EwDatabaseSongID == ewSongID {
			return &ewDatabaseLink
		}
	}
	return nil
}
