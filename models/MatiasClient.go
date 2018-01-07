package models

import "time"

type MatiasClient struct {
	ID        uint32
	Name      string
	Hostname  string
	ClientKey string
	Accepted  bool
	Connected bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
