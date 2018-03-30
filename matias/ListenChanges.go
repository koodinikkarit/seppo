package matias

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/go-clientlibs/matias"
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

func newVariationEvent(
	variation models.Variation,
	newestVariationVersion models.VariationVersion,
) *MatiasService.EventItem {
	newMatiasVariation := &MatiasService.Variation{
		Id:      variation.ID,
		Name:    newestVariationVersion.Name,
		Text:    newestVariationVersion.Text,
		Version: newestVariationVersion.Version,
	}

	if variation.AuthorID != nil {
		newMatiasVariation.AuthorId = *variation.AuthorID
	}

	if variation.CopyrightID != nil {
		newMatiasVariation.CopyrightId = *variation.CopyrightID
	}

	if variation.LanguageID != nil {
		newMatiasVariation.LanguageId = *variation.LanguageID
	}

	if variation.SongID != nil {
		newMatiasVariation.SongId = *variation.SongID
	}

	return &MatiasService.EventItem{
		EventMessage: &MatiasService.EventItem_NewVariation{
			NewVariation: newMatiasVariation,
		},
	}
}

func newSongDatabaseVariation(variationID uint32, songDatabaseID uint32) *MatiasService.EventItem {
	return &MatiasService.EventItem{
		EventMessage: &MatiasService.EventItem_NewSongDatabaseVariation{
			NewSongDatabaseVariation: &MatiasService.SongDatabaseVariation{
				SongDatabaseId: songDatabaseID,
				VariationId:    variationID,
			},
		},
	}
}

func newRemovedSongDatabaseVariation(variationID uint32, songDatabaseID uint32) *MatiasService.EventItem {
	return &MatiasService.EventItem{
		EventMessage: &MatiasService.EventItem_RemovedSongDatabaseVariation{
			RemovedSongDatabaseVariation: &MatiasService.SongDatabaseVariation{
				SongDatabaseId: songDatabaseID,
				VariationId:    variationID,
			},
		},
	}
}

func newSongDatabaseTag(songDatabaseTag models.SongDatabaseTag) *MatiasService.EventItem {
	return &MatiasService.EventItem{
		EventMessage: &MatiasService.EventItem_NewSongDatabaseTag{
			NewSongDatabaseTag: &MatiasService.SongDatabaseTag{
				Id:             songDatabaseTag.ID,
				SongDatabaseId: songDatabaseTag.SongDatabaseID,
				TagId:          songDatabaseTag.TagID,
			},
		},
	}
}

func newTagVariation(tagVariation models.TagVariation) *MatiasService.EventItem {
	return &MatiasService.EventItem{
		EventMessage: &MatiasService.EventItem_NewTagVariation{
			NewTagVariation: &MatiasService.TagVariation{
				Id:          tagVariation.ID,
				TagId:       tagVariation.TagID,
				VariationId: tagVariation.VariationID,
			},
		},
	}
}

func (ms *MatiasServiceServer) sendClientAccepted(
	db *gorm.DB,
	stream MatiasService.Matias_ListenChangesServer,
	matiasClient models.MatiasClient,
) {
	stream.Send(newClientAcceptedEvent(true))

	var ewDatabases []models.EwDatabase
	db.Where("matias_client_id = ?", matiasClient.ID).
		Find(&ewDatabases)

	var songDatabaseIds []uint32
	for _, ewDatabase := range ewDatabases {
		stream.Send(newEwDatabaseEvent(ewDatabase))
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
		stream.Send(newSongDatabaseTag(songDatabaseTag))
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
		stream.Send(newSongDatabaseVariation(
			songDatabaseVariation.VariationID,
			songDatabaseVariation.SongDatabaseID))
		variationIds = append(
			variationIds,
			songDatabaseVariation.VariationID,
		)
	}
	for _, tagVariation := range tagVariations {
		stream.Send(newTagVariation(tagVariation))
		variationIds = append(
			variationIds,
			tagVariation.VariationID,
		)
	}

	var variations []models.Variation
	db.Where("id in (?)", variationIds).
		Find(&variations)

	var variationVersions []models.VariationVersion
	err := db.Table("variation_versions").
		Joins("left join variation_versions vv2 on (variation_versions.variation_id = vv2.variation_id and variation_versions.version < vv2.version)").
		Where("vv2.id is null").
		Where("variation_versions.variation_id in (?)", variationIds).
		Find(&variationVersions).Error

	if err != nil {
		log.Fatalln("error: ", err)
	}

	for _, variation := range variations {
		var newestVersion models.VariationVersion
		for _, variationVersion := range variationVersions {
			if variationVersion.VariationID == variation.ID {
				newestVersion = variationVersion
			}
		}
		if newestVersion.ID > 0 {
			stream.Send(newVariationEvent(variation, newestVersion))
		}
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
			Hostname:  req.HostName,
		}
		db.Create(&matiasClient)
	}

	eventChannel := make(chan interface{})

	ms.pubSub.AddSub(
		eventChannel,
		"newVariation",
		"newEwDatabase",
		"clientAccepted",
		"createdVariation",
		"createdSongDatabaseVariation",
		"removedSongDatabaseVariation",
	)

	if matiasClient.Accepted == true {
		ms.sendClientAccepted(db, stream, matiasClient)
	}

	for {
		fmt.Println("kuunnellaan  event")
		eventType := <-eventChannel
		fmt.Println("newEvent", eventType)
		switch event := eventType.(type) {
		case models.EwDatabase:
			fmt.Println("ewDatabases", eventType)
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
			fmt.Println("accpeted", event.MatiasClientID, matiasClient.ID)
			if event.MatiasClientID != matiasClient.ID {
				continue
			}
			fmt.Println("accpeted", event.MatiasClientID)
			matiasClient.Accepted = event.State

			if event.State == true {
				ms.sendClientAccepted(db, stream, matiasClient)
			} else {
				stream.Send(newClientAcceptedEvent(false))
			}
		case CreatedVariationEvent:
			var newestVariationVersion models.VariationVersion
			db.Where("variation_id = ?", event.Variation.ID).Order("version").
				Find(&newestVariationVersion)
			if newestVariationVersion.ID > 0 {
				stream.Send(newVariationEvent(event.Variation, newestVariationVersion))
			}
		case CreatedSongDatabaseVariationEvent:
			stream.Send(newSongDatabaseVariation(event.VariationID, event.SongDatabaseID))
		case RemovedSongDatabaseVariationEvent:
			stream.Send(newRemovedSongDatabaseVariation(event.VariationID, event.SongDatabaseID))
		}
	}

	return nil
}
