// internal/auth/routers/admin_router.go
package routers

import (
	handler "system/internal/auth/handler/protected_handler"
	"system/internal/auth/middleware" // Importar el middleware
	"system/internal/auth/repo"

	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine, authRepo repo.Auth_Repo) {
	// Grupo de rutas protegidas
	admin := router.Group("/api/v1/admin")

	// Aplicar middleware de autenticación a TODAS las rutas admin
	admin.Use(middleware.AuthMiddleware()) // <-- Middleware global para el grupo

	// Rutas
	admin.GET("/users", handler.GetAllUsersHandler(authRepo)) // Protegida
}
