package db

import (
	"database/sql"

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
		panic(err)
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic(err)
	}
	db.Close()

	db, err = sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbIP+":"+dbPort+")/"+dbName+"?multiStatements=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	driver, _ := mysql.WithInstance(db, &mysql.Config{})

	m, migErr := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if migErr != nil {
		panic(migErr)
	}
	m.Steps(steps)
}
