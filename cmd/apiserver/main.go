package main

import (
	"github.com/1Bitcoin/http-rest-api/internal/app/apiserver"
	"github.com/BurntSushi/toml"
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
