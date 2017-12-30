package models

import "time"

type Merge struct {
	ID                            uint32
	VariationVersion1ID           uint32
	VariationVersion2ID           uint32
	DestinationVariationVersionID uint32
	CreatedAt                     time.Time
}
