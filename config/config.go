package config

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// MysqlUsername is username of mysql server
var MysqlUsername = os.Getenv("MYSQL_USERNAME")
var MysqlPassword = os.Getenv("MYSQL_PASSWORD")
var MysqlHost = os.Getenv("MYSQL_HOST")
var MysqlPort = os.Getenv("MYSQL_PORT")
var MysqlDatabase = os.Getenv("MYSQL_DATABASE")
var SeppoPort = os.Getenv("SEPPO_PORT")
var MatiasPort = os.Getenv("MATIAS_PORT")

// ValidateConfig Validated application config environment variables
func ValidateConfig() {
	if MysqlUsername == "" {
		log.Fatalln("No environment variable MYSQL_USERNAME")
	}

	if MysqlHost == "" {
		log.Fatalln("No environment variable MYSQL_HOST")
	}

	if MysqlPort == "" {
		MysqlPort = "3306"
	}

	if MysqlDatabase == "" {
		log.Fatalln("No environment variable MYSQL_DATABASE")
	}

	if SeppoPort == "" {
		log.Fatalln("No environment variable SEPPO_PORT")
	}

	if MatiasPort == "" {
		log.Fatalln("No environment variable MATIAS_PORT")
	}
}
