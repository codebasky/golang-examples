package main

import (
	"chat/client"
	"chat/server"
	"fmt"
	"os"
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
