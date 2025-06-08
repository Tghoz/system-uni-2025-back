// internal/accounts/routers/account_router.go
package routers

import (
	"system/internal/planning/handler"
	planning "system/internal/planning/interface"

	"github.com/gin-gonic/gin"
)

func PlanningRouters(router *gin.Engine, PlanningRepo planning.Planning_inteface) {
	r := router.Group("/api/v1/planning")
	r.POST("", handler.CreatePlanningHandler(PlanningRepo))
}
