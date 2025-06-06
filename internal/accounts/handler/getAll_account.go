package handler

import (
	"log"
	"net/http"
	account "system/internal/accounts/interface"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllAccountsHandler(accountRepo account.Account_inteface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener todas las cuentas del repositorio
		accounts, err := accountRepo.GetAll(c.Request.Context())
		if err != nil {
			log.Printf("Error al obtener cuentas: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error interno del servidor",
				"details": "No se pudo recuperar la lista de cuentas",
			})
			return
		}

		// Si no hay cuentas, retornar array vacío en lugar de null
		if len(accounts) == 0 {
			c.JSON(http.StatusOK, []interface{}{})
			return
		}

		// Crear respuesta estructurada para cada cuenta
		response := make([]gin.H, 0, len(accounts))
		for _, acc := range accounts {
			accountData := gin.H{
				"id":           acc.ID,
				"account_type": acc.AccountType,
				"currency":     acc.CurrencyType,
				"balance":      acc.Amount,
				"created_at":   acc.CreatedAt.Format(time.RFC3339),
				"user_id":      acc.UserID,
			}
			response = append(response, accountData)
		}

		c.JSON(http.StatusOK, response)
	}
}
