package seppo

import (
	"github.com/koodinikkarit/seppo/db"
)

func CreateSeppoServer(config Config) {
	dbUser := "root"
	dbPasswd := "qwerty"
	dbIP := "seppo-mysql"
	dbPort := "3306"
	dbName := "seppo"
	seppoPort := "3000"

	if config.DBUser != "" {
		dbUser = config.DBUser
	}
	if config.DBPasswd != "" {
		dbPasswd = config.DBPasswd
	}
	if config.DBIP != "" {
		dbIP = config.DBIP
	}
	if config.DBPort != "" {
		dbPort = config.DBPort
	}
	if config.DBName != "" {
		dbName = config.DBName
	}
	if config.SeppoPort != "" {
		seppoPort = config.SeppoPort
	}

	databaseService := SeppoDB.CreateDb(
		dbUser,
		dbPasswd,
		dbIP,
		dbPort,
		dbName)
	go databaseService.Start()
	CreateSeppoService(seppoPort, databaseService)
}
