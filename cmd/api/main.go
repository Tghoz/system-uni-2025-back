package main

import (
	"errors"
	"net/http"
	"system/internal/auth/model"
	"system/internal/auth/repo/postgre"
	"system/internal/platform/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	router.POST("/users", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Datos inválidos"})
			return
		}

		// Usar el repositorio para guardar en la DB
		if err := userRepo.CreateUser(&user); err != nil {
			c.JSON(500, gin.H{"error": "Error al crear usuario"})
			return
		}

		c.JSON(201, user)
	})

	router.GET("/users/email", GetUserByEmail(userRepo))

	router.Run(":4000")
}

func GetUserByEmail(repo *postgre.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email") // Obtener el email de los query params
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro 'email' es requerido"})
			return
		}

		// Buscar usuario por email
		user, err := repo.GetUserByEmail(c.Request.Context(), email)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar usuario"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

/*	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := ":4000"
	router.Run(port) // listen and serve on
	//
	//
	// */
