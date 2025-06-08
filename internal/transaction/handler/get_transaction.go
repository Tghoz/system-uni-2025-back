package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	transaction "system/internal/transaction/interface"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetTransactionHandler(transactionRepo transaction.Transaction_interface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el ID de la transacción de los parámetros de la URL
		transactionID := c.Param("id")

		// Validar que el ID sea un UUID válido
		if _, err := uuid.Parse(transactionID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "ID de transacción inválido",
				"details": "El formato del ID no es válido. Debe ser un UUID",
			})
			return
		}

		// Obtener la transacción del repositorio
		txn, err := transactionRepo.GetByID(c.Request.Context(), transactionID)
		if err != nil {
			// Manejar diferentes tipos de errores
			switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				c.JSON(http.StatusNotFound, gin.H{
					"error":   "Transacción no encontrada",
					"details": fmt.Sprintf("No existe una transacción con el ID: %s", transactionID),
				})
			case strings.Contains(err.Error(), "invalid input syntax for type uuid"):
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "ID inválido",
					"details": "El ID proporcionado no tiene un formato válido",
				})
			default:
				log.Printf("Error al obtener transacción: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Error interno del servidor",
					"details": "No se pudo recuperar la información de la transacción",
				})
			}
			return
		}

		// Crear respuesta estructurada
		response := gin.H{
			"id":         txn.ID,
			"type":       txn.Type,
			"created_at": txn.CreatedAt, // Ya es string según tu modelo
			"amount":     txn.Amount,
			"currency":   txn.Currency,
			"from":       txn.From,
			"to":         txn.To,
			"status":     txn.Status,
			"reference":  txn.Reference,
			"account_id": txn.AccountID,
		}

		c.JSON(http.StatusOK, response)
	}
}
