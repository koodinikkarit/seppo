package models

import (
	"time"
)

type Language struct {
	ID   uint32
	Name string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
