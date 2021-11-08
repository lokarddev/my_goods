package environ

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
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
	Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	switch Debug {
	case true:
		testEnv()
	case false:
		envFile()
	}
}

func testEnv() {
	viper.SetConfigFile("./app/configs/local.yml")
	err := viper.ReadInConfig()
	if err != nil {
		viper.SetConfigFile("./configs/local.yml")
		err = viper.ReadInConfig()
		if err != nil {
			logger.Error(err)
		}
	}
	Port = viper.GetString("port")
	DbHost = viper.GetString("db.host")
	DbPort = viper.GetString("db.port")
	DbUser = viper.GetString("db.user")
	DbPass = viper.GetString("db.pwd")
	DbName = viper.GetString("db.name")
	SslMode = viper.GetString("db.sslMode")
}

func envFile() {
	err := godotenv.Load()
	if err != nil {
		logger.Info("Error loading .env file, continue with test environment vars.")
		testEnv()
		return
	}
	_, err = godotenv.Read()
	if err != nil {
		logger.Info(".env file cannot be read, continue without it.")
		testEnv()
		return
	}
	Port = os.Getenv("PORT")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPass = os.Getenv("DB_PASS")
	DbName = os.Getenv("DB_NAME")
	SslMode = os.Getenv("SSL_MODE")
}
