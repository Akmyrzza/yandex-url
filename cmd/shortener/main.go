package main


import (
	"github.com/akmyrzza/yandex-url/internal/app"
	"github.com/akmyrzza/yandex-url/internal/config"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	if err := app.Run(cfg); err != nil {
		panic(err)
	}

}
