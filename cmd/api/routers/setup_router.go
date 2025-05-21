package routers

import (
	"system/internal/auth/repo"
	"system/internal/auth/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userRepo repo.Auth_Repo) {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4321"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // 🔥 Esto es lo clave

	}))

	routers.UserRouter(r, userRepo)
	routers.AdminRouter(r, userRepo)

	r.Run(":4000")
}
