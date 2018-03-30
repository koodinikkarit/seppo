package main

import (
	"sync"

	"github.com/cskr/pubsub"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/koodinikkarit/database"
	"github.com/koodinikkarit/seppo/config"
	"github.com/koodinikkarit/seppo/matias"
	"github.com/koodinikkarit/seppo/service"
)

func main() {
	config.ValidateConfig()

	getDB := database.CreateGetDB(
		config.MysqlUsername,
		config.MysqlPassword,
		config.MysqlHost,
		config.MysqlPort,
		config.MysqlDatabase,
	)

	pubSub := pubsub.New(1)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go service.StartSeppoService(
		getDB,
		pubSub,
		config.SeppoPort,
	)
	wg.Add(1)
	go matias.StartMatiasService(
		getDB,
		pubSub,
		config.MatiasPort,
	)
	wg.Wait()
}
