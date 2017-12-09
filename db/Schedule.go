package db

import "time"

type Schedule struct {
	ID        uint32
	Name      string
	Start     *time.Time
	End       *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
