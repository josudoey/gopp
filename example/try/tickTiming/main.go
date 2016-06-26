package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now().Unix()
	fmt.Printf("now: %-20v\n", 0)
	t1 := time.Tick(1 * time.Second)
	t3 := time.Tick(3 * time.Second)
	for i := 0; i < 3; i++ {
		select {
		case n3 := <-t3:
			fmt.Printf("t3: %-20v\n", n3.Unix()-t0)
			n1 := <-t1
			fmt.Printf("t1: %-20v\n", n1.Unix()-t0)
		}
	}
	// Output:
	// now: 0
	// t3: 3
	// t1: 1
	// t3: 6
	// t1: 4
	// t3: 9
	// t1: 7
}
