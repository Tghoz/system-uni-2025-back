package handler

import (
	"log"
	"net/http"
	transaction "system/internal/transaction/interface"

	"github.com/gin-gonic/gin"
)

func GetAllTransactionsHandler(transactionRepo transaction.Transaction_interface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener todas las transacciones del repositorio
		transactions, err := transactionRepo.GetAll(c.Request.Context())
		if err != nil {
			log.Printf("Error al obtener transacciones: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error interno del servidor",
				"details": "No se pudo recuperar la lista de transacciones",
			})
			return
		}

		// Si no hay transacciones, retornar array vacío
		if len(transactions) == 0 {
			c.JSON(http.StatusOK, []any{})
			return
		}

		// Crear respuesta estructurada para cada transacción
		response := make([]gin.H, 0, len(transactions))
		for _, tx := range transactions {
			transactionData := gin.H{
				"id":         tx.ID,
				"type":       tx.Type,
				"created_at": tx.CreatedAt,
				"amount":     tx.Amount,
				"currency":   tx.Currency,
				"from":       tx.From,
				"to":         tx.To,
				"status":     tx.Status,
				"reference":  tx.Reference,
				"account_id": tx.AccountID,
			}
			response = append(response, transactionData)
		}

		c.JSON(http.StatusOK, response)
	}
}
