package SeppoDB

import (
	_ "github.com/go-sql-driver/mysql"
)

func NewDatabaseService(
	dbUser string,
	dbPass string,
	dbIP string,
	dbPort string,
	dbName string,
) *DatabaseService {

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
	createEwDatabaseLinkChannel := make(chan createEwDatabaseLinkInternalInput)
	editEwDatabaseLinkChannel := make(chan editEwDatabaseLinkInternalInput)
	removeEwDatabaseLinkChannel := make(chan removeEwDatabaseLinkInternalInput)
	addVariationToSongDatabaseChannel := make(chan addVariationToSongDatabaseInternalInput)
	removeVariationFromSongDatabaseChannel := make(chan removeVariationFromSongDatabaseInternalInput)
	removeEwSongChannel := make(chan removeEwSongInternalInput)

	return &DatabaseService{
		dbUser:                                 dbUser,
		dbPass:                                 dbPass,
		dbIP:                                   dbIP,
		dbPort:                                 dbPort,
		dbName:                                 dbName,
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
		createEwDatabaseLinkChannel:            createEwDatabaseLinkChannel,
		editEwDatabaseLinkChannel:              editEwDatabaseLinkChannel,
		removeEwDatabaseLinkChannel:            removeEwDatabaseLinkChannel,
		addVariationToSongDatabaseChannel:      addVariationToSongDatabaseChannel,
		removeVariationFromSongDatabaseChannel: removeVariationFromSongDatabaseChannel,
		removeEwSongChannel:                    removeEwSongChannel,
	}
}
