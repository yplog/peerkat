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
	"strings"
	"syscall"
)

type Node struct {
	relayAddrStr string
	peerAddrStr  string

	Host   host.Host
	ctx    context.Context
	cancel context.CancelFunc

	stopCh chan os.Signal
}

func New(relayAddrStr string, peerAddrStr string) *Node {
	ctx, cancel := context.WithCancel(context.Background())

	node, err := libp2p.New()
	if err != nil {
		log.Fatalf("Failed to create node: %v", err)
	}

	log.Default().Println("Node ID:", node.ID().String())
	log.Default().Println("Node address:", node.Addrs()[0].String())
	log.Default().Println("Node multi address:", node.Addrs()[0].String()+"/p2p/"+node.ID().String())

	return &Node{
		relayAddrStr: relayAddrStr,
		peerAddrStr:  peerAddrStr,
		ctx:          ctx,
		cancel:       cancel,
		Host:         node,
		stopCh:       make(chan os.Signal, 1),
	}
}

func (n *Node) ConnectRelay() {
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
}

func (n *Node) StartChat() {
	log.Default().Println("Connected to relay!")

	if n.peerAddrStr == "" {
		n.startPeer()
	} else {
		rw, err := startPeerAndConnect(n.Host, n.peerAddrStr)
		if err != nil {
			log.Println(err)
			return
		}

		go writeData(rw, n)
		go readData(rw, n)
	}

	signal.Notify(n.stopCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-n.ctx.Done():
	case <-n.stopCh:
		n.Stop()
	}
}

func (n *Node) Stop() {
	n.cancel()

	err := n.Host.Close()
	if err != nil {
		log.Fatalf("Failed to close node: %v", err)
	}
}

func readData(rw *bufio.ReadWriter, n *Node) {
	for {
		select {
		case <-n.ctx.Done():
			return
		default:
			str, _ := rw.ReadString('\n')

			if str == "" {
				return
			}
			if str != "\n" {
				fmt.Printf("\x1b[32m%s\x1b[0m> ", str)
			}

			if isExit(str) {
				n.Stop()
			}
		}
	}
}

func writeData(rw *bufio.ReadWriter, n *Node) {
	stdReader := bufio.NewReader(os.Stdin)

	for {
		select {
		case <-n.ctx.Done():
			return
		default:
			fmt.Print("> ")
			sendData, err := stdReader.ReadString('\n')
			if err != nil {
				log.Println(err)
				return
			}

			_, err = rw.WriteString(fmt.Sprintf("%s\n", sendData))
			if err != nil {
				log.Fatalf("failed to write to stream: %v", err)
			}
			err = rw.Flush()
			if err != nil {
				log.Fatalf("failed to flush writer: %v", err)
			}

			if isExit(sendData) {
				n.Stop()
			}
		}
	}
}

func isExit(str string) bool {
	return strings.TrimSpace(str) == "exit"
}

func (n *Node) startPeer() {
	n.Host.SetStreamHandler("/chat/1.0.0", n.handleStream)

	var port string
	for _, la := range n.Host.Network().ListenAddresses() {
		if p, err := la.ValueForProtocol(multiaddr.P_TCP); err == nil {
			port = p
			break
		}
	}

	if port == "" {
		log.Println("was not able to find actual local port")
		return
	}

	log.Default().Println("Node Address:", n.Host.Addrs()[0].String()+"/p2p/"+n.Host.ID().String())

	log.Println("Waiting for incoming connection")
}

func (n *Node) handleStream(s network.Stream) {
	log.Println("Got a new stream!")

	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go readData(rw, n)
	go writeData(rw, n)
}

func startPeerAndConnect(h host.Host, destination string) (*bufio.ReadWriter, error) {
	log.Println("This node's multi addresses:")
	for _, la := range h.Addrs() {
		log.Printf(" - %v\n", la)
	}
	log.Println()

	multiAddr, err := multiaddr.NewMultiaddr(destination)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	info, err := peer.AddrInfoFromP2pAddr(multiAddr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	h.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)

	s, err := h.NewStream(context.Background(), info.ID, "/chat/1.0.0")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Established connection to destination")

	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	return rw, nil
}
