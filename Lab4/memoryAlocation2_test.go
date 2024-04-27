package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingArray1MBOnStack(t *testing.T) {
	var memStatsBefore, memStatsAfter runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)

	var array [1 << 20]byte
	for i := range array {
		array[i] = 0
	}

	runtime.ReadMemStats(&memStatsAfter)
	allocatedMemoryDiff := int(memStatsAfter.TotalAlloc - memStatsBefore.TotalAlloc)

	t.Logf("Memory allocated before run: %d bytes\n", memStatsAfter.TotalAlloc)
	t.Logf("Memory allocated after run: %d bytes\n", memStatsBefore.TotalAlloc)
	t.Logf("Dif: %d bytes\n", allocatedMemoryDiff)

	assert.Equal(t, allocatedMemoryDiff, 0, "Expected to have no heap allocations")

}

func TestCreatingArray1MBOnHeap(t *testing.T) {
	var memStatsBefore, memStatsAfter runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)

	array := make([]byte, 1<<20)
	for i := range array {
		array[i] = 0
	}

	runtime.ReadMemStats(&memStatsAfter)
	allocatedMemoryDiff := int(memStatsAfter.TotalAlloc - memStatsBefore.TotalAlloc)

	t.Logf("Memory allocated before run: %d bytes\n", memStatsAfter.TotalAlloc)
	t.Logf("Memory allocated after run: %d bytes\n", memStatsBefore.TotalAlloc)
	t.Logf("Dif: %d bytes\n", allocatedMemoryDiff)

	assert.GreaterOrEqual(t, allocatedMemoryDiff, 1<<20, "Expected to allocate at least 1MB")

}
