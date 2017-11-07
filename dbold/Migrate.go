package SeppoDB

import (
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&EwDatabase{})
	db.AutoMigrate(&EwDatabaseLink{})
	db.AutoMigrate(&Song{})
	db.AutoMigrate(&SongDatabase{})
	db.AutoMigrate(&SongDatabaseVariation{})
	db.AutoMigrate(&Variation{})
	db.AutoMigrate(&VariationEwSongData{})
	db.AutoMigrate(&VariationText{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Language{})
	db.AutoMigrate(&SongDatabaseTag{})
	db.AutoMigrate(&SongDatabaseVariation{})
	db.AutoMigrate(&TagVariation{})
	db.AutoMigrate(&Schedule{})
	db.AutoMigrate(&ScheduleVariation{})
}