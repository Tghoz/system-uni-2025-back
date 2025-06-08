package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	models "system/internal/models"

	transaction "system/internal/transaction/interface"

	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTransactionHandler(transactionRepo transaction.Transaction_interface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Type      string  `json:"type"`
			Amount    float64 `json:"amount"`
			Currency  string  `json:"currency"`
			From      string  `json:"from"`
			To        string  `json:"to"`
			Status    string  `json:"status"`
			Reference string  `json:"reference"`
			AccountID string  `json:"account_id"`
		}

		// Bind del JSON
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Datos de solicitud inválidos",
				"details": "Verifique que todos los campos estén completos y en el formato correcto",
			})
			return
		}

		// Validar campos requeridos
		requiredFields := []string{input.Type, input.Currency, input.From, input.To, input.Status, input.Reference}
		requiredNames := []string{"type", "currency", "from", "to", "status", "reference"}

		for i, field := range requiredFields {
			if field == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Campos requeridos faltantes",
					"details": fmt.Sprintf("El campo '%s' es obligatorio", requiredNames[i]),
				})
				return
			}
		}

		// Validar monto positivo
		if input.Amount <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Monto inválido",
				"details": "El monto debe ser mayor que cero",
			})
			return
		}

		// Validar UUIDs para cuentas
		if _, err := uuid.Parse(input.From); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "ID de cuenta (origen) inválido",
				"details": "El ID de la cuenta de origen debe tener un formato UUID válido",
			})
			return
		}

		if _, err := uuid.Parse(input.To); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "ID de cuenta (destino) inválido",
				"details": "El ID de la cuenta de destino debe tener un formato UUID válido",
			})
			return
		}

		// Validar tipo de transacción
		validTypes := []string{"transfer", "deposit", "withdrawal", "payment", "refund"}
		if !contains(validTypes, input.Type) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Tipo de transacción inválido",
				"details": fmt.Sprintf("Tipos válidos: %s", strings.Join(validTypes, ", ")),
			})
			return
		}

		// Validar estado de transacción
		validStatuses := []string{"pending", "completed", "failed", "cancelled", "reversed"}
		if !contains(validStatuses, input.Status) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Estado de transacción inválido",
				"details": fmt.Sprintf("Estados válidos: %s", strings.Join(validStatuses, ", ")),
			})
			return
		}

		// Validar moneda (reutilizando función de cuentas)
		// if !isValidCurrencyType(input.Currency) {
		// 	validCurrencies := getValidCurrencyTypes()
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"error":   "Tipo de moneda inválido",
		// 		"details": fmt.Sprintf("Monedas válidas: %s", strings.Join(validCurrencies, ", ")),
		// 	})
		// 	return
		// }

		// Crear objeto Transaction
		transaction := models.Transaction{
			Type:      input.Type,
			Amount:    input.Amount,
			Currency:  input.Currency,
			From:      input.From,
			To:        input.To,
			Status:    input.Status,
			Reference: input.Reference,
			CreatedAt: time.Now().Format(time.RFC3339), // Generar timestamp actual
			AccountID: input.AccountID,
		}

		// Intentar crear la transacción
		if err := transactionRepo.Create(c.Request.Context(), &transaction); err != nil {
			// Manejar errores específicos
			switch {
			case strings.Contains(err.Error(), "violates foreign key constraint"):
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Cuenta no encontrada",
					"details": "Las cuentas de origen o destino no existen en el sistema",
				})
			case strings.Contains(err.Error(), "insufficient funds"):
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Fondos insuficientes",
					"details": "La cuenta de origen no tiene suficiente saldo",
				})
			default:
				log.Printf("Error interno al crear transacción: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Error interno del servidor",
					"details": "No se pudo completar la operación",
				})
			}
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":        "Transacción creada exitosamente",
			"transaction_id": transaction.ID,
			"reference":      transaction.Reference,
			"status":         transaction.Status,
		})
	}
}

// Función auxiliar para verificar valores en slice
func contains(slice []string, item string) bool {
	return slices.Contains(slice, item)
}
