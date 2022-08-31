package main

import (
	"github.com/istudko/go-project-template/config"
	"github.com/istudko/go-project-template/db"
	"github.com/istudko/go-project-template/logger"
	"github.com/istudko/go-project-template/server"
	"github.com/istudko/go-project-template/validator"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// setup will load configuration and initialize application dependency
func setup() {
	config.Load()
	logger.SetupLogger(config.GetLogger())
	validator.Init()
	if err := db.Init(config.GetDB().DatabaseConfig()); err != nil {
		log.Panicf("failure to initialize db connection: %+v", err)
	}
}

// teardown will decommission application dependency
func teardown() {
	_ = db.Close()
}

func main() {
	setup()
	defer teardown()

	cliApp := cli.NewApp()
	cliApp.Commands = []*cli.Command{
		{
			Name:        "start",
			Description: "Start the service",
			Action: func(c *cli.Context) error {
				server.Start()
				return nil
			},
		},
		{
			Name:        "migrate",
			Description: "Run Database migration",
			Action: func(c *cli.Context) error {
				return db.RunDatabaseMigrations(config.GetDB().MigrationConfig())
			},
		},
		{
			Name:        "rollback",
			Description: "Rollback latest database migration",
			Action: func(c *cli.Context) error {
				return db.RollbackLatestMigration(config.GetDB().MigrationConfig())
			},
		},
	}
	// Run app via command line
	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}
