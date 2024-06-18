package filetransfer

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type NodeInterface interface {
	Stop()
	Done() <-chan struct{}
}

func ReadFileData(rw *bufio.ReadWriter, n NodeInterface) {
	for {
		select {
		case <-n.Done():
			return
		default:
			str, _ := rw.ReadString('\n')

			if str == "" {
				return
			}

			if strings.HasPrefix(str, "/file ") {
				parts := strings.Fields(str)
				if len(parts) < 3 {
					fmt.Println("Invalid file message")
					return
				}
				filename := parts[1]
				encoded := parts[2]
				data, err := base64.StdEncoding.DecodeString(encoded)
				if err != nil {
					fmt.Printf("Failed to decode file data: %v\n", err)
					return
				}
				err = os.WriteFile(filename, data, 0644)
				if err != nil {
					fmt.Printf("Failed to write file: %v\n", err)
					return
				}
				fmt.Printf("Received file: %s\n", filename)
				return
			}
		}
	}
}

func WriteFileData(rw *bufio.ReadWriter, n NodeInterface) {
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

			if strings.HasPrefix(sendData, "/sendfile ") {
				parts := strings.Fields(sendData)
				if len(parts) != 2 {
					fmt.Println("Usage: /sendfile <filepath>")
					return
				}
				fp := parts[1]
				data, err := os.ReadFile(fp)
				if err != nil {
					fmt.Printf("Failed to read file: %v\n", err)
					return
				}
				encoded := base64.StdEncoding.EncodeToString(data)
				_, err = rw.WriteString(fmt.Sprintf("/file %s %s\n", filepath.Base(fp), encoded))
				if err != nil {
					log.Fatalf("failed to write to stream: %v", err)
				}
				err = rw.Flush()
				if err != nil {
					log.Fatalf("failed to flush writer: %v", err)
				}
			}
		}
	}
}
