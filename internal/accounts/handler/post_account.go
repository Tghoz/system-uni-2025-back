// system/internal/handler/account_handler.go
package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	account "system/internal/accounts/interface"

	models "system/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateAccountHandler(accountRepo account.Account_inteface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			AccountType  string  `json:"account"`
			CurrencyType string  `json:"currency"`
			Amount       float64 `json:"balance"`
			UserID       string  `json:"user_id"`
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
		if input.AccountType == "" || input.CurrencyType == "" || input.UserID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Campos requeridos faltantes",
				"details": "Los campos 'account', 'currency' y 'user_id' son obligatorios",
			})
			return
		}

		// Validar formato UUID
		if _, err := uuid.Parse(input.UserID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "ID de usuario inválido",
				"details": "El ID de usuario debe tener un formato UUID válido (ej: 550e8400-e29b-41d4-a716-446655440000)",
			})
			return
		}

		// Validar tipo de cuenta
		if !isValidAccountType(input.AccountType) {
			validTypes := getValidAccountTypes()
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Tipo de cuenta inválido",
				"details": fmt.Sprintf("Los tipos válidos son: %s", strings.Join(validTypes, ", ")),
			})
			return
		}

		// Validar tipo de moneda
		if !isValidCurrencyType(input.CurrencyType) {
			validCurrencies := getValidCurrencyTypes()
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Tipo de moneda inválido",
				"details": fmt.Sprintf("Las monedas válidas son: %s", strings.Join(validCurrencies, ", ")),
			})
			return
		}

		// Crear objeto Account DTO
		account := models.Account{
			AccountType:  input.AccountType,
			CurrencyType: input.CurrencyType,
			Amount:       input.Amount,
			UserID:       input.UserID,
		}

		// Intentar crear la cuenta
		if err := accountRepo.Create(c.Request.Context(), &account); err != nil {
			// Manejar errores específicos
			switch {
			case strings.Contains(err.Error(), "violates foreign key constraint"):
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Usuario no encontrado",
					"details": "El ID de usuario proporcionado no existe en el sistema",
				})
			case strings.Contains(err.Error(), "invalid account type"):
				// Este caso debería estar cubierto por la validación previa
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Tipo de cuenta no válido",
					"details": "Por favor seleccione un tipo de cuenta válido",
				})
			default:
				// Log interno para desarrollo/depuración
				log.Printf("Error interno al crear cuenta: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Error interno del servidor",
					"details": "No se pudo completar la operación. Por favor intente más tarde",
				})
			}
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":    "Cuenta creada exitosamente",
			"account_id": account.ID,
			"balance":    account.Amount,
		})
	}
}

// Funciones auxiliares para validación
func isValidAccountType(accountType string) bool {
	validTypes := map[string]bool{
		"checking": true,
		"savings":  true,
		"credit":   true,
	}
	return validTypes[accountType]
}

func isValidCurrencyType(currency string) bool {
	validCurrencies := map[string]bool{
		"USD": true,
		"EUR": true,
		"COP": true,
		"Bs":  true,
	}
	return validCurrencies[currency]
}

func getValidAccountTypes() []string {
	return []string{"checking", "savings", "credit"}
}

func getValidCurrencyTypes() []string {
	return []string{"USD", "EUR", "COP", "Bs"}
}
