package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkWithoutGC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Generate some objects
		objects := make([]interface{}, 1<<10)
		for i := range objects {
			objects[i] = new(interface{})
		}
	}
}

func BenchmarkWithGC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Generate some objects
		objects := make([]interface{}, 1<<10)
		for i := range objects {
			objects[i] = new(interface{})
		}
		// Force garbage collection
		runtime.GC()
	}
}

func TestGarbageCollection(t *testing.T) {
	var withoutGCMem, withGCMem uint64

	// Benchmark without forcing GC
	withoutGCMem = testing.Benchmark(BenchmarkWithoutGC).MemBytes

	// Benchmark with forcing GC after object creation
	withGCMem = testing.Benchmark(BenchmarkWithGC).MemBytes

	t.Logf("Memory used without GC: %d bytes\n", withoutGCMem)
	t.Logf("Memory used with GC: %d bytes\n", withGCMem)

	assert.Greater(t, withoutGCMem, withGCMem, "Expected memory usage to be lower with GC")
}
