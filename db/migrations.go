package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // get db migration from path
	"github.com/pkg/errors"
)

const defaultMigrationsPath = "file://./migration"

var ErrNoMigrations = errors.New("no migrations")
var ErrFindingDriver = errors.New("no migrate driver instance found")

type MigrationConfig struct {
	Driver string
	URL    string
	Path   string
}

func (cfg MigrationConfig) MigrationPath() string {
	if cfg.Path == "" {
		return defaultMigrationsPath
	}
	return cfg.Path
}

func RunDatabaseMigrations(config *MigrationConfig) error {
	db, err := sql.Open(config.Driver, config.URL)
	if err != nil {
		return err
	}

	driver, err := getDBDriverInstance(db, config.Driver)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(config.MigrationPath(), config.Driver, driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err == migrate.ErrNoChange || err == nil {
		return nil
	}

	return err
}

func RollbackLatestMigration(config *MigrationConfig) error {
	m, err := migrate.New(config.MigrationPath(), config.URL)
	if err != nil {
		return err
	}

	err = m.Steps(-1)
	if err == migrate.ErrNoChange || err == nil {
		return nil
	}

	return err
}

func CreateMigration(filename string, config *MigrationConfig) error {
	if len(filename) == 0 {
		return errors.New("filename is not provided")
	}

	timeStamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", config.MigrationPath(), timeStamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", config.MigrationPath(), timeStamp, filename)

	if err := createFile(upMigrationFilePath); err != nil {
		return err
	}
	fmt.Printf("created %s\n", upMigrationFilePath)

	if err := createFile(downMigrationFilePath); err != nil {
		os.Remove(upMigrationFilePath)
		return err
	}

	fmt.Printf("created %s\n", downMigrationFilePath)

	return nil
}

func getDBDriverInstance(db *sql.DB, driver string) (database.Driver, error) {
	switch driver {
	case "postgres":
		return postgres.WithInstance(db, &postgres.Config{})
	default:
		return nil, ErrFindingDriver
	}
}

func createFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	err = f.Close()

	return err
}
