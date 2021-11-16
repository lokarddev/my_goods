package environ

import (
	"github.com/joho/godotenv"
	"my_goods/pkg/logger"
	"os"
)

var (
	Port    string
	DbHost  string
	DbPort  string
	DbUser  string
	DbPass  string
	DbName  string
	SslMode string

	Salt string
)

func Env() {
	err := godotenv.Load()
	if err != nil {
		logger.Info("Error loading .env file, continue with default environment vars.")
	}
	_, err = godotenv.Read()
	if err != nil {
		logger.Info(".env file cannot be read, continue with default environment vars.")
	}
	Port = os.Getenv("PORT")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPass = os.Getenv("DB_PASS")
	DbName = os.Getenv("DB_NAME")
	SslMode = os.Getenv("SSL_MODE")
	Salt = os.Getenv("SALT")
}
