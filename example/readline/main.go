package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bio := bufio.NewReader(os.Stdin)
	for {
		line, isPrefix, err := bio.ReadLine()
		if err != nil {
			break
		}
		if isPrefix {
			fmt.Printf("%s", string(line))
		} else {
			fmt.Println(string(line))
		}
	}
}
