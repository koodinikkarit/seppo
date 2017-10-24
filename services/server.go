package seppo

import (
	"sync"

	"github.com/koodinikkarit/seppo/db"
)

func CreateSeppoServer(config Config) {
	databaseService := SeppoDB.NewDatabaseService(
		config.DBUser,
		config.DBPasswd,
		config.DBIP,
		config.DBPort,
		config.DBName)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		go databaseService.Start()
	}()
	wg.Add(1)
	go func() {
		go CreateSeppoService(config.SeppoPort, databaseService)
	}()
	wg.Add(1)
	go func() {
		go NewMatiasService(config.MatiasPort, databaseService)
	}()
	wg.Wait()
}
