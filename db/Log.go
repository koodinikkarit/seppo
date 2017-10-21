package db

import "time"

type Log struct {
	ID          uint32
	LogType     uint32
	Message     string
	MessageDate time.Time
}
