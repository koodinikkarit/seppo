package SeppoDB

type Tag struct {
	ID   uint32
	Name string

	Variations    []Variation
	SongDatabases []SongDatabase
}
