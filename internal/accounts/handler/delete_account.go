package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	account "system/internal/accounts/interface"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteAccountHandler(accountRepo account.Account_inteface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el ID de la cuenta de los parámetros de la URL
		accountID := c.Param("id")

		// Validar que el ID sea un UUID válido
		if _, err := uuid.Parse(accountID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "ID de cuenta inválido",
				"details": "El formato del ID no es válido. Debe ser un UUID",
			})
			return
		}

		// Eliminar la cuenta usando el repositorio
		err := accountRepo.Delete(c.Request.Context(), accountID)
		if err != nil {
			// Manejar diferentes tipos de errores
			switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				c.JSON(http.StatusNotFound, gin.H{
					"error":   "Cuenta no encontrada",
					"details": fmt.Sprintf("No existe una cuenta con el ID: %s", accountID),
				})

			case errors.Is(err, gorm.ErrForeignKeyViolated):
				c.JSON(http.StatusConflict, gin.H{
					"error":   "No se puede eliminar la cuenta",
					"details": "La cuenta tiene transacciones asociadas",
				})

			default:
				log.Printf("Error al eliminar cuenta: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Error interno del servidor",
					"details": "No se pudo eliminar la cuenta",
				})
			}
			return
		}

		// Respuesta exitosa sin contenido
		c.Status(http.StatusNoContent)
	}
}
