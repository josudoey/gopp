package main

import "fmt"

func describe(args ...interface{}) {
	fmt.Printf("args = (%v, %T)\n", args, args)
	for i, v := range args {
		fmt.Printf("args = [%d] (%v, %T)\n", i, v, v)
	}
}

func main() {
	describe(nil, true, 1, 1.0, "a", describe, make(chan []interface{}, 1))
}
