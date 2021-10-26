package init

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"my_goods/internal/entity"
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

func DB(cfg *DbConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.Username, cfg.Password, cfg.DbName, cfg.Port, cfg.SslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
	}
	err = db.AutoMigrate(&entity.Goods{})
	err = db.AutoMigrate(&entity.Dish{})
	err = db.AutoMigrate(&entity.List{})
	if err != nil {
		logrus.Error(err)
	}
	return db
}
