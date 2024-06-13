package main

import (
	"flag"
	"fmt"
	"github.com/yplog/peerkat/pkg/node"
	"os"
)

func main() {
	relayAddr := flag.String("relay", "", "Relay address")
	peerAddr := flag.String("peer", "", "Peer address (optional)")
	mode := flag.String("mode", "", "Mode (chat, file-transfer)")

	flag.Parse()

	if *relayAddr == "" || *mode == "" {
		fmt.Println("Please provide relay address and mode as flags")
		os.Exit(1)
	}

	if *mode != "chat" && *mode != "file-transfer" {
		fmt.Println("Mode must be either 'chat' or 'file-transfer'")
		os.Exit(1)
	}

	fmt.Print("Relay address: ", *relayAddr, "\n")
	fmt.Print("Mode: ", *mode, "\n")

	if *peerAddr != "" {
		fmt.Print("Peer address: ", *peerAddr, "\n")
	}
	fmt.Println("Starting peerkat node...")

	fmt.Print("Relay address: ", relayAddr, "\n")
	fmt.Println("Starting peerkat node...")

	peerNode := node.New(*relayAddr, *peerAddr)

	peerNode.ConnectRelay()

	if *mode == "chat" {
		peerNode.StartChat()
	} else {
		peerNode.StartFileTransfer()
	}
}
