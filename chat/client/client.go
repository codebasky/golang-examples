package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Run(port string) error {
	c, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Printf("client dialup error: %s \n", err)
		return err
	}
	go func(con net.Conn) {
		sr := bufio.NewScanner(con)
		for {
			if sr.Scan() {
				fmt.Println(sr.Text())
			}
			if sr.Err() != nil {
				fmt.Println("terminating server scan for client")
				return
			}
		}
	}(c)
	ur := bufio.NewScanner(os.Stdin)
	name := "client>"
	for {
		if ur.Scan() && ur.Err() == nil {
			fmt.Fprintf(c, name+ur.Text()+"\n")
		} else {
			fmt.Printf("terminating client: %s\n", ur.Err())
			return nil
		}
	}
}
