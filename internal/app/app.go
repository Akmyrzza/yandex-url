package app

import (
	"github.com/akmyrzza/yandex-url/internal/handler"
	"github.com/akmyrzza/yandex-url/internal/server"
)

func Run() error {

	hndlr := handler.New()
	srv := server.New(hndlr.InitRouter())
	srv.Start()
	
	return nil
}
