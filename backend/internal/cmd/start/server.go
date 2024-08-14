package start

import (
	"backend/internal/cmd/app"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func StartServer(app *app.App) error {

	router := gin.New()

	router.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				defer func() {
					if err := recover(); err != nil {
						c.AbortWithStatusJSON(
							http.StatusInternalServerError,
							gin.H{"error": "Ошибка при инициализации роутера"},
						)
					}
				}()
			}
		}()
		c.Next()
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//api := router.Group("/api")

	httpServer := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Сервер запущен на " + httpServer.Addr + " порту")

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return httpServer.Shutdown(ctx)
}
