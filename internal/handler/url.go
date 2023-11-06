package handler

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShortURL(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось прочитать тело запроса"})
		return
	}

	OriginalURL := strings.TrimSpace(string(body))
	ShortURL := GetRandomURL()
	h.StorageURL[ShortURL] = OriginalURL

	ctx.Header("Content-Type", "text/plain")
	ctx.String(http.StatusCreated, h.BaseURL + ShortURL)
}

func (h *Handler) OriginalURL(ctx *gin.Context) {
	ShortURL, exist := ctx.Params.Get("id")
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить ID"})
		return
	}

	OriginalURL, ok := h.StorageURL[ShortURL]
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить оригинальную ссылку"})
		return
	}

	ctx.Header("Location", OriginalURL)
	ctx.Redirect(http.StatusTemporaryRedirect, OriginalURL)
}

