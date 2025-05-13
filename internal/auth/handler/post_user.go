package handler

import (
	"system/internal/auth/model"
	"system/internal/auth/repo"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(authRepo repo.Auth_Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Datos inválidos"})
			return
		}
		// Usar el repositorio para guardar en la DB
		if err := authRepo.CreateUser(c.Request.Context(), &user); err != nil {
			c.JSON(500, gin.H{"error": "Error al crear usuario"})
			return
		}
		c.JSON(201, user)
	}
}
