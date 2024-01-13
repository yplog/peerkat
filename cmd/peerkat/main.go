package main

import (
	"github.com/yplog/peerkat/pkg/cli"
	"github.com/yplog/peerkat/pkg/config"
)

func main() {
	config.Setup()
	cli.Execute()
}
