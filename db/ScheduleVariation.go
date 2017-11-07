package db

import "time"

type ScheduleVariation struct {
	ID          uint32
	ScheduleID  uint32
	VariationID uint32
	OrderNumber uint32
	CreatedAt   *time.Time
}
