package app

import (
	"github.com/akmyrzza/yandex-url/internal/config"
	"github.com/akmyrzza/yandex-url/internal/handler"
	"github.com/akmyrzza/yandex-url/internal/router"
	"github.com/akmyrzza/yandex-url/internal/service"
	"github.com/akmyrzza/yandex-url/internal/storage"
)

func Run(cfg *config.Config) error {

	strg := storage.New()
	srvs := service.New(*strg)
	hndlr := handler.New(*srvs, cfg.BaseURL)
	srv := service.NewShortener(router.InitRouter(hndlr), cfg.ServerAddr)
	err := srv.Start()
	if err != nil {
		panic(err)
	}

	return nil
}
