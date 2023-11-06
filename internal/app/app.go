package app

import (
	"github.com/akmyrzza/yandex-url/internal/config"
	"github.com/akmyrzza/yandex-url/internal/handler"
	"github.com/akmyrzza/yandex-url/internal/server"
)

func Run(cfg *config.Config) error {

	hndlr := handler.New(cfg.BASE_URL)
	srv := server.New(hndlr.InitRouter(), cfg.SERVER_ADDR)
	srv.Start()

	return nil
}
