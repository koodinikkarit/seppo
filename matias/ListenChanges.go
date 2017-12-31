package matias

import (
	"github.com/koodinikkarit/seppo/matias_service"
	"github.com/koodinikkarit/seppo/models"
)

func newClientAcceptedEvent(state bool) *MatiasService.EventItem {
	return &MatiasService.EventItem{
		EventMessage: &MatiasService.EventItem_AcceptedKey{
			AcceptedKey: state,
		},
	}
}

func newEwDatabaseEvent(ewDatabase models.EwDatabase) *MatiasService.EventItem {
	return &MatiasService.EventItem{
		EventMessage: &MatiasService.EventItem_NewEwDatabase{
			NewEwDatabase: &MatiasService.EwDatabase{
				Id:             ewDatabase.ID,
				FilesystemPath: ewDatabase.FilesystemPath,
				SongDatabaseId: ewDatabase.SongDatabaseID,
			},
		},
	}
}

func (ms *MatiasServiceServer) ListenChanges(
	req *MatiasService.ListenEventsRequest,
	stream MatiasService.Matias_ListenChangesServer,
) error {
	db := ms.getDB()

	var matiasClient models.MatiasClient
	db.Where("client_key = ?", req.Key).
		First(&matiasClient)

	if matiasClient.ID == 0 {
		matiasClient := models.MatiasClient{
			ClientKey: req.Key,
		}
		db.Create(&matiasClient)
	}

	eventChannel := make(chan interface{})

	ms.pubSub.AddSub(
		eventChannel,
		"newVariation",
		"newEwDatabase",
		"clientAccepted",
	)

	if matiasClient.Accepted == true {
		stream.Send(newClientAcceptedEvent(true))

		var ewDatabases []models.EwDatabase
		db.Where("matias_client_id = ?", matiasClient.ID).
			Find(&ewDatabases)

		var songDatabaseIds []uint32
		for _, ewDatabase := range ewDatabases {
			songDatabaseIds = append(
				songDatabaseIds,
				ewDatabase.SongDatabaseID,
			)
		}

		var songDatabaseVariations []models.SongDatabaseVariation
		db.Where("song_database_id in (?)", songDatabaseIds).
			Find(&songDatabaseVariations)

		var songDatabaseTags []models.SongDatabaseTag
		db.Where("song_database_id in (?)", songDatabaseTags).
			Find(&songDatabaseTags)

		var tagIds []uint32
		for _, songDatabaseTag := range songDatabaseTags {
			tagIds = append(
				tagIds,
				songDatabaseTag.TagID,
			)
		}

		var tagVariations []models.TagVariation
		db.Where("tag_id in (?)", tagIds).
			Find(&tagVariations)

		var variationIds []uint32
		for _, songDatabaseVariation := range songDatabaseVariations {
			variationIds = append(
				variationIds,
				songDatabaseVariation.VariationID,
			)
		}
		for _, tagVariation := range tagVariations {
			variationIds = append(
				variationIds,
				tagVariation.VariationID,
			)
		}

		var variations []models.Variation
		db.Where("id in (?)", variationIds).
			Find(&variations)
	}

	for {
		eventType := <-eventChannel
		switch event := eventType.(type) {
		case models.EwDatabase:
			if matiasClient.ID != event.MatiasClientID {
				continue
			}
			stream.Send(newEwDatabaseEvent(event))
		case models.SongDatabaseVariation:
		case models.SongDatabaseTag:
		case models.TagVariation:
		case models.Variation:
		case models.Song:
		case IsClientAcceptedEvent:
			if event.MatiasClientID != matiasClient.ID {
				continue
			}
			matiasClient.Accepted = event.State
			stream.Send(newClientAcceptedEvent(event.State))
		}
	}

	return nil
}
