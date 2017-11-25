package main

import (
	"os"
	"sync"

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
	matiasPort := os.Getenv("SEPPO_MATIAS_PORT")

	db.Migrate(
		dbUser,
		dbPassword,
		dbIP,
		dbPort,
		dbName,
		1,
	)

	getDb := db.CreateGetDB(
		dbUser,
		dbPassword,
		dbIP,
		dbPort,
		dbName,
	)

	getGormDB := db.CreateGetGormDB(
		dbUser,
		dbPassword,
		dbIP,
		dbPort,
		dbName,
	)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		go services.StartSeppoService(
			seppoPort,
			getDb,
			getGormDB,
		)
	}()
	wg.Add(1)
	go func() {
		go services.StartMatiasService(
			matiasPort,
			getDb,
			getGormDB,
		)
	}()
	wg.Wait()
}
