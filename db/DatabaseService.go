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
	createEwDatabaseLinkChannel            chan createEwDatabaseLinkInternalInput
	editEwDatabaseLinkChannel              chan editEwDatabaseLinkInternalInput
	removeEwDatabaseLinkChannel            chan removeEwDatabaseLinkInternalInput
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
		db.AutoMigrate(&VariationEwSongData{})
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
				Name:    createVariationInput.input.Name,
				Text:    createVariationInput.input.Text,
				Version: 1,
			}

			ds.GetDb().Create(&variation)

			createVariationInput.returnChannel <- variation
		case in := <-ds.removeVariationChannel:
			var variation Variation
			ds.GetDb().First(&variation, in.variationID)
			ds.GetDb().Delete(&variation)
			ds.GetDb().Where("variation_id = ?", in.variationID).Delete(EwDatabaseLink{})
			ds.GetDb().Where("variation_id", in.variationID).Delete(SongDatabaseVariation{})
			ds.GetDb().Where("variation_id", in.variationID).Delete(VariationEwSongData{})

			in.returnChannel <- true
		case in := <-ds.editVariationChannel:
			var variation Variation
			ds.GetDb().First(&variation, in.input.VariationID)

			changed := false

			if in.input.Name != "" && in.input.Name != variation.Name {
				changed = true
				variation.Name = in.input.Name
			}

			if in.input.Text != "" && in.input.Text != variation.Text {
				changed = true
				variation.Text = in.input.Text
			}

			if in.input.SongID != 0 && in.input.SongID != variation.SongID {
				changed = true
				variation.SongID = in.input.SongID
			}

			if changed == true {
				variation.Version++
			}

			ds.GetDb().Save(&variation)
			in.returnChannel <- &variation
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
		case in := <-ds.createEwDatabaseLinkChannel:
			ewDatabaseLink := &EwDatabaseLink{
				EwDatabaseID:     in.ewDatabaseID,
				EwDatabaseSongID: in.ewDatabaseSongID,
				VariationID:      in.variationID,
				Version:          in.version,
			}
			ds.GetDb().Create(&ewDatabaseLink)
			in.returnChannel <- ewDatabaseLink
		case in := <-ds.editEwDatabaseLinkChannel:
			var ewDatabaseLink EwDatabaseLink
			ds.GetDb().First(&ewDatabaseLink, in.ewDatabaseLinkID)
			if ewDatabaseLink.Version > 0 {
				ewDatabaseLink.Version = in.version
			}
			ds.GetDb().Save(&ewDatabaseLink)
			in.returnChannel <- &ewDatabaseLink
		case in := <-ds.removeEwDatabaseLinkChannel:
			var ewDatabaseLink EwDatabaseLink
			ds.GetDb().First(&ewDatabaseLink, in.ewDatabaseLinkID)
			if ewDatabaseLink.ID > 0 {
				ds.GetDb().Delete(&ewDatabaseLink)
				in.returnChnnel <- true
			} else {
				in.returnChnnel <- false
			}
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
