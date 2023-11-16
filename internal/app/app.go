package app

import (
	"github.com/akmyrzza/yandex-url/internal/config"
	"github.com/akmyrzza/yandex-url/internal/handler"
	"github.com/akmyrzza/yandex-url/internal/router"
	"github.com/akmyrzza/yandex-url/internal/shortener"
	"github.com/akmyrzza/yandex-url/internal/storage"
	"net/http"
)

func Run(cfg *config.Config) error {

	strg := storage.New()
	srvs := shortener.New(*strg)
	hndlr := handler.New(*srvs, cfg.BaseURL)

	srv := &http.Server{
		Handler: router.InitRouter(hndlr),
		Addr:    cfg.ServerAddr,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

	return nil
}
