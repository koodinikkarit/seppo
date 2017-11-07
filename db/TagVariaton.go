package db

import "time"

type TagVariation struct {
	ID          uint32
	TagID       uint32
	VariationID uint32
	CreatedAt   *time.Time
}
