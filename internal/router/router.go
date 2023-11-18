package router

import (
	"github.com/akmyrzza/yandex-url/internal/Middleware"
	"github.com/akmyrzza/yandex-url/internal/handler"
	"github.com/akmyrzza/yandex-url/internal/logger"
	"github.com/gin-gonic/gin"
)

func InitRouter(h *handler.Handler, lg *logger.Logger) *gin.Engine {
	router := gin.Default()

	router.Use(Middleware.Logging(lg))

	router.POST("/", h.ShortURL)
	router.GET("/:id", h.OriginalURL)

	return router
}
