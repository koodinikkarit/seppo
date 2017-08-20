package SeppoDB

type VariationText struct {
	ID          uint32
	VariationID uint32
	Text        string `gorm:"size:4096"`
}
