package main

//ref https://godoc.org/golang.org/x/net/websocket#Handler
//ref https://github.com/golang-samples/websocket/blob/master/cli/src/server.go
//ref https://github.com/golang/net/blob/master/websocket/server.go
//ref http://stackoverflow.com/questions/19708330/serving-a-websocket-in-go
import (
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

func EchoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

// This example demonstrates a trivial echo server.
func main() {
	wsServer := websocket.Server{Handler: websocket.Handler(EchoHandler)}

	http.HandleFunc("/echo",
		func(w http.ResponseWriter, req *http.Request) {
			wsServer.ServeHTTP(w, req)
		})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
