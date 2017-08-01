package seppo

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	DBUser   string `yaml:"dbUser"`
	DBPasswd string `yaml:"dbPasswd"`
	DBIP     string `yaml:"dbIp"`
	DBPort   string `yaml:"dbPort"`
	DBName   string `yaml:"dbName"`
}

func (c *Config) GetConf() *Config {

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

func LoadConfig(path string) Config {
	c := Config{}

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
