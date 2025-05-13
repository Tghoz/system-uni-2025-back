package handler

import (
	"net/http"
	"system/internal/auth/repo"
	"system/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func GetUserByEmail(authRepo repo.Auth_Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro 'email' es requerido"})
			return
		}

		user, err := authRepo.GetUserByEmail(c.Request.Context(), email)
		if err != nil {
			helpers.HandleUserError(c, err)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
