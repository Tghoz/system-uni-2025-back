package routers

import (
	auth_routers "system/internal/auth/routers"
	"system/internal/repo" // Importa el paquete del contenedor

	account_routers "system/internal/accounts/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cambia el parámetro para aceptar el contenedor completo
func SetupRouter(container *repo.RepositoryContainer) {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4321"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Pasa el repositorio específico que necesitan los routers
	auth_routers.UserRouter(r, container.User)
	auth_routers.AdminRouter(r, container.User)

	account_routers.AccountRouters(r, container.Account)

	r.Run(":4000")
}
