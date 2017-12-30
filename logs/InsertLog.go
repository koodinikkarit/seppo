package logs

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/models"
)

func InsertLog(
	db *gorm.DB,
	logType uint32,
	message string,
) {
	log := models.Log{
		LogType:     logType,
		Message:     message,
		MessageDate: time.Now(),
	}
	db.Create(&log)
}
