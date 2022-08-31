package config

import (
	"fmt"
	"github.com/istudko/go-project-template/logger"
	"os"

	"github.com/spf13/viper"
)

var config Config

type Config struct {
	appName string
	port    string
	log     *logger.LogConfig
	db      DBConfig
}

func Load() {
	viper.SetDefault("APP_PORT", "8080")

	viper.SetConfigName("application")
	if os.Getenv("ENVIRONMENT") == "test" {
		viper.SetConfigName("test")
	}

	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("unable to read config file: %v\n", err)
	}

	config = Config{
		appName: extractStringValue("APP_NAME"),
		port:    extractStringValue("APP_PORT"),
		log:     newLoggerConfig(),
		db:      newDBConfig(),
	}
}

func newLoggerConfig() *logger.LogConfig {
	return &logger.LogConfig{
		Out:    os.Stdout,
		Level:  extractStringValue("LOG_LEVEL"),
		Format: extractStringValue("LOG_FORMAT"),
	}
}

func GetAppName() string {
	return config.appName
}

func GetPort() string {
	return config.port
}

func GetLogger() *logger.LogConfig {
	return config.log
}

func GetDB() DBConfig {
	return config.db
}
