package seppo

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
)

func (s *SeppoServiceServer) CreateVariation(ctx context.Context, in *SeppoService.CreateVariationRequest) (*SeppoService.CreateVariationResponse, error) {
	variation := s.databaseService.CreateVariation(SeppoDB.NewCreateVariationFromServiceType(in))
	return &SeppoService.CreateVariationResponse{
		Variation: NewVariationToServiceType(variation),
	}, nil
}

func (s *SeppoServiceServer) EditVariation(ctx context.Context, in *SeppoService.EditVariationRequest) (*SeppoService.EditVariationResponse, error) {
	res := &SeppoService.EditVariationResponse{}
	res.Variation = NewVariationToServiceType(s.databaseService.EditVariation(SeppoDB.NewEditVariationFromService(in)))
	return res, nil
}

func (s *SeppoServiceServer) RemoveVariation(ctx context.Context, in *SeppoService.RemoveVariationRequest) (*SeppoService.RemoveVariationResponse, error) {
	res := &SeppoService.RemoveVariationResponse{}
	s.databaseService.RemoveVariation(in.VariationId)
	return res, nil
}

func (s *SeppoServiceServer) CreateSongDatabase(ctx context.Context, in *SeppoService.CreateSongDatabaseRequest) (*SeppoService.CreateSongDatabaseResponse, error) {
	res := &SeppoService.CreateSongDatabaseResponse{}
	songDatabase := s.databaseService.CreateSongDatabase(SeppoDB.NewCreateSongDatabaseInputFromServiceType(in))
	res.SongDatabase = NewSongDatabaseToServiceType(songDatabase)
	return res, nil
}

func (s *SeppoServiceServer) EditSongDatabase(ctx context.Context, in *SeppoService.EditSongDatabaseRequest) (*SeppoService.EditSongDatabaseResponse, error) {
	res := &SeppoService.EditSongDatabaseResponse{}
	res.SongDatabase = NewSongDatabaseToServiceType(s.databaseService.EditSongDatabase(SeppoDB.NewEditSongDatabaseInputFromServiceType(in)))
	return res, nil
}

func (s *SeppoServiceServer) RemoveSongDatabase(ctx context.Context, in *SeppoService.RemoveSongDatabaseRequest) (*SeppoService.RemoveSongDatabaseResponse, error) {
	res := &SeppoService.RemoveSongDatabaseResponse{}
	s.databaseService.RemoveSongDatabase(in.SongDatabaseId)
	return res, nil
}

func (s *SeppoServiceServer) CreateEwDatabase(ctx context.Context, in *SeppoService.CreateEwDatabaseRequest) (*SeppoService.CreateEwDatabaseResponse, error) {
	res := &SeppoService.CreateEwDatabaseResponse{}
	ewDatabase := s.databaseService.CreateEwDatabase(SeppoDB.NewCreateEwDatabaseFromServiceType(in))
	res.EwDatabase = NewEwDatabaseToServiceType(ewDatabase)
	return res, nil
}

func (s *SeppoServiceServer) EditEwDatabase(ctx context.Context, in *SeppoService.EditEwDatabaseRequest) (*SeppoService.EditEwDatabaseResponse, error) {
	res := &SeppoService.EditEwDatabaseResponse{}
	ewDatabase := s.databaseService.EditEwDatabase(SeppoDB.NewEditEwDatabaseFromServiceType(in))
	res.EwDatabase = NewEwDatabaseToServiceType(ewDatabase)
	return res, nil
}

func (s *SeppoServiceServer) RemoveEwDatabase(ctx context.Context, in *SeppoService.RemoveEwDatabaseRequest) (*SeppoService.RemoveEwDatabaseResponse, error) {
	res := &SeppoService.RemoveEwDatabaseResponse{}
	s.databaseService.RemoveEwDatabase(in.EwDatabaseId)
	return res, nil
}

func (s *SeppoServiceServer) AddVariationToSongDatabase(ctx context.Context, in *SeppoService.AddVariationToSongDatabaseRequest) (*SeppoService.AddVariationToSongDatabaseResponse, error) {
	res := &SeppoService.AddVariationToSongDatabaseResponse{}
	songDatabaseVariation := s.databaseService.AddVariationToSongDatabase(in.SongDatabaseId, in.VariationId)
	res.SongDatabaseVariation = NewSongDatabaseVariationToServiceType(songDatabaseVariation)
	return res, nil
}

func (s *SeppoServiceServer) RemoveVariationFromSongDatabase(ctx context.Context, in *SeppoService.RemoveVariationFromSongDatabaseRequest) (*SeppoService.RemoveVariationFromSongDatabaseResponse, error) {
	res := &SeppoService.RemoveVariationFromSongDatabaseResponse{}
	s.databaseService.RemoveVariationFromSongDatabase(in.SongDatabaseId, in.VariationId)
	return res, nil
}

// InsertEwSongIds Asettaa ewtietokantaan luodun laulun idn
func (s *SeppoServiceServer) InsertEwSongIds(
	ctx context.Context,
	in *SeppoService.InsertEwSongIdsRequest) (*SeppoService.InsertEwSongIdsResponse, error) {
	res := &SeppoService.InsertEwSongIdsResponse{}

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

func (s *SeppoServiceServer) ChangeEwSongIds(ctx context.Context, in *SeppoService.ChangeEwSongIdsRequest) (*SeppoService.ChangeEwSongIdsResponse, error) {
	response := &SeppoService.ChangeEwSongIdsResponse{}

	return response, nil
}

func (s *SeppoServiceServer) SyncEwDatabase(ctx context.Context, in *SeppoService.SyncEwDatabaseRequest) (*SeppoService.SyncEwDatabaseResponse, error) {
	response := &SeppoService.SyncEwDatabaseResponse{}

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
									response.EwSongs = append(response.EwSongs, &SeppoService.EwSong{
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
					response.EwSongs = append(response.EwSongs, &SeppoService.EwSong{
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

func (s SeppoServiceServer) CreateTag(ctx context.Context, in *SeppoService.CreateTagRequest) (*SeppoService.CreateTagResponse, error) {
	res := &SeppoService.CreateTagResponse{}
	tag := s.databaseService.CreateTag(SeppoDB.CreateTagInput{
		Name: in.Name,
	})
	res.Tag = NewTagToServiceType(&tag)
	return res, nil
}

func (s SeppoServiceServer) EditTag(ctx context.Context, in *SeppoService.EditTagRequest) (*SeppoService.EditTagResponse, error) {
	res := &SeppoService.EditTagResponse{}
	tag := s.databaseService.EditTag(SeppoDB.EditTagInput{
		TagID: in.TagId,
		Name:  in.Name,
	})
	res.Tag = NewTagToServiceType(&tag)
	return res, nil
}

func (s SeppoServiceServer) RemoveTag(ctx context.Context, in *SeppoService.RemoveTagRequest) (*SeppoService.RemoveTagResponse, error) {
	res := &SeppoService.RemoveTagResponse{}
	res.Success = s.databaseService.RemoveTag(in.TagId)
	return res, nil
}

func (s SeppoServiceServer) CreateLanguage(ctx context.Context, in *SeppoService.CreateLanguageRequest) (*SeppoService.CreateLanguageResponse, error) {
	res := &SeppoService.CreateLanguageResponse{}
	language := s.databaseService.CreateLanguage(SeppoDB.CreateLanguageInput{
		Name: in.Name,
	})
	res.Language = NewLanguageToServiceType(&language)
	return res, nil
}

func (s SeppoServiceServer) EditLanguage(ctx context.Context, in *SeppoService.EditLanguageRequest) (*SeppoService.EditLanguageResponse, error) {
	res := &SeppoService.EditLanguageResponse{}
	language := s.databaseService.EditLanguage(SeppoDB.EditLanguageInput{
		LanguageID: in.LanguageId,
		Name:       in.Name,
	})
	res.Language = NewLanguageToServiceType(&language)
	return res, nil
}

func (s SeppoServiceServer) RemoveLanguage(ctx context.Context, in *SeppoService.RemoveLanguageRequest) (*SeppoService.RemoveLanguageResponse, error) {
	res := &SeppoService.RemoveLanguageResponse{}
	res.Success = s.databaseService.RemoveLanguage(in.LanguageId)
	return res, nil
}
