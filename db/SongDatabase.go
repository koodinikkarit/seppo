package db

import "time"

type SongDatabase struct {
	ID        uint32
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time

	SongDatabaseTags []SongDatabaseTag
	Tags             []Tag       `gorm:"many2many:tag_variations"`
	Variations       []Variation `gorm:"many2many:song_database_variations"`
}

func (sdb *SongDatabase) HasSongDatabaseTag(
	tagID uint32,
) bool {
	for _, songDatabaseTag := range sdb.SongDatabaseTags {
		if songDatabaseTag.TagID == tagID {
			return true
		}
	}
	return false
}
