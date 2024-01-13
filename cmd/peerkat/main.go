package main

import (
	"fmt"
	"log"

	"github.com/yplog/peerkat/pkg/cli"
	"github.com/yplog/peerkat/pkg/config"
)

func setup() {
	configFileExists := config.Exists()
	if !configFileExists {
		log.Println("Config file not found, generating...")
		config.Generate()
		fmt.Println("")
	}
}

func main() {
	setup()
	cli.Execute()
}
