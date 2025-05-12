package handler

import (
	"system/internal/auth/model"
	"system/internal/auth/repo/postgre"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(repo *postgre.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Datos inválidos"})
			return
		}
		// Usar el repositorio para guardar en la DB
		if err := repo.CreateUser(&user); err != nil {
			c.JSON(500, gin.H{"error": "Error al crear usuario"})
			return
		}
		c.JSON(201, user)
	}
}
