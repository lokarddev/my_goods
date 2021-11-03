package environ

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"my_goods/pkg/logger"
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
		viper.AddConfigPath("/configs")
		viper.SetConfigName("local")
		viper.SetConfigType("yml")
		viper.AddConfigPath("/app/configs/")
		err = viper.ReadInConfig()
		if err != nil {
			logger.Error(err)
		}
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
	Host = os.Getenv("HOST")
	Port = os.Getenv("PORT")

	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPass = os.Getenv("DB_PASS")
	DbName = os.Getenv("DB_NAME")
	SslMode = os.Getenv("SSL_MODE")
}

func testEnv() {
	Host = viper.GetString("Host")
	Port = viper.GetString("Port")

	DbHost = viper.GetString("db.Host")
	DbPort = viper.GetString("db.Port")
	DbUser = viper.GetString("db.User")
	DbPass = viper.GetString("db.Pass")
	DbName = viper.GetString("db.Name")
	SslMode = viper.GetString("db.SslMode")
}
