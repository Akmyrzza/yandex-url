package app

import (
	"github.com/akmyrzza/yandex-url/internal/config"
	"github.com/akmyrzza/yandex-url/internal/handler"
	"github.com/akmyrzza/yandex-url/internal/logger"
	"github.com/akmyrzza/yandex-url/internal/router"
	"github.com/akmyrzza/yandex-url/internal/shortener"
	"github.com/akmyrzza/yandex-url/internal/storage"
	"net/http"
)

func Run(cfg *config.Config) error {

	strg := storage.NewStorage()
	srvs := shortener.NewShortener(*strg)
	hndlr := handler.NewHandler(*srvs, cfg.BaseURL)
	lg := logger.NewLogger()

	srv := &http.Server{
		Handler: router.InitRouter(hndlr, lg),
		Addr:    cfg.ServerAddr,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
