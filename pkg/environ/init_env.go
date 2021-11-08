package environ

import (
	"github.com/joho/godotenv"
	"my_goods/pkg/logger"
	"os"
	"strconv"
)

var (
	Port    string
	Debug   bool
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
		logger.Info("Error loading .env file, continue with default environment vars.")
	}
	_, err = godotenv.Read()
	if err != nil {
		logger.Info(".env file cannot be read, continue with default environment vars.")
	}
	Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	Port = os.Getenv("PORT")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPass = os.Getenv("DB_PASS")
	DbName = os.Getenv("DB_NAME")
	SslMode = os.Getenv("SSL_MODE")
}
