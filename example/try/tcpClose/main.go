package main

import (
	"fmt"
	"net"
	"time"
)

var done = make(chan bool, 1)

func WaitClient(ln net.Listener) {
	conn, err := ln.Accept()
	if err != nil {
		// handle error
	}
	ln.Close()
	fmt.Printf("lisnter close\n")
	time.Sleep(time.Second)
	fmt.Printf("client close\n")
	conn.Close()
}

func ClientConnet(addr net.Addr) {
	conn, err := net.Dial("tcp", addr.String())
	if err != nil {
		// handle error
	}
	b := make([]byte, 1024)
	for {
		n, err := conn.Read(b)
		fmt.Printf("client read %v %v\n", n, err)
		if n == 0 {
			break
		}
	}
	done <- true
}

func main() {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		// handle error
	}
	fmt.Printf("listen on %v\n", ln.Addr())
	go WaitClient(ln)
	go ClientConnet(ln.Addr())

	_ = <-done
}
