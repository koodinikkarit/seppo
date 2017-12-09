package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/services"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	DBUser   string `yaml:"dbUser"`
	DBPasswd string `yaml:"dbPasswd"`
	DBIP     string `yaml:"dbIp"`
	DBPort   string `yaml:"dbPort"`
	DBName   string `yaml:"dbName"`
	Port     string `yaml:"port"`
}

func readConfig() *Config {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := Config{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &c
}

func main() {
	var dbUser string
	var dbPassword string
	var dbIP string
	var dbPort string
	var dbName string
	var port string

	if os.Getenv("SEPPO_USE_CONFIG_FILE") == "true" {
		c := readConfig()
		fmt.Println("use configfile ", c)
		dbUser = c.DBUser
		if dbUser == "" {
			panic("Configfile dbUser is empty")
		}
		dbPassword = c.DBPasswd
		if dbPassword == "" {
			panic("Configfile dbPasswd is empty")
		}
		dbIP = c.DBIP
		if dbIP == "" {
			panic("Configfile dbIp is empty")
		}
		dbPort = c.DBPort
		if dbPort == "" {
			panic("Configfile dbPort is empty")
		}
		dbName = c.DBName
		if dbName == "" {
			panic("Configfile dbName is empty")
		}
		port = c.Port
		if port == "" {
			panic("Configfile port is empty")
		}
	} else {
		dbUser = os.Getenv("SEPPO_DB_USERNAME")
		if dbUser == "" {
			panic("No enviroment variable dbUser")
		}
		dbPassword = os.Getenv("SEPPO_DB_PASSWORD")
		if dbPassword == "" {
			panic("No enviroment variable SEPPO_DB_PASSWORD")
		}
		dbIP = os.Getenv("SEPPO_DB_IP")
		if dbIP == "" {
			panic("No enviroment variable SEPPO_DB_IP")
		}
		dbPort = os.Getenv("SEPPO_DB_PORT")
		if dbPort == "" {
			panic("No enviroment variable SEPPO_DB_PORT")
		}
		dbName = os.Getenv("SEPPO_DB_NAME")
		if dbName == "" {
			panic("No enviroment variable  SEPPO_DB_NAME")
		}
		port = os.Getenv("SEPPO_PORT")
		if port == "" {
			panic("No enviroment variable SEPPO_PORT")
		}
	}

	fmt.Println("Creating mysql connection with parameters")
	fmt.Println(dbUser)
	fmt.Println(dbPassword)
	fmt.Println(dbIP)
	fmt.Println(dbPort)
	fmt.Println(dbName)
	fmt.Println(port)

	db.Migrate(
		dbUser,
		dbPassword,
		dbIP,
		dbPort,
		dbName,
		1,
	)

	getGormDB := db.CreateGetGormDB(
		dbUser,
		dbPassword,
		dbIP,
		dbPort,
		dbName,
	)

	services.StartSeppoService(
		port,
		getGormDB,
	)
}
