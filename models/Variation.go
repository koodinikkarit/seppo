package models

import (
	"time"
)

type Variation struct {
	ID          uint32
	SongID      *uint32
	LanguageID  *uint32
	AuthorID    *uint32
	CopyrightID *uint32
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time

	Song              Song
	Language          Language
	Author            Author
	Copyright         Copyright
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
		if newest.ID == 0 ||
			variationVersion.Version > newest.Version {
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
		if variation.Name == name &&
			variation.Text == text {
			return &variation
		}
	}
	return nil
}

func (v *Variation) FindVersionWithVersionNumber(
	versionNumber uint32,
) *VariationVersion {
	for _, variationVersion := range v.VariationVersions {
		if variationVersion.Version == versionNumber {
			return &variationVersion
		}
	}
	return nil
}

func (v *Variation) HasTag(
	tagID uint32,
) bool {
	for _, tagVariation := range v.TagVariations {
		if tagVariation.TagID == tagID {
			return true
		}
	}
	return false
}
