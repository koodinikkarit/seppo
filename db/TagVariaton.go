package db

import "time"

type TagVariation struct {
	ID                 uint32
	TagID              uint32
	VariationVersionID uint32
	CreatedAt          *time.Time
}
