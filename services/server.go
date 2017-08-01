package seppo

import (
	"github.com/koodinikkarit/seppo/db"
)

func CreateSeppoServer(config Config) {

	databaseService := SeppoDB.CreateDb(
		config.DBUser,
		config.DBPasswd,
		config.DBIP,
		config.DBPort,
		config.DBName)
	go databaseService.Start()
	CreateSeppoService("3214", databaseService)
}
