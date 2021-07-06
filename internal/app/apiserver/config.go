package apiserver

import "github.com/1Bitcoin/http-rest-api/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level""`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{}
}
