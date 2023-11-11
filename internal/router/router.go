package router

import (
	"github.com/akmyrzza/yandex-url/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.POST("/", h.ShortURL)
	router.GET("/:id", h.OriginalURL)

	return router
}
