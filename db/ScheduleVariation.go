package db

import "time"

type ScheduleVariation struct {
	ID                 uint32
	ScheduleID         uint32
	VariationVersionID uint32
	OrderNumber        uint32
	CreatedAt          *time.Time
}
