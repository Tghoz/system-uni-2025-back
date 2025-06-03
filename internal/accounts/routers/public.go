// internal/accounts/routers/account_router.go
package routers

import (
	"system/internal/accounts/handler"
	account "system/internal/accounts/interface"

	"github.com/gin-gonic/gin"
)

func AccountRouters(router *gin.Engine, accountRepo account.Account_inteface) {
	r := router.Group("/api/v1/accounts")
	r.POST("", handler.CreateAccountHandler(accountRepo))
	r.GET("/:id", handler.GetAccountHandler(accountRepo))
	r.PUT("/:id", handler.UpdateAccountHandler(accountRepo))
}
