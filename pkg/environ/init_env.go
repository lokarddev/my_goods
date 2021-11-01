package environ

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	Host string
	Port string

	DbHost  string
	DbPort  string
	DbUser  string
	DbPass  string
	DbName  string
	SslMode string
)

func Env() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warning("Error loading .env file, continue without it.")
	}
	_, err = godotenv.Read()
	if err != nil {
		logrus.Warning(".env file cannot be read, continue without it.")
	}
	Host = os.Getenv("HOST")
	Port = os.Getenv("PORT")

	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPass = os.Getenv("DB_PASS")
	DbName = os.Getenv("DB_NAME")
	SslMode = os.Getenv("SSL_MODE")
}
