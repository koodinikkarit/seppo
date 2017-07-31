package seppo

import (
	"io/ioutil"
	"log"

	"github.com/koodinikkarit/seppo/db"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DBUser   string `yaml:"dbUser"`
	DBPasswd string `yaml:"dbPasswd"`
	DBIP     string `yaml:"dbIp"`
	DBPort   string `yaml:"dbPort"`
	DBName   string `yaml:"dbName"`
}

func (c *Config) getConf() *Config {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func CreateSeppoServer() {
	config := &Config{}
	config.getConf()

	databaseService := SeppoDB.CreateDb(
		config.DBUser,
		config.DBPasswd,
		config.DBIP,
		config.DBPort,
		config.DBName)
	go databaseService.Start()
	CreateSeppoService("3214", databaseService)
}
