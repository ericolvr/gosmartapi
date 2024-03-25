package main

import (
	"github.com/ericolvr/goapi/internal/adapter/database"
	"github.com/ericolvr/goapi/internal/adapter/http"
	"github.com/ericolvr/goapi/internal/repository"
	"github.com/ericolvr/goapi/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize router
	router := gin.Default()

	// initialize mysql connection
	db, err := database.NewMySQLConnection("root:secret@tcp(localhost:3306)/goapi")
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
