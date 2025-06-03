package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	account "system/internal/accounts/interface"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	// Necesitas importar uuid para trabajar con los IDs de modelo
)

// Funciones auxiliares (asumiendo que existen en tu paquete handler o son globales si es necesario)

func UpdateAccountHandler(accountRepo account.Account_inteface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener ID de la cuenta de los parámetros de la URL
		accountID := c.Param("id")

		// Validar que el ID sea un UUID válido
		if _, err := uuid.Parse(accountID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "ID de cuenta inválido",
				"details": "El formato del ID no es válido. Debe ser un UUID",
			})
			return
		}

		// Obtener la cuenta desde la base de datos
		existingAccount, err := accountRepo.GetByID(c.Request.Context(), accountID)
		if err != nil {
			log.Printf("Error al obtener cuenta: %v", err)

			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{
					"error":   "Cuenta no encontrada",
					"details": fmt.Sprintf("No se encontró una cuenta con ID: %s", accountID),
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Error interno",
					"details": "No se pudo recuperar la cuenta",
				})
			}
			return
		}

		// Bind de los datos de actualización
		var updateData struct {
			CurrencyType *string  `json:"currency"` // Puntero para diferenciar entre no enviado y valor cero
			Amount       *float64 `json:"amount"`
		}

		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Datos inválidos",
				"details": "Verifique el formato de los campos",
			})
			return
		}

		// Validar que al menos un campo sea proporcionado
		if updateData.CurrencyType == nil && updateData.Amount == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Datos insuficientes",
				"details": "Debe proporcionar al menos un campo para actualizar (currency o amount)",
			})
			return
		}

		// Aplicar cambios a la cuenta existente
		changed := false

		// Actualizar moneda si se proporciona
		if updateData.CurrencyType != nil {
			existingAccount.CurrencyType = *updateData.CurrencyType
			changed = true
		}

		// Actualizar amount si se proporciona
		if updateData.Amount != nil {
			existingAccount.Amount = *updateData.Amount
			changed = true
		}

		// Si no hay cambios reales
		if !changed {
			c.JSON(http.StatusOK, gin.H{
				"message": "No se detectaron cambios para aplicar",
				"account": gin.H{
					"id":           existingAccount.ID,
					"account_type": existingAccount.AccountType,
					"currency":     existingAccount.CurrencyType,
					"amount":       existingAccount.Amount,
				},
			})
			return
		}

		// Actualizar marca de tiempo

		// Aplicar actualización
		if err := accountRepo.Update(c.Request.Context(), existingAccount); err != nil {
			log.Printf("Error al actualizar cuenta: %v", err)

			// Manejar diferentes tipos de errores
			var errorMsg string
			var statusCode int

			switch {
			case strings.Contains(err.Error(), "violates foreign key constraint"):
				errorMsg = "La moneda especificada no existe en el sistema"
				statusCode = http.StatusBadRequest
			case strings.Contains(err.Error(), "duplicate key value"):
				errorMsg = "Conflicto de datos únicos"
				statusCode = http.StatusConflict
			default:
				errorMsg = "Error interno del servidor"
				statusCode = http.StatusInternalServerError
			}

			c.JSON(statusCode, gin.H{
				"error":   "Error al actualizar cuenta",
				"details": errorMsg,
			})
			return
		}

		// Crear respuesta
		response := gin.H{
			"message": "Cuenta actualizada exitosamente",
			"account": gin.H{
				"id":           existingAccount.ID,
				"account_type": existingAccount.AccountType,
				"currency":     existingAccount.CurrencyType,
				"amount":       existingAccount.Amount,
			},
		}

		c.JSON(http.StatusOK, response)
	}
}
