package handler

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShortURL(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось прочитать тело запроса"})
		return
	}

	url := h.service.MakeShortPath(h.BaseURL, body)

	ctx.Header("Content-Type", "text/plain")
	ctx.String(http.StatusCreated, url)
}

func (h *Handler) OriginalURL(ctx *gin.Context) {
	ShortURL, exist := ctx.Params.Get("id")
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить ID"})
		return
	}

	url, ok := h.service.GetOriginalURL(ShortURL)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить оригинальную ссылку"})
		return
	}

	ctx.Header("Location", url)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}
