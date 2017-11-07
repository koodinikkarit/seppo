package db

import "time"

type VariationVersion struct {
	ID          uint32
	VariationID uint32
	Name        string
	Text        string
	Version     uint32
	Newest      bool
	CreatedAt   *time.Time
	DisabledAt  *time.Time
}
