package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	account "system/internal/accounts/interface"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAccountHandler(accountRepo account.Account_inteface) gin.HandlerFunc {
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

		// Obtener la cuenta del repositorio
		account, err := accountRepo.GetByID(c.Request.Context(), accountID)
		if err != nil {
			// Manejar diferentes tipos de errores
			switch {	
			case errors.Is(err, gorm.ErrRecordNotFound):
				c.JSON(http.StatusNotFound, gin.H{
					"error":   "Cuenta no encontrada",
					"details": fmt.Sprintf("No existe una cuenta con el ID: %s", accountID),
				})
				
			case strings.Contains(err.Error(), "invalid input syntax for type uuid"):
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "ID inválido",
					"details": "El ID proporcionado no tiene un formato válido",
				})
			default:
				log.Printf("Error al obtener cuenta: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Error interno del servidor",
					"details": "No se pudo recuperar la información de la cuenta",
				})
			}
			return
		}

		// Crear respuesta estructurada
		response := gin.H{
			"id":           account.ID,
			"account_type": account.AccountType,
			"currency":     account.CurrencyType,
			"balance":      account.Amount,
			"created_at":   account.CreatedAt.Format(time.RFC3339),
			"user_id":      account.UserID,
		}

		c.JSON(http.StatusOK, response)
	}
}
