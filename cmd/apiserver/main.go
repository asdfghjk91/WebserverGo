package main

import (
	"flag"
	"log"
	"webgo/internal/app/apiserver"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

// Init flag
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	//Init config
	config := apiserver.NewConfig()

	//Creat apiserver
	s := apiserver.New(config)

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	//Start apiserver
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
