package SeppoDB

import (
	"fmt"

	mysql "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func CreateDb(
	dbUser string,
	dbPass string,
	dbIp string,
	dbPort string,
	dbName string,
) *DatabaseService {
	config, _ := mysql.ParseDSN("db")

	fmt.Println("addr", config.Addr)

	fmt.Println("CreateDb", dbUser, dbPass, dbIp, dbPort, dbName)
	fmt.Println("Connection string ", dbUser+":"+dbPass+"@tcp("+dbIp+":"+dbPort+")/"+dbName)
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
	removeVariationChannel := make(chan removeVariationInternalInput)
	createSongDatabaseChannel := make(chan createSongDatabaseInternalInput)
	editSongDatabaseChannel := make(chan editSongDatabaseInternalInput)
	removeSongDatabaseChannel := make(chan removeSongDatabaseInternalInput)
	createEwDatabaseChannel := make(chan createEwDatabaseInternalInput)
	editEwDatabaseChannel := make(chan editEwDatabaseInternalInput)
	removeEwDatabaseChannel := make(chan removeEwDatabaseInternalInput)
	addVariationToSongDatabaseChannel := make(chan addVariationToSongDatabaseInternalInput)
	removeVariationFromSongDatabaseChannel := make(chan removeVariationFromSongDatabaseInternalInput)

	return &DatabaseService{
		db:                                     db,
		CreateSongChannel:                      createSongChannel,
		createVariationChannel:                 createVariationChannel,
		editVariationChannel:                   editVariationChannel,
		removeVariationChannel:                 removeVariationChannel,
		createSongDatabaseChannel:              createSongDatabaseChannel,
		editSongDatabaseChannel:                editSongDatabaseChannel,
		removeSongDatabaseChannel:              removeSongDatabaseChannel,
		createEwDatabaseChannel:                createEwDatabaseChannel,
		editEwDatabaseChannel:                  editEwDatabaseChannel,
		removeEwDatabaseChannel:                removeEwDatabaseChannel,
		addVariationToSongDatabaseChannel:      addVariationToSongDatabaseChannel,
		removeVariationFromSongDatabaseChannel: removeVariationFromSongDatabaseChannel,
	}
}
