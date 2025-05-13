package handler

import (
	"system/internal/auth/model"
	"system/internal/auth/repo"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(authRepo repo.Auth_Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Datos inválidos"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		if err != nil {
			c.JSON(500, gin.H{"error": "Error al procesar la contraseña"})
			return
		}
		user.Password = string(hashedPassword)

		if err := authRepo.CreateUser(c.Request.Context(), &user); err != nil {
			c.JSON(500, gin.H{"error": "Error al crear usuario"})
			return
		}
		c.JSON(201, user)
	}
}
