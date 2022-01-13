package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	Production, _  = strconv.ParseBool(os.Getenv("PRODUCTION"))
	Automigrations bool

	DbHost   string
	DbPort   string
	DbUser   string
	DbPass   string
	DbName   string
	DbSsl    string
	DbSchema string
)

func InitEnvVariables() error {
	if err := godotenv.Load(".env"); err != nil {
		if !Production {
			log.Fatalf("Error loading .env file, continue without it.")
		}
	}
	_, err := godotenv.Read(".env")
	if err != nil {
		if !Production {
			log.Fatalf(".env file cannot be read, continue without it.")
		}
	}
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPass = os.Getenv("DB_PASS")
	DbName = os.Getenv("DB_NAME")
	DbSsl = os.Getenv("SSL_MODE")
	DbSchema = os.Getenv("DB_SCHEMA")

	Automigrations, _ = strconv.ParseBool(os.Getenv("AUTOMIGRATIONS"))
	return err
}
