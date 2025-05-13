package routers

import (
	"system/internal/auth/repo"
	"system/internal/auth/routers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userRepo repo.Auth_Repo) {
	r := gin.Default()

	routers.UserRouter(r, userRepo)

	r.Run(":4000")
}
