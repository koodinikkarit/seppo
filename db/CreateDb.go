package SeppoDB

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func CreateDb(
	dbUser string,
	dbPass string,
	dbIp string,
	dbPort string,
	dbName string,
) *DatabaseService {
	db, err := gorm.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbIp+":"+dbPort+")/"+dbName)
	if err != nil {
		fmt.Println("Creating database connection failed", err)
	}

	db.AutoMigrate(&EwDatabase{})
	db.AutoMigrate(&EwDatabaseLink{})
	db.AutoMigrate(&Song{})
	db.AutoMigrate(&SongDatabase{})
	db.AutoMigrate(&SongDatabaseVariation{})
	db.AutoMigrate(&Variation{})
	db.AutoMigrate(&Verse{})

	createSongChannel := make(chan CreateSongInput)
	createVariationChannel := make(chan createVariationInternalInput)
	editVariationChannel := make(chan editVariationInternalInput)

	return &DatabaseService{
		db:                     db,
		CreateSongChannel:      createSongChannel,
		createVariationChannel: createVariationChannel,
		editVariationChannel:   editVariationChannel,
	}
}