package seppo

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
