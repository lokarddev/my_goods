package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"my_goods/pkg/env"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	DbSchema string
	SSLMode  string
}

func NewDatabasePostgres() (*pgxpool.Pool, error) {
	cfg := DatabaseConfig{
		Host:     env.DbHost,
		Port:     env.DbPort,
		Username: env.DbUser,
		Password: env.DbPass,
		DBName:   env.DbName,
		DbSchema: env.DbSchema,
		SSLMode:  env.DbSsl,
	}
	dsnDB := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s&sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.DbSchema, cfg.SSLMode)
	e, err := pgxpool.ParseConfig(dsnDB)
	e.MaxConns = int32(20)
	if err != nil {
		log.Fatalf("ERROR INITIALIZING DB[%s]: %s", cfg.DBName, err.Error())
	}
	db, err := pgxpool.ConnectConfig(context.Background(), e)
	if err != nil {
		log.Fatalf("ERROR INITIALIZING DB[%s]: %s", cfg.DBName, err.Error())
	}
	log.Printf("SUCCESSFUL CONNECTION TO DB[%s]\n", cfg.DBName)
	return db, err
}
