// ref https://blog.golang.org/defer-panic-and-recover
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("happen :", r)
		}
	}()
	panic("oh my god!!")
}
