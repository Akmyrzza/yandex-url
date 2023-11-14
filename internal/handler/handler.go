package handler

import "github.com/akmyrzza/yandex-url/internal/shortener"

type Handler struct {
	service shortener.Service
	BaseURL string
}

func New(service shortener.Service, BaseURL string) *Handler {
	return &Handler{
		service: service,
		BaseURL: BaseURL,
	}
}
