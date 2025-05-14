package handler

import (
	"system/internal/auth/repo"

	"github.com/gin-gonic/gin"
)

func GetAllUsersHandler(authRepo repo.Auth_Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := authRepo.GetAllUsers(c.Request.Context())
		if err != nil {
			c.JSON(500, gin.H{"error": "Error al obtener usuarios"})
			return
		}

		// Limpiar contraseñas antes de responder
		for i := range users {
			users[i].Password = "--"
		}

		c.JSON(200, users)
	}
}
