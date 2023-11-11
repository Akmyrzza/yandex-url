package handler

import "github.com/akmyrzza/yandex-url/internal/service"

type Handler struct {
	service service.Service
	BaseURL string
}

func New(service service.Service, BaseURL string) *Handler {
	return &Handler{
		service: service,
		BaseURL: BaseURL,
	}
}
