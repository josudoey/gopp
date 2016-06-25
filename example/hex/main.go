package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

func main() {
	str := "Hello World!"
	var buf bytes.Buffer
	buf.Write([]byte(str))
	hexStr := hex.EncodeToString(buf.Bytes())
	decStr, _ := hex.DecodeString(hexStr)
	fmt.Printf("%s -> %s", decStr, hexStr)
}
