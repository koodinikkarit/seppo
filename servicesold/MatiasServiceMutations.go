package seppo

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/help"
	"github.com/koodinikkarit/seppo/matias_service"
)

// InsertEwSongIds Asettaa ewtietokantaan luodun laulun idn
func (s *MatiasServiceServer) InsertEwSongIds(
	ctx context.Context,
	in *MatiasService.InsertEwSongIdsRequest) (*MatiasService.InsertEwSongIdsResponse, error) {
	res := &MatiasService.InsertEwSongIdsResponse{}

	var variationIds []uint32
	for _, variationIdEwSongId := range in.VariationIdEwSongIds {
		variationIds = append(variationIds, variationIdEwSongId.VariationId)
	}

	variations := []SeppoDB.Variation{}
	s.databaseService.GetDb().Where("id in (?)", variationIds).Find(&variations)

	var ewDatabase SeppoDB.EwDatabase

	s.databaseService.GetDb().Where("ew_databases.key = ?", in.EwDatabaseKey).First(&ewDatabase)

	for _, variationIdEwSongId := range in.VariationIdEwSongIds {
		for _, variation := range variations {
			if variationIdEwSongId.VariationId == variation.ID {
				s.databaseService.CreateEwDatabaseLink(
					ewDatabase.ID,
					variationIdEwSongId.EwSongId,
					variation.ID,
					variation.Version,
				)
			}
		}
	}

	var ewSongIDs []uint32
	ewDatabaseLinks := []SeppoDB.EwDatabaseLink{}

	for _, link := range in.NewSongIds {
		ewSongIDs = append(ewSongIDs, link.OldEwSongId)
	}

	s.databaseService.GetDb().Where("ew_database_song_id in (?)", ewSongIDs).Find(&ewDatabaseLinks)

	for _, ewDatabaseLink := range ewDatabaseLinks {
		for i := 0; i < len(in.NewSongIds); i++ {
			if ewDatabaseLink.EwDatabaseSongID == in.NewSongIds[i].OldEwSongId {
				s.databaseService.EditEwDatabaseLink(SeppoDB.EditEwDatabaseLinkInput{
					EwDatabaseLinkID: ewDatabaseLink.ID,
					EwDatabaseSongID: ewDatabaseLink.EwDatabaseSongID,
				})
			}
		}
	}

	return res, nil
}

func (s *MatiasServiceServer) SyncEwDatabase(ctx context.Context, in *MatiasService.SyncEwDatabaseRequest) (*MatiasService.SyncEwDatabaseResponse, error) {
	response := &MatiasService.SyncEwDatabaseResponse{}

	var ewDatabase SeppoDB.EwDatabase
	s.databaseService.GetDb().Where("ew_databases.key = ?", in.EwDatabaseKey).Find(&ewDatabase)

	if ewDatabase.ID > 0 {

		ewDatabaseVariations := []SeppoDB.Variation{}
		variationTexts := []SeppoDB.VariationText{}

		s.databaseService.GetDb().Debug().Table("song_database_variations").
			Joins("JOIN variations ON variations.id = song_database_variations.variation_id").
			Where("song_database_variations.song_database_id = ?", ewDatabase.SongDatabaseID).
			Select("variations.id, variations.name, variations.song_id, variations.version").
			Find(&ewDatabaseVariations)
		s.databaseService.GetDb().Table("song_database_variations").
			Joins("JOIN variation_texts ON variation_texts.variation_id = song_database_variations.variation_id").
			Where("song_database_variations.song_database_id = ?", ewDatabase.SongDatabaseID).
			Select("variation_texts.id, variation_texts.variation_id, variation_texts.text").
			Find(&variationTexts)

		var ewDatabaseLinks []SeppoDB.EwDatabaseLink
		s.databaseService.GetDb().Debug().Where("ew_database_id = ?", ewDatabase.ID).Find(&ewDatabaseLinks)

		for _, ewSong := range in.EwSongs {
			ewDatabaseLink := help.FindEwDatabaseLinkWithEwSong
			foundEwDatabaseLink := false
			for _, ewdatabaseLink := range ewDatabaseLinks {
				if ewSong.Id == ewdatabaseLink.EwDatabaseSongID {
					foundEwDatabaseLink = true
					foundVariation := false
					for _, a := range ewDatabaseVariations {
						if ewdatabaseLink.VariationID == a.ID {
							foundVariation = true
							var text string
							for _, vt := range variationTexts {
								if vt.VariationID == a.ID {
									text = vt.Text
								}
							}
							if a.Name != ewSong.Title || text != ewSong.Text {
								if ewdatabaseLink.Version >= a.Version {
									v := s.databaseService.EditVariation(SeppoDB.EditVariationInput{
										VariationID: a.ID,
										Name:        ewSong.Title,
										Text:        ewSong.Text,
									})

									s.databaseService.EditEwDatabaseLink(SeppoDB.EditEwDatabaseLinkInput{
										EwDatabaseLinkID: ewdatabaseLink.ID,
										Version:          v.Version,
									})
								} else {
									response.EwSongs = append(response.EwSongs, &MatiasService.EwSong{
										Id:          ewSong.Id,
										VariationId: a.ID,
										Title:       a.Name,
										Text:        text,
									})
									s.databaseService.EditEwDatabaseLink(SeppoDB.EditEwDatabaseLinkInput{
										EwDatabaseLinkID: ewdatabaseLink.ID,
										Version:          a.Version,
									})
								}
							}
						}
					}
					if foundVariation == false {
						response.RemoveEwSongIds = append(response.RemoveEwSongIds, ewSong.Id)
						s.databaseService.RemoveDatabaseLink(ewdatabaseLink.ID)
					}
				}
			}
			if foundEwDatabaseLink == false {
				variation := s.databaseService.CreateVariation(SeppoDB.CreateVariationInput{
					Name: ewSong.Title,
					Text: ewSong.Text,
				})
				s.databaseService.CreateEwDatabaseLink(
					ewDatabase.ID,
					ewSong.Id,
					variation.ID,
					variation.Version,
				)
				s.databaseService.AddVariationToSongDatabase(ewDatabase.SongDatabaseID, variation.ID)
			}
		}

		for _, variation := range ewDatabaseVariations {
			var text string
			for _, vt := range variationTexts {
				if vt.VariationID == variation.ID {
					text = vt.Text
				}
			}
			var foundEwDatabaseLink bool
			var ewDatabaseLink2 SeppoDB.EwDatabaseLink
			var foundEwSong bool
			for _, ewDatabaseLink := range ewDatabaseLinks {
				if variation.ID == ewDatabaseLink.VariationID {
					foundEwDatabaseLink = true
					ewDatabaseLink2 = ewDatabaseLink
					for _, e := range in.EwSongs {
						if ewDatabaseLink.EwDatabaseSongID == e.Id {
							foundEwSong = true
							break
						}
					}

				}
			}
			if foundEwDatabaseLink == false {
				foundSameEwSong := false
				for _, ews := range in.EwSongs {
					if ews.Title == variation.Name && ews.Text == text {
						foundSameEwSong = true
					}
				}
				if foundSameEwSong == false {
					response.EwSongs = append(response.EwSongs, &MatiasService.EwSong{
						VariationId: variation.ID,
						Title:       variation.Name,
						Text:        text,
					})
				}
			} else {
				if foundEwSong == false {
					s.databaseService.RemoveEwSong(ewDatabase.SongDatabaseID, ewDatabaseLink2.EwDatabaseSongID)
				}
			}
		}
	}

	return response, nil
}
