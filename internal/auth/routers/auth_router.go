// internal/auth/routers/user_router.go
package routers

import (
	"system/internal/auth/handler"
	"system/internal/auth/repo" // Importar paquete de la interfaz

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine, authRepo repo.Auth_Repo) { // Usar interfaz aquí
	r := router.Group("/api/v1/user")

	// Handlers reciben la interfaz
	r.POST("/login", handler.LoginUserHandler(authRepo))
	r.POST("/register", handler.CreateUserHandler(authRepo))

}
