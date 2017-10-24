package logs

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
)

func InsertLog(
	newDb *gorm.DB,
	logType uint32,
	message string,
) {
	log := db.Log{
		LogType:     logType,
		Message:     message,
		MessageDate: time.Now(),
	}
	newDb.Create(&log)
}
