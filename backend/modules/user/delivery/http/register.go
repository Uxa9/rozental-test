package http

import (
	"backend/internal/cmd/app"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndPoints(router *gin.RouterGroup, app *app.App) {
	h := NewHandler(app.UserModuleUC)

	endpoints := router.Group("user")
	{
		endpoints.POST("", h.CreateUser)
		endpoints.GET("", h.GetUser)
		endpoints.PUT("", h.UpdateUser)
		endpoints.DELETE("", h.DeleteUser)
	}
}
