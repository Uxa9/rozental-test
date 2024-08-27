package http

import (
	"backend/internal/cmd/app"
	"github.com/gin-gonic/gin"
)

func RegiterHTTPEndPoints(router *gin.RouterGroup, app *app.App) {
	h := NewHandler(app.StatusUC)

	endpoints := router.Group("status")
	{
		endpoints.PUT("", h.CreateStatus)
		endpoints.GET("", h.GetStatus)
	}
}
