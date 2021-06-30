package apiserver

import "http-rest-api/internal/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level""`
	Store *store.Config
}

func NewConfig() *Config {
	return &Config{
		Store: store.NewConfig()
	}
}
