package main

import (
	"fmt"
	"os"

	"github.com/codebasky/golang-examples/chat/client"
	"github.com/codebasky/golang-examples/chat/server"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("send args type (s/c) and port to program")
		return
	}
	switch os.Args[1] {
	case "s":
		server.Run(os.Args[2])
	case "c":
		client.Run(os.Args[2])
	default:
		fmt.Println("wrong type")
	}
}
