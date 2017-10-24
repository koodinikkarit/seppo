package SeppoDB

type EwDatabase struct {
	ID             uint32
	Name           string
	SongDatabaseID uint32
	Key string


	SongDatabase SongDatabase
}
