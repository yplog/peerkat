package node

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/multiformats/go-multiaddr"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Node struct {
	relayAddrStr string

	host   host.Host
	ctx    context.Context
	stopCh chan os.Signal
}

func New(relayAddrStr string) *Node {
	ctx := context.Background()

	node, err := libp2p.New()
	if err != nil {
		log.Fatalf("Failed to create node: %v", err)
	}

	fmt.Println("Node ID:", node.ID().String())

	return &Node{
		relayAddrStr: relayAddrStr,
		ctx:          ctx,
		host:         node,
		stopCh:       make(chan os.Signal, 1),
	}
}

func (n *Node) Start() {
	relayAddr, err := multiaddr.NewMultiaddr(n.relayAddrStr)
	if err != nil {
		log.Fatalf("failed to parse relay address: %v", err)
	}

	relayInfo, err := peer.AddrInfoFromP2pAddr(relayAddr)
	if err != nil {
		log.Fatalf("failed to parse relay address: %v", err)
	}

	n.host.Peerstore().AddAddrs(relayInfo.ID, relayInfo.Addrs, peerstore.PermanentAddrTTL)
	if err := n.host.Connect(n.ctx, *relayInfo); err != nil {
		log.Fatalf("failed to connect to relay: %v", err)
	}

	log.Default().Println("Connected to the relay")

	signal.Notify(n.stopCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-n.stopCh:
		n.Stop()
	}
}

func (n *Node) Stop() {
	err := n.host.Close()
	if err != nil {
		log.Fatalf("Failed to close node: %v", err)
	}
}
