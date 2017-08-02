package SeppoDB

type CreateSongInput struct {
	Name          string
	Author        string
	Copyright     string
	Administrator string
	Description   string
	Tags          string
	SongID        uint32
	EwDatabaseId  uint32
}

type CreateVariationInput struct {
	Name string
	Text string
}

type EditVariationInput struct {
	VariationID uint32
	Name        string
	Text        string
	SongID      uint32
}

type UpdateEwSongInput struct {
	Title         string
	Author        string
	Copyright     string
	Administrator string
	Description   string
	Tags          string
	EwDatabaseId  uint32
}

type CreateSongDatabaseInput struct {
	Name string
}

type EditSongDatabaseInput struct {
	SongDatabaseId uint32
	Name           string
}

type CreateEwDatabaseInput struct {
	Name           string
	SongDatabaseId uint32
}

type EditEwDatabaseInput struct {
	EwDatabaseID   uint32
	Name           string
	SongDatabaseID uint32
}
