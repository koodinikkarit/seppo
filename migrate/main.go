package main

import (
	"os"

	"github.com/koodinikkarit/seppo/db"
)

func main() {
	dbUser := os.Getenv("SEPPO_DB_USERNAME")
	dbPassword := os.Getenv("SEPPO_DB_PASSWORD")
	dbIP := os.Getenv("SEPPO_DB_IP")
	dbPort := os.Getenv("SEPPO_DB_PORT")
	dbName := os.Getenv("SEPPO_DB_NAME")

	db.Migrate(
		dbUser,
		dbPassword,
		dbIP,
		dbPort,
		dbName,
		1,
	)
}
