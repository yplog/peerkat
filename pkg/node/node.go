package node

import (
	"bufio"
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
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
	peerAddrStr  string

	Host   host.Host
	ctx    context.Context
	stopCh chan os.Signal
}

func New(relayAddrStr string, peerAddrStr string) *Node {
	ctx := context.Background()

	node, err := libp2p.New()
	if err != nil {
		log.Fatalf("Failed to create node: %v", err)
	}

	log.Default().Println("Node ID:", node.ID().String())
	log.Default().Println("Node address:", node.Addrs()[0].String())

	return &Node{
		relayAddrStr: relayAddrStr,
		peerAddrStr:  peerAddrStr,
		ctx:          ctx,
		Host:         node,
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

	n.Host.Peerstore().AddAddrs(relayInfo.ID, relayInfo.Addrs, peerstore.PermanentAddrTTL)
	if err := n.Host.Connect(n.ctx, *relayInfo); err != nil {
		log.Fatalf("failed to connect to relay: %v", err)
	}

	log.Default().Println("Connected to relay!")

	if n.peerAddrStr != "" {
		peerAddr, err := multiaddr.NewMultiaddr(n.peerAddrStr)
		if err != nil {
			log.Fatalf("failed to parse peer address: %v", err)
		}

		peerInfo, err := peer.AddrInfoFromP2pAddr(peerAddr)
		if err != nil {
			log.Fatalf("failed to parse peer address: %v", err)
		}

		n.Host.Peerstore().AddAddrs(peerInfo.ID, peerInfo.Addrs, peerstore.PermanentAddrTTL)
		if err := n.Host.Connect(n.ctx, *peerInfo); err != nil {
			log.Fatalf("failed to connect to peer: %v", err)
		}

		fmt.Println("Connecting to Peer A...")
		err = n.Host.Connect(n.ctx, *peerInfo)
		if err != nil {
			log.Fatal(err)
		}

		log.Default().Println("Connected Peer ID:", peerInfo.ID.String())

		stream, err := n.Host.NewStream(n.ctx, peerInfo.ID, "/chat/1.0.0")
		if err != nil {
			log.Fatal(err)
		}

		_, err = stream.Write([]byte("Hello from Peer B\n"))
		if err != nil {
			log.Fatalf("Failed to write to stream: %v", err)
		}

		err = stream.Close()
		if err != nil {
			log.Fatalf("Failed to close stream: %v", err)
		}
	} else {
		n.Host.SetStreamHandler("/chat/1.0.0", handleStream)
	}

	signal.Notify(n.stopCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-n.stopCh:
		n.Stop()
	}
}

func (n *Node) Stop() {
	err := n.Host.Close()
	if err != nil {
		log.Fatalf("Failed to close node: %v", err)
	}
}

func handleStream(stream network.Stream) {
	log.Default().Printf("Stream ID: %s\n", stream.ID())
	r := bufio.NewReader(stream)
	str, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	log.Default().Printf("Received: %s", str)

	err = stream.Close()
	if err != nil {
		log.Fatalf("Failed to close stream: %v", err)
	}

	log.Default().Println("Stream closed")
}
