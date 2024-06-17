package main

import (
	"flag"
	"github.com/yplog/peerkat/pkg/node"
	"log"
)

func main() {
	relayAddr := flag.String("relay", "", "Relay address")
	peerAddr := flag.String("peer", "", "Peer address (optional)")
	mode := flag.String("mode", "", "Mode (chat, file-transfer)")

	flag.Parse()

	if *relayAddr == "" || *mode == "" {
		log.Fatalln("Usage: peerkat -relay <relay-address> -mode <mode> [-peer <peer-address>]")
	}

	if *mode != "chat" && *mode != "file-transfer" {
		log.Fatalln("Mode must be either 'chat' or 'file-transfer'")
	}

	log.Println("Relay address: ", *relayAddr)
	log.Println("mode: ", *mode)

	if *peerAddr != "" {
		log.Println("Peer address: ", *peerAddr)
	}
	log.Println("Starting peerkat node...")

	peerNode := node.New(*relayAddr, *peerAddr)

	peerNode.ConnectRelay()

	if *mode == "chat" {
		peerNode.StartChat()
	} else if *mode == "file-transfer" {
		peerNode.StartFileTransfer()
	} else {
		log.Fatalln("Invalid mode")
	}
}
