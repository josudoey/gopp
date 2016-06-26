package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var mem runtime.MemStats
	fmt.Printf("%-10v %-15v %-15v %-15v %-15v %-15v\n", "time(ns)", "tick count", "Alloc", "TotalAlloc", "HeapAlloc", "HeapSys")
	runtime.ReadMemStats(&mem)
	i := 0
	t0 := time.Now().UnixNano()
	fmt.Printf("%-10v %-15v %-15v %-15v %-15v %-15v\n", time.Now().UnixNano()-t0, i, mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
	done := make(chan bool, 1)
	go func() {
		for ; i < 10000; i++ {
			_ = time.Tick(time.Second)
		}
		done <- true
	}()
	_ = <-time.Tick(time.Millisecond)
	runtime.ReadMemStats(&mem)
	fmt.Printf("%-10v %-15v %-15v %-15v %-15v %-15v\n", time.Now().UnixNano()-t0, i, mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
	_ = <-done
	runtime.ReadMemStats(&mem)
	fmt.Printf("%-10v %-15v %-15v %-15v %-15v %-15v\n", time.Now().UnixNano()-t0, i, mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
	runtime.GC()
	runtime.ReadMemStats(&mem)
	fmt.Printf("%-10v %-15v %-15v %-15v %-15v %-15v\n", time.Now().UnixNano()-t0, "After GC", mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
}
