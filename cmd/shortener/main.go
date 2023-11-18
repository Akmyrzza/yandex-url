package main

import (
	"github.com/akmyrzza/yandex-url/internal/app"
	"github.com/akmyrzza/yandex-url/internal/config"
	"log"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("error, initializing config: %d", err)
	}

	if err := app.Run(cfg); err != nil {
		log.Fatalf("error, running app: %d", err)
	}
}
