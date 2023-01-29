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
	ureader := bufio.NewReader(os.Stdin)
	sreader := bufio.NewReader(c)
	for {
		fmt.Print(">> ")
		data, err := ureader.ReadString('\n')
		if err != nil {
			fmt.Println("Reading user input error")
			return
		}
		fmt.Fprint(c, data)
		message, _ := sreader.ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
