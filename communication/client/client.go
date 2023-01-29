package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Client(addr string) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("Client dialup error %s \n", err)
		return
	}
	defer c.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Reading user input error")
			return
		}
		fmt.Fprint(c, data)
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
