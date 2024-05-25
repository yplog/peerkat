package main

import (
	"bufio"
	"fmt"
	"github.com/yplog/peerkat/pkg/node"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter relay address: ")
	relayAddr, _ := reader.ReadString('\n')
	relayAddr = relayAddr[:len(relayAddr)-1]

	fmt.Print("Relay address: ", relayAddr, "\n")
	fmt.Println("Starting peerkat node...")

	peerNode := node.New(relayAddr)
	peerNode.Start()
}
