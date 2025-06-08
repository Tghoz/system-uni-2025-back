// internal/accounts/routers/account_router.go
package routers

import (
	"system/internal/transaction/handler"
	transaction "system/internal/transaction/interface"

	"github.com/gin-gonic/gin"
)

func TransactionRouters(router *gin.Engine, TransactionRepo transaction.Transaction_interface) {
	r := router.Group("/api/v1/transaction")
	r.POST("", handler.CreateTransactionHandler(TransactionRepo))
	r.GET("", handler.GetAllTransactionsHandler(TransactionRepo))
	r.GET("/:id", handler.GetTransactionHandler(TransactionRepo))
}
