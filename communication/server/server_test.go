package server_test

import (
	"fmt"
	"golang-examples/communication/server"
	"net"
	"testing"

	"go.uber.org/goleak"
)

func TestA(t *testing.T) {
	defer goleak.VerifyNone(t)
	go server.Server("localhost:8065")
	c, err := net.Dial("tcp", "localhost:8065")
	if err != nil {
		fmt.Printf("Client dialup error %s \n", err)
		return
	}
	defer c.Close()
	_, err = c.Write([]byte("client test"))
	if err != nil {
		t.Error("write failed")
	}
}
