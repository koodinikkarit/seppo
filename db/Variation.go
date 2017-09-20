package SeppoDB

import "github.com/jinzhu/gorm"

type Variation struct {
	gorm.Model

	ID         uint32
	Name       string
	SongID     uint32
	Version    uint64
	LanguageID uint32
	Origin     string

	Song     Song
	Language Language
}
