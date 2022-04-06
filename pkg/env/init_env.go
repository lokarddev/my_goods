package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	AutoMigrate bool
	Port        string

	Debug    bool
	DbHost   string
	DbPort   string
	DbUser   string
	DbPass   string
	DbName   string
	DbSsl    string
	DbSchema string

	JWTSign       string
	JWTExpiration int
	RefreshTTL    int
)

func InitEnvVariables() error {
	var err error
	if err = envFileLookup(); err != nil {
		Debug, err = checkOSDebug()
		if err != nil {
			if err = loadDefaultEnv(); err != nil {
				log.Println("error loading .env.dev default variables")
				return err
			}
		}
	}
	Port = os.Getenv("PORT")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPass = os.Getenv("DB_PASS")
	DbName = os.Getenv("DB_NAME")
	DbSsl = os.Getenv("SSL_MODE")
	DbSchema = os.Getenv("DB_SCHEMA")

	JWTSign = os.Getenv("JWT_SIGN")
	JWTExpiration, _ = strconv.Atoi(os.Getenv("JWT_TTL"))
	RefreshTTL, _ = strconv.Atoi(os.Getenv("REFRESH_TTL"))

	AutoMigrate, _ = strconv.ParseBool(os.Getenv("AUTO_MIGRATE"))
	return err
}

func envFileLookup() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	_, err := godotenv.Read(".env")
	if err == nil {
		log.Println("loaded .env file environment variables")
	}
	return err
}

func checkOSDebug() (bool, error) {
	isDebug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	return isDebug, err
}

func loadDefaultEnv() error {
	err := godotenv.Load(".env.dev")
	_, err = godotenv.Read(".env.dev")
	if err == nil {
		log.Println(fmt.Sprintf("loaded env variables from .env.dev file"))
	}
	return err
}
