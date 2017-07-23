package seppo

import (
	"golang.org/x/net/context"

	"fmt"

	seppo "github.com/koodinikkarit/seppo/db"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
)

func (s *SeppoServiceServer) SyncEwDatabase(ctx context.Context, in *SeppoService.SyncEwDatabaseRequest) (*SeppoService.SyncEwDatabaseResponse, error) {
	response := &SeppoService.SyncEwDatabaseResponse{}

	var songIds []uint32
	var ewDatabaseLinks []seppo.EwDatabaseLink

	for _, ewSong := range in.EwSongs {
		songIds = append(songIds, ewSong.Id)
	}

	s.databaseService.GetDb().Debug().Where("ew_database_song_id in (?)", songIds).Where("ew_database_id = ?", in.EwDatabaseId).Find(&ewDatabaseLinks)

	for _, ewSong := range in.EwSongs {
		var found bool
		for _, ewDatabaseLink := range ewDatabaseLinks {
			if ewDatabaseLink.EwDatabaseSongID == ewSong.Id {
				found = true
			}
		}

		if found == false {
			s.databaseService.CreateSongChannel <- seppo.CreateSongInput{
				Name:         ewSong.Title,
				SongID:       ewSong.Id,
				EwDatabaseId: in.EwDatabaseId,
			}
		}
	}

	var variationIdt []uint32

	for _, ewDatabaseLink := range ewDatabaseLinks {
		var found bool
		for _, ewSong := range in.EwSongs {
			if ewDatabaseLink.EwDatabaseSongID == ewSong.Id {
				found = true
			}
		}

		if found == false {
			variationIdt = append(variationIdt, ewDatabaseLink.VariationID)
		}
	}

	variations := []seppo.Variation{}

	s.databaseService.GetDb().Where("id in (?)", variationIdt).Find(&variations)

	for _, variation := range variations {
		response.EwSongs = append(response.EwSongs, &SeppoService.EwSong{
			Id:    variation.ID,
			Title: variation.Name,
		})
	}

	fmt.Println("SyncDatabase")

	fmt.Println("channel ", s.createSongChannel)

	return response, nil
}
