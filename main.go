package main

import (
	"log"
	"os"
	"sync"

	"github.com/cskr/pubsub"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/koodinikkarit/seppo/database"
	"github.com/koodinikkarit/seppo/matias"
	"github.com/koodinikkarit/seppo/service"
)

func main() {
	var dbUser string
	var dbPassword string
	var dbIP string
	var dbPort string
	var dbName string
	var seppoPort string
	var matiasPort string

	dbUser = os.Getenv("DB_USERNAME")
	if dbUser == "" {
		log.Fatalln("No enviroment variable DB_USERNAME")
	}
	dbPassword = os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatalln("No enviroment variable DB_PASSWORD")
	}
	dbIP = os.Getenv("DB_IP")
	if dbIP == "" {
		log.Fatalln("No enviroment variable DB_IP")
	}
	dbPort = os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Println("No enviroment variable DB_PORT using 3306")
		dbPort = "3306"
	}
	dbName = os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatalln("No enviroment variable  DB_NAME")
	}
	seppoPort = os.Getenv("SEPPO_PORT")
	if seppoPort == "" {
		log.Println("No enviroment variable SEPPO_PORT using 3214")
		seppoPort = "3214"
	}
	matiasPort = os.Getenv("MATIAS_PORT")
	if matiasPort == "" {
		log.Println("No enviroment variable MATIAS_PORT using 6755")
		matiasPort = "6755"
	}
	log.Println("Creating mysql connection with parameters")
	log.Printf("dbUser: %v\n", dbUser)
	log.Printf("dbPassword: %v\n", dbPassword)
	log.Printf("dbIP: %v\n", dbIP)
	log.Printf("dbPort: %v\n", dbPort)
	log.Printf("dbName: %v\n", dbName)
	log.Printf("seppoPort: %v\n", seppoPort)
	log.Printf("matiasPort: %v\n", matiasPort)

	getDB := database.CreateGetDB(
		dbUser,
		dbPassword,
		dbIP,
		dbPort,
		dbName,
		1,
	)

	pubSub := pubsub.New(1)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go service.StartSeppoService(
		getDB,
		pubSub,
		seppoPort,
	)
	wg.Add(1)
	go matias.StartMatiasService(
		getDB,
		pubSub,
		matiasPort,
	)
	wg.Wait()
}
