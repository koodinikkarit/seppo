package db

import "time"

type Branch struct {
	ID                            uint32
	SourceVariationVersionID      uint32
	DestinationVariationVersionID uint32
	CreatedAt                     time.Time
}
