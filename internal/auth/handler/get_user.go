package handler

import (
	"errors"
	"net/http"
	"system/internal/auth/repo/postgre"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
