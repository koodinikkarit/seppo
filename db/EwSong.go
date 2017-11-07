package db

type EwSong struct {
	ID            uint32
	Title         string
	Author        string
	Copyright     string
	Administrator string
	Description   string
	Tags          string
	Text          string

	Variation Variation
}
