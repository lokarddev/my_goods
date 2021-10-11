package configs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"my_goods/internal/models"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SslMode  string
}

func NewDatabaseConf() *DbConfig {
	return &DbConfig{
		Host:     DbHost,
		Port:     DbPort,
		DbName:   DbName,
		Username: DbUser,
		Password: DbPass,
		SslMode:  SslMode,
	}
}

func InitDB(cfg *DbConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.Username, cfg.Password, cfg.DbName, cfg.Port, cfg.SslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
	}
	err = db.AutoMigrate(&models.Goods{})
	err = db.AutoMigrate(&models.Dish{})
	err = db.AutoMigrate(&models.List{})
	if err != nil {
		logrus.Error(err)
	}
	return db
}
