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
	Title string
	Author        string
	Copyright     string
	Administrator string
	Description   string
	Tags          string
	EwDatabaseId  uint32
}
