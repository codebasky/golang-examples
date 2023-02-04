package server

import (
	"fmt"
	"net"
	"sync"
)

type room struct {
	name    string
	clients map[chan<- string]struct{}
	msgCh   chan string
	lock    sync.RWMutex
}

func createRoom(n string) *room {
	r := room{
		name:    n,
		clients: make(map[chan<- string]struct{}),
		msgCh:   make(chan string),
	}

	go r.broadcast()
	return &r
}

func (r *room) handleConnection(c net.Conn, quit chan bool) {
	ch, done := AddClient(c, r.msgCh, quit)
	r.lock.Lock()
	r.clients[ch] = struct{}{}
	r.lock.Unlock()

	<-done
	r.removeClient(ch)
	if len(r.clients) == 0 && len(quit) > 0 {
		fmt.Println("closing the msg channel for the room")
		close(r.msgCh)
	}

}

func (r *room) broadcast() {
	for msg := range r.msgCh {
		fmt.Printf("broadcast msg: %s\n", msg)
		r.lock.RLock()
		for ch := range r.clients {
			go func(c chan<- string) {
				c <- msg
			}(ch)
		}
		r.lock.RUnlock()
	}
}

func (r *room) removeClient(ch chan string) {
	fmt.Println("removing client")
	r.lock.Lock()
	delete(r.clients, ch)
	r.lock.Unlock()
}
