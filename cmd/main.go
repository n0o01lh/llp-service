package main

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"github.com/n0o01lh/llp/internals/core/services"
	db_configuration "github.com/n0o01lh/llp/internals/db"
	"github.com/n0o01lh/llp/internals/handlers"
	"github.com/n0o01lh/llp/internals/repositories"
	"github.com/n0o01lh/llp/internals/server"
)

func main() {

	var envFile string
	if len(os.Args) > 1 {
		envFile = os.Args[1]
	} else {
		envFile = ".env"
	}

	err := godotenv.Load(envFile)

	if err != nil {
		log.Error(err)
		panic(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	port, _ := strconv.Atoi(dbPort)

	//database connection
	db_configuration.Connect(dbHost, dbUser, dbPassword, port)

	resourceRepository := repositories.NewResourceRepository(db_configuration.Database)
	resourceService := services.NewResourceService(resourceRepository)
	resourceHandlers := handlers.NewResourceHandlers(resourceService)

	server := server.NewServer(resourceHandlers)

	server.Initialize()
}
