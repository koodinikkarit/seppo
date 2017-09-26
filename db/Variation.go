package SeppoDB

type Variation struct {
	ID         uint32
	Name       string
	SongID     uint32
	Version    uint64
	LanguageID uint32
	Origin     string
	DeletedAt  string

	Song     Song
	Language Language
}
