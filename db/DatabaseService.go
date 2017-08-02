package SeppoDB

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type DatabaseService struct {
	db                                     *gorm.DB
	CreateSongChannel                      chan CreateSongInput
	createVariationChannel                 chan createVariationInternalInput
	editVariationChannel                   chan editVariationInternalInput
	removeVariationChannel                 chan removeVariationInternalInput
	createSongDatabaseChannel              chan createSongDatabaseInternalInput
	editSongDatabaseChannel                chan editSongDatabaseInternalInput
	removeSongDatabaseChannel              chan removeSongDatabaseInternalInput
	createEwDatabaseChannel                chan createEwDatabaseInternalInput
	removeEwDatabaseChannel                chan removeEwDatabaseInternalInput
	addVariationToSongDatabaseChannel      chan addVariationToSongDatabaseInternalInput
	removeVariationFromSongDatabaseChannel chan removeVariationFromSongDatabaseInternalInput
}

func (ds *DatabaseService) insertSong(name string, songID uint32) {

}

func (ds *DatabaseService) GetDb() *gorm.DB {
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

			ds.db.Create(&variation)

			if createSongInput.EwDatabaseId != 0 {
				ewDatabaseLink := &EwDatabaseLink{
					EwDatabaseID:     createSongInput.EwDatabaseId,
					EwDatabaseSongID: createSongInput.SongID,
					VariationID:      variation.ID,
				}

				ds.db.Create(&ewDatabaseLink)

				fmt.Println("uusi ewsong")
			}
			fmt.Println("uusi laulu", createSongInput)
		case createVariationInput := <-ds.createVariationChannel:
			variation := &Variation{
				Name: createVariationInput.input.Name,
				Text: createVariationInput.input.Text,
			}

			ds.db.Create(&variation)

			createVariationInput.returnChannel <- variation
		case removeVariation := <-ds.removeVariationChannel:
			var variation Variation
			ds.db.First(&variation, removeVariation.variationID)

			ds.db.Delete(&variation)
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

			ds.db.Save(&variation)
			editVariationInput.returnChannel <- &variation
		case createSongDatabase := <-ds.createSongDatabaseChannel:
			songDatabase := &SongDatabase{
				Name: createSongDatabase.input.Name,
			}
			ds.db.Create(&songDatabase)
			createSongDatabase.returnChannel <- songDatabase
		case editSongDatabase := <-ds.editSongDatabaseChannel:
			var songDatabase SongDatabase
			ds.db.First(&songDatabase, editSongDatabase.input.SongDatabaseId)

			if editSongDatabase.input.Name != "" {
				songDatabase.Name = editSongDatabase.input.Name
			}

			ds.db.Save(&songDatabase)
			editSongDatabase.returnChannel <- &songDatabase
		case removeSongDatabase := <-ds.removeSongDatabaseChannel:
			var songDatabase SongDatabase
			ds.db.First(&songDatabase, removeSongDatabase.songDatabaseID)

			ds.db.Delete(&songDatabase)
			removeSongDatabase.returnChannel <- true
		case createEwDatabase := <-ds.createEwDatabaseChannel:
			ewDatabase := &EwDatabase{
				Name:           createEwDatabase.input.Name,
				SongDatabaseID: createEwDatabase.input.SongDatabaseId,
			}
			ds.db.Create(&ewDatabase)
			createEwDatabase.returnChnnel <- ewDatabase
		case removeEwDatabase := <-ds.removeEwDatabaseChannel:
			var ewDatabase EwDatabase
			ds.db.First(&ewDatabase, removeEwDatabase.ewDatabaseID)
			ds.db.Delete(&ewDatabase)
			removeEwDatabase.returnChannel <- true
		case in := <-ds.addVariationToSongDatabaseChannel:
			var songDatabaseVariation SongDatabaseVariation
			ds.db.Where("song_database_id = ?", in.songDatabaseID).Where("variation_id = ?", in.variationID).First(&songDatabaseVariation)
			if songDatabaseVariation.ID == 0 {
				songDatabaseVariation = SongDatabaseVariation{
					SongDatabaseID: in.songDatabaseID,
					VariationID:    in.variationID,
				}
			}

			ds.db.Create(&songDatabaseVariation)
			in.returnChannel <- &songDatabaseVariation
		case in := <-ds.removeVariationFromSongDatabaseChannel:
			var songDatabaseVariation SongDatabaseVariation

			ds.db.Where("song_database_id =  (?)", in.songDatabaseID).Where("variation_id = ?", in.variationID).Find(&songDatabaseVariation)
			ds.db.Delete(&songDatabaseVariation)
			in.returnChannel <- true
		}
	}
}
