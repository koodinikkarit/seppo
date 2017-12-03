package main

import (
	"os"

	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/services"
)

func main() {
	dbUser := os.Getenv("SEPPO_DB_USERNAME")
	dbPassword := os.Getenv("SEPPO_DB_PASSWORD")
	dbIP := os.Getenv("SEPPO_DB_IP")
	dbPort := os.Getenv("SEPPO_DB_PORT")
	dbName := os.Getenv("SEPPO_DB_NAME")
	seppoPort := os.Getenv("SEPPO_PORT")

	db.Migrate(
		dbUser,
		dbPassword,
		dbIP,
		dbPort,
		dbName,
		1,
	)

	getGormDB := db.CreateGetGormDB(
		dbUser,
		dbPassword,
		dbIP,
		dbPort,
		dbName,
	)

	services.StartSeppoService(
		seppoPort,
		getGormDB,
	)
}
