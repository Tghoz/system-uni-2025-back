package handler

import (
	"log"
	"net/http"
	"strings"

	planning "system/internal/planning/interface"

	models "system/internal/models"

	"github.com/gin-gonic/gin"
)

func CreatePlanningHandler(planningRepo planning.Planning_inteface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Name          string  `json:"name"`
			Description   string  `json:"description"`
			Amount        float64 `json:"amount"`
			Service       string  `json:"service"`
			Value         float64 `json:"value"`
			Month         string  `json:"date"`
			TransactionID string  `json:"transaction_id"`
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
		if input.Name == "" || input.Service == "" || input.Month == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Campos requeridos faltantes",
				"details": "Los campos 'name', 'service' y 'date' son obligatorios",
			})
			return
		}

		// Validar valores numéricos
		if input.Amount < 0 || input.Value < 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Valores numéricos inválidos",
				"details": "Los campos 'amount' y 'value' deben ser valores positivos",
			})
			return
		}

		// Crear objeto Planning
		plan := models.Planning{
			Name:          input.Name,
			Description:   input.Description,
			Amount:        input.Amount,
			Service:       input.Service,
			Value:         input.Value,
			Month:         input.Month,
			TransactionID: input.TransactionID,
		}

		// Intentar crear el planning
		if err := planningRepo.Create(c.Request.Context(), &plan); err != nil {
			// Manejar errores específicos
			switch {
			case strings.Contains(err.Error(), "duplicate key"):
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Registro duplicado",
					"details": "Ya existe un planning con estos datos",
				})
			case strings.Contains(err.Error(), "invalid service"):
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Servicio no válido",
					"details": "Por favor seleccione un servicio de la lista permitida",
				})
			default:
				log.Printf("Error interno al crear planning: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Error interno del servidor",
					"details": "No se pudo completar la operación",
				})
			}
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":     "Planning creado exitosamente",
			"planning_id": plan.ID,
			"name":        plan.Name,
			"date":        plan.Month,
		})
	}
}
