package main

import (
	"fmt"
	"runtime"
)

type node struct {
	next *node
	val  int
}

func allocCircleLinkedList() *node {
	head := &node{}
	prev := head
	for i := 0; i < 1<<20; i++ {
		nextNode := &node{
			val: i,
		}
		prev.next = nextNode
		prev = nextNode
	}
	return prev
}

func showMemStats(memStats *runtime.MemStats) {
	runtime.ReadMemStats(memStats)
	fmt.Printf("HeapAlloc:%v NextGC:%v LastGC:%v\n", memStats.HeapAlloc, memStats.NextGC, memStats.LastGC)
}

func main() {
	var memStats runtime.MemStats
	showMemStats(&memStats)

	node := allocCircleLinkedList()
	showMemStats(&memStats)

	runtime.GC()
	fmt.Println("after gc")
	showMemStats(&memStats)

	fmt.Printf("node addr: %p\n", node)
	runtime.GC()
	fmt.Println("after gc")
	showMemStats(&memStats)

	// Output:
	// HeapAlloc:131712 NextGC:4473924 LastGC:0
	// HeapAlloc:6442580688 NextGC:10748656560 LastGC:1655118758376422000
	// after gc
	// HeapAlloc:6442583344 NextGC:12885164720 LastGC:1655118769143768000
	// node addr: 0x1400000c018
	// after gc
	// HeapAlloc:132376 NextGC:4194304 LastGC:1655118770251760000
}
