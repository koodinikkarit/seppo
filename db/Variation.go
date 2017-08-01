package SeppoDB

type Variation struct {
	ID      uint32
	Name    string
	Text    string `gorm:"size:2048"`
	SongID  uint32
	Version uint64
}
