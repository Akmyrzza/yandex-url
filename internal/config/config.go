package config

import "flag"

type Config struct {
	SERVER_ADDR string
	BASE_URL    string
}

func InitConfig() (*Config, error) {
	cfg := new(Config)

	flag.StringVar(&cfg.SERVER_ADDR, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&cfg.BASE_URL, "b", "http://localhost:8080", "base address of the resulting shortened URL")

	flag.Parse()

	return cfg, nil
}
