package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func CreateGetDB(
	dbUser string,
	dbPass string,
	dbIP string,
	dbPort string,
	dbName string,
) func() *sql.DB {
	return func() *sql.DB {
		db, err := sql.Open(
			"mysql",
			dbUser+":"+dbPass+"@tcp("+dbIP+":"+dbPort+")/"+dbName+"?parseTime=True&loc=Local",
		)
		if err != nil {
			panic(err)
		}
		return db
	}
}
