package main

import (
	"github.com/BurntSushi/toml"
	"http-rest-api/internal/app/apiserver"
	"log"
)

func main() {
	configPath := "configs/apiserver.toml"

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
