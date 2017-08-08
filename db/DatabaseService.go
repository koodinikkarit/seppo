package SeppoDB

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type DatabaseService struct {
	db                                     *gorm.DB
	dbUser                                 string
	dbPass                                 string
	dbIP                                   string
	dbPort                                 string
	dbName                                 string
	CreateSongChannel                      chan CreateSongInput
	createVariationChannel                 chan createVariationInternalInput
	editVariationChannel                   chan editVariationInternalInput
	removeVariationChannel                 chan removeVariationInternalInput
	createSongDatabaseChannel              chan createSongDatabaseInternalInput
	editSongDatabaseChannel                chan editSongDatabaseInternalInput
	removeSongDatabaseChannel              chan removeSongDatabaseInternalInput
	createEwDatabaseChannel                chan createEwDatabaseInternalInput
	editEwDatabaseChannel                  chan editEwDatabaseInternalInput
	removeEwDatabaseChannel                chan removeEwDatabaseInternalInput
	addVariationToSongDatabaseChannel      chan addVariationToSongDatabaseInternalInput
	removeVariationFromSongDatabaseChannel chan removeVariationFromSongDatabaseInternalInput
}

func (ds *DatabaseService) insertSong(name string, songID uint32) {

}

func (ds *DatabaseService) GetDb() *gorm.DB {
	if ds.db == nil {
		fmt.Println("CreateDb", ds.dbUser, ds.dbPass, ds.dbIP, ds.dbPort, ds.dbName)
		fmt.Println("Connection string ", ds.dbUser+":"+ds.dbPass+"@tcp("+ds.dbIP+":"+ds.dbPort+")/"+ds.dbName)
		db, err := gorm.Open("mysql", ds.dbUser+":"+ds.dbPass+"@tcp("+ds.dbIP+":"+ds.dbPort+")/"+ds.dbName)
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
		ds.db = db
	}
	return ds.db
}

func (ds *DatabaseService) Start() {
	for {
		select {
		case createSongInput := <-ds.CreateSongChannel:
			variation := &Variation{
				Name:    createSongInput.Name,
				Version: 1,
			}

			ds.GetDb().Create(&variation)

			if createSongInput.EwDatabaseId != 0 {
				ewDatabaseLink := &EwDatabaseLink{
					EwDatabaseID:     createSongInput.EwDatabaseId,
					EwDatabaseSongID: createSongInput.SongID,
					VariationID:      variation.ID,
				}

				ds.GetDb().Create(&ewDatabaseLink)

				fmt.Println("uusi ewsong")
			}
			fmt.Println("uusi laulu", createSongInput)
		case createVariationInput := <-ds.createVariationChannel:
			variation := &Variation{
				Name: createVariationInput.input.Name,
				Text: createVariationInput.input.Text,
			}

			ds.GetDb().Create(&variation)

			createVariationInput.returnChannel <- variation
		case removeVariation := <-ds.removeVariationChannel:
			var variation Variation
			ds.GetDb().First(&variation, removeVariation.variationID)

			ds.GetDb().Delete(&variation)
			removeVariation.returnChannel <- true
		case editVariationInput := <-ds.editVariationChannel:
			var variation Variation
			ds.GetDb().First(&variation, editVariationInput.input.VariationID)

			if editVariationInput.input.Name != "" {
				variation.Name = editVariationInput.input.Name
			}

			if editVariationInput.input.Text != "" {
				variation.Text = editVariationInput.input.Text
			}

			if editVariationInput.input.SongID != 0 {
				variation.SongID = editVariationInput.input.SongID
			}

			ds.GetDb().Save(&variation)
			editVariationInput.returnChannel <- &variation
		case createSongDatabase := <-ds.createSongDatabaseChannel:
			songDatabase := &SongDatabase{
				Name: createSongDatabase.input.Name,
			}
			ds.GetDb().Create(&songDatabase)
			createSongDatabase.returnChannel <- songDatabase
		case editSongDatabase := <-ds.editSongDatabaseChannel:
			var songDatabase SongDatabase
			ds.GetDb().First(&songDatabase, editSongDatabase.input.SongDatabaseId)

			if editSongDatabase.input.Name != "" {
				songDatabase.Name = editSongDatabase.input.Name
			}

			ds.GetDb().Save(&songDatabase)
			editSongDatabase.returnChannel <- &songDatabase
		case removeSongDatabase := <-ds.removeSongDatabaseChannel:
			var songDatabase SongDatabase
			ds.GetDb().First(&songDatabase, removeSongDatabase.songDatabaseID)

			ds.GetDb().Delete(&songDatabase)
			removeSongDatabase.returnChannel <- true
		case createEwDatabase := <-ds.createEwDatabaseChannel:
			ewDatabase := &EwDatabase{
				Name:           createEwDatabase.input.Name,
				SongDatabaseID: createEwDatabase.input.SongDatabaseId,
			}
			ds.GetDb().Create(&ewDatabase)
			createEwDatabase.returnChnnel <- ewDatabase
		case in := <-ds.editEwDatabaseChannel:
			var ewDatabase EwDatabase
			ds.GetDb().First(&ewDatabase, in.input.EwDatabaseID)
			if in.input.Name != "" {
				ewDatabase.Name = in.input.Name
			}
			if in.input.SongDatabaseID > 0 {
				ewDatabase.SongDatabaseID = in.input.SongDatabaseID
			}
			ds.GetDb().Save(&ewDatabase)
			in.returnChannel <- &ewDatabase
		case removeEwDatabase := <-ds.removeEwDatabaseChannel:
			var ewDatabase EwDatabase
			ds.GetDb().First(&ewDatabase, removeEwDatabase.ewDatabaseID)
			ds.GetDb().Delete(&ewDatabase)
			removeEwDatabase.returnChannel <- true
		case in := <-ds.addVariationToSongDatabaseChannel:
			var songDatabaseVariation SongDatabaseVariation
			ds.GetDb().Where("song_database_id = ?", in.songDatabaseID).Where("variation_id = ?", in.variationID).First(&songDatabaseVariation)
			if songDatabaseVariation.ID == 0 {
				songDatabaseVariation = SongDatabaseVariation{
					SongDatabaseID: in.songDatabaseID,
					VariationID:    in.variationID,
				}
			}

			ds.GetDb().Create(&songDatabaseVariation)
			in.returnChannel <- &songDatabaseVariation
		case in := <-ds.removeVariationFromSongDatabaseChannel:
			var songDatabaseVariation SongDatabaseVariation

			ds.GetDb().Where("song_database_id =  (?)", in.songDatabaseID).Where("variation_id = ?", in.variationID).Find(&songDatabaseVariation)
			ds.GetDb().Delete(&songDatabaseVariation)
			in.returnChannel <- true
		}
	}
}
