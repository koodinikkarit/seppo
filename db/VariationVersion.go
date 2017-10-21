package db

import "time"

type VariationVersion struct {
	ID          uint32
	VariationID uint32
	Name        string
	Text        string
	Version     uint32
	CreatedAt   *time.Time
	DisabledAt  *time.Time
}
