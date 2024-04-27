package main

import (
	"testing"
)

func BenchmarkStackAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = stackIt()
	}
}

func stackIt() int {
	y := 2
	return y * 2
}

// 1000000000               0.2636 ns/op          0 B/op          0 allocs/op

func BenchmarkHeapAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = heapIt()
	}
}

func heapIt() []int {
	slice := make([]int, 1)
	return append(slice, 42)
}

//42108364                26.46 ns/op           16 B/op          1 allocs/op <- heap
