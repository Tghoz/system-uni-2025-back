package routers

import (
	"system/internal/auth/repo"
	"system/internal/auth/routers"
	potect "system/internal/auth/routers/protected_routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userRepo repo.Auth_Repo) {
	r := gin.Default()

	r.Use(cors.Default())

	routers.UserRouter(r, userRepo)
	potect.AdminRouter(r, userRepo)

	r.Run(":4000")
}
