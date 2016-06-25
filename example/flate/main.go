package main

import (
	"bytes"
	"compress/flate"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	var src, dst bytes.Buffer
	enc, err := flate.NewWriter(&dst, 9)
	dec := flate.NewReader(&dst)
	if err != nil {
		log.Fatal(err)
	}

	decBuf := make([]byte, 100)
	for i := 0; i < 10; i++ {
		src.Write([]byte{0})
		dst.Reset()
		enc.Write(src.Bytes())
		enc.Flush()
		dstHex := hex.EncodeToString(dst.Bytes())
		n, _ := dec.Read(decBuf)
		decHex := hex.EncodeToString(decBuf[:n])
		fmt.Printf("%-24v -> %-24v\n", decHex, dstHex)
	}
}
