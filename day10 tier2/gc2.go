package main

import (
	"fmt"
	"runtime"
)

func printStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: Alloc=%d KB, TotalAlloc=%d KB, NumGC=%d\n", label, m.Alloc/1024, m.TotalAlloc/1024, m.NumGC)
}

func main() {
	printStats("before")
	num := make([]int, 1000000)
	_ = num
	printStats("after alloc")
	runtime.GC()
	printStats("after GC")
}
