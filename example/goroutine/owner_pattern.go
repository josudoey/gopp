package main

import (
	"fmt"
	"time"
)

type Operation int

const (
	Read Operation = iota
	Increment
)

type Task struct {
	Op     Operation
	Result chan int
}

func counterManager(requests <-chan Task) {
	counter := 0 // safe counter
	for req := range requests {
		switch req.Op {
		case Read:
			req.Result <- counter // current counter
		case Increment:
			counter++
			fmt.Printf("Counter incremented to: %d\n", counter)
		}
	}
}

func main() {
	requests := make(chan Task) // Must unbuffered channel
	go counterManager(requests)

	// Goroutine 1
	go func() {
		respCh := make(chan int)
		requests <- Task{Op: Read, Result: respCh}
		val := <-respCh
		fmt.Printf("Goroutine 1: Initial counter value: %d\n", val)

		requests <- Task{Op: Increment}
		time.Sleep(50 * time.Millisecond) // like some task
		requests <- Task{Op: Increment}
	}()

	// Goroutine 2
	go func() {
		time.Sleep(100 * time.Millisecond) // like some task
		requests <- Task{Op: Increment}
		time.Sleep(50 * time.Millisecond)
		respCh := make(chan int)
		requests <- Task{Op: Read, Result: respCh}
		val := <-respCh
		fmt.Printf("Goroutine 2: Current counter value: %d\n", val)
	}()

	time.Sleep(2 * time.Second)
	close(requests)
}
