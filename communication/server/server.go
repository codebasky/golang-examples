package server

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var count = 0

func Server(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Printf("Accept error on server %s", err)
			return err
		}
		count++
		go handleConnection(c)
	}

}

func handleConnection(c net.Conn) {
	defer func() {
		count--
		c.Close()
	}()
	reader := bufio.NewScanner(c)
	for reader.Scan() {
		data := reader.Text()
		fmt.Println("rcv->" + data)
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
		c.Write([]byte(strconv.Itoa(count) + "\n"))
	}
	if err := reader.Err(); err != nil {
		fmt.Printf("Reading input err: %s\n", err)
	}
}
