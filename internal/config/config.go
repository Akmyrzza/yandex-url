package config

import "flag"

type Config struct {
	ServerAddr string
	BaseURL    string
}

func InitConfig() (*Config, error) {
	cfg := new(Config)

	flag.StringVar(&cfg.ServerAddr, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&cfg.BaseURL, "b", "http://localhost:8080", "base address of the resulting shortened URL")

	flag.Parse()

	return cfg, nil
}
