package http

import (
	"backend/internal/cmd/app"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndPoints(router *gin.RouterGroup, app *app.App) {
	h := NewHandler(app.InventoryOperationUC)

	endpoints := router.Group("inventory-operation")
	{
		endpoints.POST("/getOperations", h.GetInventoryOperation)
		endpoints.POST("/getInventory", h.GetInventory)
	}
}
