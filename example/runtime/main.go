//ref https://golang.org/pkg/runtime/#MemStats
package main

import (
	"fmt"
	"runtime"
)

func main() {
	var mem runtime.MemStats
	fmt.Printf("%-15v %-15v %-15v %-15v\n", "Alloc", "TotalAlloc", "HeapAlloc", "HeapSys")
	runtime.ReadMemStats(&mem)
	fmt.Printf("%-15v %-15v %-15v %-15v\n", mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
}
