package helpers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleUserError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar la solicitud"})
	}
}
