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
	return &DatabaseService{
		dbUser:                                 dbUser,
		dbPass:                                 dbPass,
		dbIP:                                   dbIP,
		dbPort:                                 dbPort,
		dbName:                                 dbName,
		CreateSongChannel:                      make(chan CreateSongInput),
		createVariationChannel:                 make(chan createVariationInternalInput),
		editVariationChannel:                   make(chan editVariationInternalInput),
		removeVariationChannel:                 make(chan removeVariationInternalInput),
		createSongDatabaseChannel:              make(chan createSongDatabaseInternalInput),
		editSongDatabaseChannel:                make(chan editSongDatabaseInternalInput),
		removeSongDatabaseChannel:              make(chan removeSongDatabaseInternalInput),
		createEwDatabaseChannel:                make(chan createEwDatabaseInternalInput),
		editEwDatabaseChannel:                  make(chan editEwDatabaseInternalInput),
		removeEwDatabaseChannel:                make(chan removeEwDatabaseInternalInput),
		createEwDatabaseLinkChannel:            make(chan createEwDatabaseLinkInternalInput),
		editEwDatabaseLinkChannel:              make(chan editEwDatabaseLinkInternalInput),
		removeEwDatabaseLinkChannel:            make(chan removeEwDatabaseLinkInternalInput),
		addVariationToSongDatabaseChannel:      make(chan addVariationToSongDatabaseInternalInput),
		removeVariationFromSongDatabaseChannel: make(chan removeVariationFromSongDatabaseInternalInput),
		removeEwSongChannel:                    make(chan removeEwSongInternalInput),
		createTagChannel:                       make(chan createTagInternalInput),
		editTagChannel:                         make(chan editTagInternalInput),
		removeTagChannel:                       make(chan removeTagInternalInput),
		createLanguageChannel:                  make(chan createLanguageInternalInput),
		editLanguageChannel:                    make(chan editLanguageInternalInput),
		removeLanguageChannel:                  make(chan removeLanguageInternalInput),
		addTagToVariationChannel:               make(chan addTagToVariationInternalInput),
		removeTagFromVariationChannel:          make(chan removeTagFromVariationInternalInput),
		addTagToSongDatabaseChannel:            make(chan addTagToSongDatabaseInternalInput),
		removeTagFromSongDatabaseChannel:       make(chan removeTagFromSongDatabaseInternalInput),
	}
}
