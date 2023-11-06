package app

import (
	"github.com/akmyrzza/yandex-url/internal/config"
	"github.com/akmyrzza/yandex-url/internal/handler"
	"github.com/akmyrzza/yandex-url/internal/server"
)

func Run(cfg *config.Config) error {

	hndlr := handler.New(cfg.BaseURL)
	srv := server.New(hndlr.InitRouter(), cfg.ServerAddr)
	srv.Start()

	return nil
}
