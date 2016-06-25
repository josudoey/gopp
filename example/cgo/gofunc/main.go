package main

//ref https://golang.org/misc/cgo/life/main.go
//ref https://github.com/golang/go/wiki/cgo
//ref https://golang.org/cmd/cgo/

import (
	"./cfunc"
)

func main() {
	cfunc.Say("Joey", "hello")
}
