# go-project-template

A Simple template for golang project

### Prerequisite

* go 1.17 or above

### Development setup

1. `docker-compose up -d` - Start project dependencies
2. `cp application.yml.sample application.yml` - Create application config
3. `go build -o out/app` - Compile and build an application
4. `./out/app migrate` - Run database migration
5. `./out/app start` - Run application

### Development teardown

1. Exit application
2. `docker-compose down` - Stop project dependencies
