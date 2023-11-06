package main

import "github.com/akmyrzza/yandex-url/internal/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
