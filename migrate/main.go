package main

import (
	"os"

	"github.com/koodinikkarit/seppo/database"
)

func main() {
	dbUser := os.Getenv("SEPPO_DB_USERNAME")
	dbPassword := os.Getenv("SEPPO_DB_PASSWORD")
	dbIP := os.Getenv("SEPPO_DB_IP")
	dbPort := os.Getenv("SEPPO_DB_PORT")
	dbName := os.Getenv("SEPPO_DB_NAME")

	database.Migrate(
		dbUser,
		dbPassword,
		dbIP,
		dbPort,
		dbName,
		1,
	)
}
