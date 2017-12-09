package db

import "github.com/jinzhu/gorm"

func CreateGetGormDB(
	dbUser string,
	dbPass string,
	dbIP string,
	dbPort string,
	dbName string,
) func() *gorm.DB {
	return func() *gorm.DB {
		db, err := gorm.Open(
			"mysql",
			dbUser+":"+dbPass+"@tcp("+dbIP+":"+dbPort+")/"+dbName+"?parseTime=True&loc=Local",
		)
		if err != nil {
			panic(err)
		}
		return db
	}
}
