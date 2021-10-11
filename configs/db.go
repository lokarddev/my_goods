package configs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

func NewDatabaseConf() *DbConfig {
	return &DbConfig{
		Host:     DbHost,
		Port:     DbPort,
		DbName:   DbName,
		Username: DbUser,
		Password: DbPass,
	}
}

func InitDB(cfg *DbConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8mb4", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
	}
	return db
}
