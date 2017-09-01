package SeppoDB

type Variation struct {
	ID         uint32
	Name       string
	SongID     uint32
	Version    uint64
	LanguageID uint32

	Song     Song
	Language Language
}
