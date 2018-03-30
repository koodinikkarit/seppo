package matias

import (
	"github.com/koodinikkarit/seppo/models"
)

type IsClientAcceptedEvent struct {
	MatiasClientID uint32
	State          bool
}

type CreatedEwDatabaseEvent struct {
	EwDatabase models.EwDatabase
}

type UpdatedEwDatabaseEvent struct {
	EwDatabase models.EwDatabase
}

type RemovedEwDatabaseEvent struct {
	EwDatabaseId uint32
}

type CreatedVariationEvent struct {
	Variation models.Variation
}

type UpdatedVariationEvent struct {
	Variation models.Variation
}

type RemovedVariationEvent struct {
	VariationId uint32
}

type CreatedSongDatabaseEvent struct {
	SongDatabase models.SongDatabase
}

type UpdateSongDatabaseEvent struct {
	SongDatabase models.SongDatabase
}

type RemovedSongDatabaseEvent struct {
	SongDatabaseID uint32
}

type CreatedSongDatabaseVariationEvent struct {
	SongDatabaseID uint32
	VariationID    uint32
}

type RemovedSongDatabaseVariationEvent struct {
	SongDatabaseID uint32
	VariationID    uint32
}
