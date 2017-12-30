package database
import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
)

func Migrate(
	dbUser string,
	dbPass string,
	dbIP string,
	dbPort string,
	dbName string,
	steps int,
) {
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbIP+":"+dbPort+")/")
	if err != nil {
		log.Fatalf("Failed to open db connection for db creation %v", err)
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatalf("Failed to create database %v error %v", dbName, err)
	}
	db.Close()

	db, err = sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbIP+":"+dbPort+")/"+dbName+"?multiStatements=true")
	if err != nil {
		log.Fatalf("Error while opening connnection for migrations %v", err)
	}
	defer db.Close()

	driver, _ := mysql.WithInstance(db, &mysql.Config{})

	m, migErr := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if migErr != nil {
		log.Fatalf("Migration error %v", migErr)
	}
	m.Steps(steps)
}
