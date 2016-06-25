package main

//ref https://github.com/golang-samples/websocket/blob/master/cli/src/client.go

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

var origin = "http://localhost/"
var url = "ws://echo.websocket.org/"

func main() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = ws.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("WebSocket close\n")
	}()

	message := []byte("hello, world!")
	n, err := ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send(%d): %s\n", n, message)

	var msg = make([]byte, 512)
	n, err = ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive(%d): %s\n", n, msg)
}
