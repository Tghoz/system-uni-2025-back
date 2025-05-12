package main

import (
	"system/internal/auth/handler"
	"system/internal/auth/model"
	"system/internal/auth/repo/postgre"
	"system/internal/platform/db"

	"github.com/gin-gonic/gin"
)

func main() {

	database := db.ConnectDB()
	database.AutoMigrate(&model.User{})

	repo := postgre.NewUserRepository(database)

	setupRouter(repo)

	defer db.DisconnectDB(database)

}

func setupRouter(userRepo *postgre.UserRepository) {
	router := gin.Default()
	// Ruta para crear un usuario
	router.POST("/users", handler.CreateUserHandler(userRepo))
	
	router.GET("/users/email", handler.GetUserByEmail(userRepo))

	router.Run(":4000")
}
