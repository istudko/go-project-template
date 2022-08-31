package config

import (
	"fmt"
	"github.com/istudko/go-project-template/db"
	"time"

)

const dbDriver = "postgres"

type DBConfig struct {
	databaseName        string
	databaseHost        string
	databaseUser        string
	databasePassword    string
	databasePort        int
	databaseMaxPoolSize int
	ReadTimeout         time.Duration
	WriteTimeout        time.Duration
}

func newDBConfig() DBConfig {
	return DBConfig{
		databaseName:        extractStringValue("DB_NAME"),
		databaseHost:        extractStringValue("DB_HOSTNAME"),
		databaseUser:        extractStringValue("DB_USERNAME"),
		databasePassword:    extractStringValue("DB_PASSWORD"),
		databasePort:        extractIntValue("DB_PORT"),
		databaseMaxPoolSize: extractIntValue("DB_POOL"),
		ReadTimeout:         time.Millisecond * time.Duration(extractIntValue("DB_READ_TIMEOUT_MS")),
		WriteTimeout:        time.Millisecond * time.Duration(extractIntValue("DB_WRITE_TIMEOUT_MS")),
	}
}

func (dc DBConfig) ConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dc.databaseUser, dc.databasePassword, dc.databaseHost, dc.databasePort, dc.databaseName)
}

func (dc DBConfig) DatabaseMaxPoolSize() int {
	return dc.databaseMaxPoolSize
}

func (dc DBConfig) DatabaseConfig() *db.Config {
	return &db.Config{
		Driver:          dbDriver,
		URL:             GetDB().ConnectionURL(),
		MaxIdleConns:    GetDB().DatabaseMaxPoolSize(),
		MaxOpenConns:    GetDB().DatabaseMaxPoolSize(),
		ConnMaxLifeTime: 30 * time.Minute,
	}
}

func (dc DBConfig) MigrationConfig() *db.MigrationConfig {
	return &db.MigrationConfig{
		Driver: dbDriver,
		URL:    dc.ConnectionURL(),
	}
}
