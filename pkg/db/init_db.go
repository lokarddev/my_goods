package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
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

	err = migrationsUp()

	return db, err
}

func migrationsUp() error {
	dsnMigrations := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		env.DbUser, env.DbPass, env.DbHost, env.DbPort, env.DbName, env.DbSsl)
	mDB, err := sql.Open("postgres", dsnMigrations)
	driver, err := postgres.WithInstance(mDB, &postgres.Config{
		MigrationsTable:       "my_goods_migrations",
		MigrationsTableQuoted: false,
		MultiStatementEnabled: false,
		DatabaseName:          env.DbName,
		SchemaName:            "",
		StatementTimeout:      0,
		MultiStatementMaxSize: 0,
	})
	m, err := migrate.NewWithDatabaseInstance("file://migrations/", env.DbName, driver)
	err = m.Up()
	return err
}
