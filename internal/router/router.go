package router

import (
	"github.com/akmyrzza/yandex-url/internal/handler"
	"github.com/akmyrzza/yandex-url/internal/logger"
	"github.com/akmyrzza/yandex-url/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(h *handler.Handler, lg *logger.Logger) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.LoggingRequest(lg))

	router.POST("/", h.ShortURL)
	router.GET("/:id", h.OriginalURL)

	return router
}
