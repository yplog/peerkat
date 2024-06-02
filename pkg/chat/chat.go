package chat

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type NodeInterface interface {
	Stop()
	Done() <-chan struct{}
}

func ReadData(rw *bufio.ReadWriter, n NodeInterface) {
	for {
		select {
		case <-n.Done():
			return
		default:
			str, _ := rw.ReadString('\n')

			if isCommand(str) {
				commandHandler(str, n)
			}

			if str == "" {
				return
			}

			if str != "\n" {
				fmt.Printf("\x1b[32m%s\x1b[0m> ", str)
			}
		}
	}
}

func WriteData(rw *bufio.ReadWriter, n NodeInterface) {
	stdReader := bufio.NewReader(os.Stdin)

	for {
		select {
		case <-n.Done():
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

			if isCommand(sendData) {
				commandHandler(sendData, n)
			}
		}
	}
}

func isCommand(str string) bool {
	return strings.HasPrefix(str, "/")
}

func commandHandler(str string, n NodeInterface) {
	switch strings.TrimSpace(str) {
	case "/help":
		fmt.Println("Available commands:")
		fmt.Println("/help - show this message")
		fmt.Println("/exit - exit the chat")
	case "/exit":
		fmt.Println("Exiting chat...")
		n.Stop()
	default:
		fmt.Println("Unknown command. Type /help to see available commands")
	}
}
