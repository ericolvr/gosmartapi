package main

import (
	"github.com/ericolvr/goapi/internal/adapter/database"
	"github.com/ericolvr/goapi/internal/adapter/http"
	"github.com/ericolvr/goapi/internal/repository"
	"github.com/ericolvr/goapi/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db, err := database.NewMySQLConnection()
	if err != nil {
		panic("Failed to connect to MySQL database: " + err.Error())
	}

	userRepo := repository.NewMySQLUserRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo)

	http.NewUserHandler(router, userUsecase)

	router.Run(":8080")
}
