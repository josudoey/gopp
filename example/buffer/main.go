//ref https://golang.org/pkg/bytes/
package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func Example1() {
	var b bytes.Buffer // A Buffer needs no initialization.
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!\n")
	b.WriteTo(os.Stdout)
}

func Example2() {
	// A Buffer can turn a string or a []byte into an io.Reader.
	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec)
}

func main() {
	fmt.Println("=Example1=")
	Example1()
	fmt.Println("=Example2=")
	Example2()
}
