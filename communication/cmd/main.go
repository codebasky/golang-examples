package main

import (
	"fmt"
	"golang-examples/communication/client"
	"golang-examples/communication/server"
	"os"
)

func main() {
	fmt.Println("welcome back basky")
	if len(os.Args) < 2 {
		fmt.Println("pass the server/client type argument")
		return
	}
	if os.Args[1] == "server" {
		server.Server("localhost:8065")
	} else {
		client.Client("localhost:8065")
	}
}
