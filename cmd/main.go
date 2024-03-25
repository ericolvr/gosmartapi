package main

import (
	"log"
	"os"

	"github.com/ericolvr/goapi/internal/adapter/database"
	"github.com/ericolvr/goapi/internal/adapter/http"
	"github.com/ericolvr/goapi/internal/repository"
	"github.com/ericolvr/goapi/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// initialize router
	router := gin.Default()

	// load variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_NAME")
	connString := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + databaseName

	// initialize mysql connection
	db, err := database.NewMySQLConnection(connString)
	if err != nil {
		panic("Failed to connect to MySQL database: " + err.Error())
	}

	// Initialize user repository
	userRepo := repository.NewMySQLUserRepository(db)

	// Initialize user use case
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Initialize HTTP handlers
	http.NewUserHandler(router, userUsecase)

	// Start server
	router.Run(":8080")
}
