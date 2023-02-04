package server

import (
	"bufio"
	"fmt"
	"net"
)

type client struct {
	done chan bool
	msg  chan string
	cn   net.Conn
}

func AddClient(conn net.Conn, msgCh chan string, quit chan bool) (chan string, chan bool) {
	c := client{
		done: make(chan bool),
		msg:  make(chan string),
		cn:   conn,
	}

	go func(cn net.Conn) {
		sr := bufio.NewScanner(cn)
		for {
			if sr.Scan() && sr.Err() == nil {
				msg := sr.Text()
				fmt.Printf("recevied msg: %s\n", msg)
				msgCh <- msg
			}
			if sr.Err() != nil {
				break
			}
		}
		fmt.Println("client scan failed terminating it")
		c.done <- true
	}(c.cn)

	go func() {
		select {
		case <-quit:
			c.cn.Close()
		case <-c.done:
		}
	}()

	go c.send()
	return c.msg, c.done
}

func (c *client) send() {
	for m := range c.msg {
		fmt.Printf("sending msg to client: %s\n", m)
		fmt.Fprintf(c.cn, m+"\n")
	}
}
