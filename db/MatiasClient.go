package db

import "time"

type MatiasClient struct {
	id        uint32
	ClientKey string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
