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
	removeEwSongChannel                    chan removeEwSongInternalInput
	createTagChannel                       chan createTagInternalInput
	editTagChannel                         chan editTagInternalInput
	removeTagChannel                       chan removeTagInternalInput
	createLanguageChannel                  chan createLanguageInternalInput
	editLanguageChannel                    chan editLanguageInternalInput
	removeLanguageChannel                  chan removeLanguageInternalInput
	addTagToVariationChannel               chan addTagToVariationInternalInput
	removeTagFromVariationChannel          chan removeTagFromVariationInternalInput
	addTagToSongDatabaseChannel            chan addTagToSongDatabaseInternalInput
	removeTagFromSongDatabaseChannel       chan removeTagFromSongDatabaseInternalInput
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
		Migrate(db)
		ds.db = db
	}
	return ds.db.Debug()
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
		case in := <-ds.createVariationChannel:
			variation := &Variation{}
			if in.input.Name != "" && in.input.Text != "" {
				ds.GetDb().Table("variations").
					Joins("JOIN variation_texts ON variation_texts.variation_id = variations.id").
					Where("variations.name = ?", in.input.Name).
					Where("variation_texts.text = ?", in.input.Text).
					First(&variation)
			}
			if variation.ID == 0 {
				variation := &Variation{
					Name:    in.input.Name,
					Version: 1,
				}
				ds.GetDb().Create(&variation)
				variationText := &VariationText{
					VariationID: variation.ID,
					Text:        in.input.Text,
				}
				ds.GetDb().Create(&variationText)
			}
			in.returnChannel <- variation
		case in := <-ds.removeVariationChannel:
			var variation Variation
			ds.GetDb().First(&variation, in.variationID)
			ds.GetDb().Delete(&variation)
			ds.GetDb().Where("variation_id = ?", in.variationID).Delete(SongDatabaseVariation{})
			ds.GetDb().Where("variation_id = ?", in.variationID).Delete(VariationEwSongData{})
			ds.GetDb().Where("variation_id = ?", in.variationID).Delete(VariationText{})
			ds.GetDb().Where("variation_id = ?", in.variationID).Delete(TagVariation{})

			in.returnChannel <- true
		case in := <-ds.editVariationChannel:
			var variation Variation
			ds.GetDb().First(&variation, in.input.VariationID)

			changed := false

			if in.input.Name != "" && in.input.Name != variation.Name {
				changed = true
				variation.Name = in.input.Name
			}

			if in.input.Text != "" {
				var variationText VariationText
				ds.GetDb().Where("variation_id = ?", in.input.VariationID).First(&variationText)
				if variationText.ID > 0 {
					if variationText.Text != in.input.Text {
						changed = true
						variationText.Text = in.input.Text
						ds.GetDb().Save(&variationText)
					}
				}
			}

			if in.input.SongID != 0 && in.input.SongID != variation.SongID {
				changed = true
				variation.SongID = in.input.SongID
			}

			if in.input.LanguageID != 0 && in.input.LanguageID != variation.LanguageID {
				changed = true
				variation.LanguageID = in.input.LanguageID
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
		case in := <-ds.removeSongDatabaseChannel:
			var songDatabase SongDatabase
			success := false
			ds.GetDb().First(&songDatabase, in.songDatabaseID)
			if songDatabase.ID > 0 {
				ds.GetDb().Where("song_database_id = ?", in.songDatabaseID).Delete(EwDatabase{})
				ds.GetDb().Delete(&songDatabase)
				success = true
			}
			in.returnChannel <- success
		case in := <-ds.createEwDatabaseChannel:

			s, _ := GenerateRandomString(10)

			ewDatabase := &EwDatabase{
				Name:           in.input.Name,
				SongDatabaseID: in.input.SongDatabaseId,
				Key:            s,
			}
			ds.GetDb().Create(&ewDatabase)
			in.returnChnnel <- ewDatabase
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
			ds.GetDb().First(&ewDatabaseLink, in.input.EwDatabaseLinkID)
			if ewDatabaseLink.ID > 0 {
				if ewDatabaseLink.Version > 0 {
					ewDatabaseLink.Version = in.input.Version
				}
				if ewDatabaseLink.EwDatabaseSongID > 0 {
					ewDatabaseLink.EwDatabaseSongID = in.input.EwDatabaseSongID
				}
				ds.GetDb().Save(&ewDatabaseLink)
			}
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
				ds.GetDb().Create(&songDatabaseVariation)
			}
			in.returnChannel <- &songDatabaseVariation
		case in := <-ds.removeVariationFromSongDatabaseChannel:
			var ewDatabase EwDatabase
			ds.GetDb().Where("song_database_id = ?", in.songDatabaseID).First(&ewDatabase)
			//ds.GetDb().Where("ew_database_id = ?", ewDatabase.ID).Where("variation_id = ?", in.variationID).Delete(&EwDatabaseLink{})
			ds.GetDb().Where("song_database_id =  (?)", in.songDatabaseID).Where("variation_id = ?", in.variationID).Delete(SongDatabaseVariation{})
			in.returnChannel <- true
		case in := <-ds.removeEwSongChannel:
			var ewDatabaseLink EwDatabaseLink
			ds.GetDb().Where("ew_database_song_id = ?", in.ewSongID).First(&ewDatabaseLink)
			if ewDatabaseLink.ID > 0 {
				ds.GetDb().Delete(ewDatabaseLink)
				ds.GetDb().Where("song_database_id =  (?)", in.songDatabaseID).Where("variation_id = ?", ewDatabaseLink.VariationID).Delete(SongDatabaseVariation{})
				in.returnChannel <- true
			} else {
				in.returnChannel <- false
			}
		case in := <-ds.createTagChannel:
			tag := Tag{
				Name: in.input.Name,
			}
			ds.GetDb().Create(&tag)
			in.returnChannel <- &tag
		case in := <-ds.editTagChannel:
			var tag Tag
			ds.GetDb().First(&tag, in.input.TagID)
			if tag.ID > 0 {
				tag.Name = in.input.Name
			}
			ds.GetDb().Save(&tag)
			in.returnChannel <- &tag
		case in := <-ds.removeTagChannel:
			var tag Tag
			ds.GetDb().First(&tag, in.tagID)
			if tag.ID > 0 {
				ds.GetDb().Delete(&tag)
				in.returnChannel <- true
			} else {
				in.returnChannel <- false
			}
		case in := <-ds.createLanguageChannel:
			language := Language{
				Name: in.input.Name,
			}
			ds.GetDb().Create(&language)
			in.returnChannel <- &language
		case in := <-ds.editLanguageChannel:
			var language Language
			ds.GetDb().First(&language, in.input.LanguageID)
			if language.ID > 0 {
				language.Name = in.input.Name
			}
			ds.GetDb().Save(&language)
			in.returnChannel <- &language
		case in := <-ds.removeLanguageChannel:
			var language Language
			ds.GetDb().First(&language, in.languageID)
			if language.ID > 0 {
				ds.GetDb().Delete(&language)
				in.returnChannel <- true
			} else {
				in.returnChannel <- false
			}
		case in := <-ds.addTagToVariationChannel:
			var tagVariation TagVariation
			ds.GetDb().Where("tag_id = ?", in.tagID).
				Where("variation_id = ?", in.variationID)

			if tagVariation.ID == 0 {
				tagVariation = TagVariation{
					TagID:       in.tagID,
					VariationID: in.variationID,
				}
				ds.GetDb().Create(&tagVariation)
			}
			in.returnChannel <- &tagVariation
		case in := <-ds.removeTagFromVariationChannel:
			var tagVariation TagVariation
			ds.GetDb().Where("tag_id = ?", in.tagID).
				Where("variation_id = ?", in.variationID).
				Find(&tagVariation)

			if tagVariation.ID > 0 {
				ds.GetDb().Delete(&tagVariation)
				in.returnChannel <- true
			} else {
				in.returnChannel <- false
			}
		case in := <-ds.addTagToSongDatabaseChannel:
			var songDatabaseTag SongDatabaseTag
			ds.GetDb().Where("tag_id = ?", in.tagID).
				Where("song_database_id = ?", in.songDatabaseID).
				Find(&songDatabaseTag)
			if songDatabaseTag.ID == 0 {
				songDatabaseTag := SongDatabaseTag{
					TagID:          in.tagID,
					SongDatabaseID: in.songDatabaseID,
				}
				ds.GetDb().Create(&songDatabaseTag)
			}
			in.returnChannel <- &songDatabaseTag
		case in := <-ds.removeTagFromSongDatabaseChannel:
			var songDatabaseTag SongDatabaseTag
			ds.GetDb().Where("tag_id = ?", in.tagID).
				Where("song_database_id = ?", in.songDatabaseID).
				Find(&songDatabaseTag)
			if songDatabaseTag.ID > 0 {
				ds.GetDb().Delete(&songDatabaseTag)
				in.returnChannel <- true
			} else {
				in.returnChannel <- false
			}
		}
	}
}
