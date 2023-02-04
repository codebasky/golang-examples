package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run(port string) error {
	quit := make(chan bool)
	r := createRoom("Techy World")

	l, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		return err
	}

	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		l.Close()
		fmt.Println("terminating client connections")
		quit <- true
		if len(r.clients) > 0 {
			<-r.msgCh
		}
		os.Exit(0)
	}()

	for {
		con, err := l.Accept()
		if err != nil {
			return err
		}
		go r.handleConnection(con, quit)
	}
}
