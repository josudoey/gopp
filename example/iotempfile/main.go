package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// example ref
// https://github.com/kubernetes/kubectl/blob/master/pkg/cmd/get/customcolumn_flags_test.go#L35

func main() {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "example")
	if err != nil {
		log.Fatal(err)
	}
	defer func(tempFile *os.File) {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}(tmpFile)

	log.Printf("%v", tmpFile.Name())
	fmt.Fprintf(tmpFile, "hello world")
}
