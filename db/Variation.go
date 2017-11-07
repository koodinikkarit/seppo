package db

import (
	"time"
)

type Variation struct {
	ID              uint32
	SongID          *uint32
	LanguageID      *uint32
	EwSongID        *uint32
	JyvaskylaSongID *uint32
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	DeletedAt       *time.Time

	Song              Song
	Language          Language
	VariationVersions []VariationVersion
	SongDatabases     []SongDatabase `gorm:"many2many:song_database_variations"`
	TagVariations     []TagVariation
	Tags              []Tag `gorm:"many2many:tag_variations"`
}

func (v *Variation) FindSongDatabaseByID(
	ID uint32,
) *SongDatabase {
	var songDatabase *SongDatabase
	for i := 0; i < len(v.SongDatabases); i++ {
		if ID == v.SongDatabases[i].ID {
			songDatabase = &v.SongDatabases[i]
		}
	}
	return songDatabase
}

func (v *Variation) FindNewestVersion() VariationVersion {
	var newest VariationVersion
	for _, variationVersion := range v.VariationVersions {
		if newest.ID == 0 || variationVersion.Version > newest.Version {
			newest = variationVersion
		}
	}

	return newest
}

func (v *Variation) FindVariationVersionByNameAndText(
	name string,
	text string,
) *VariationVersion {
	for _, variation := range v.VariationVersions {
		if variation.Name == name && variation.Text == text {
			return &variation
		}
	}
	return nil
}
