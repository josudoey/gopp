package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.Tick(1 * time.Second)
	t := time.Now().Unix() + 3
	for now := range c {
		fmt.Printf("%v %v\n", now.UnixNano(), now)
		if now.Unix() > t {
			break
		}
	}
}
