package pool

import (
	"testing"
	"runtime"
)

type Object struct {
	a int
	b int
	c int
}

// measurements use M1 Max Pro arm64 platform
// produced with: go test -v -bench=. -benchtime=10s -benchmem

// This is a benchmark of the Pool implementation
// minimum space: 3072 btes. (4 lines * 32 entries * 24 bytes)
// measured space: 6416 bytes.
// measured time: 2833 ns/op.
// memory overhead: 208%
func BenchmarkPool(b *testing.B) {
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	lineSize := []uint16{64, 64, 64, 64}
	handles := new([256]Handle)
	ObjectQ := Object{2, 1, 0xFFFFFFFF}

	Pool, err := New(ObjectQ, 4, lineSize...)
	if err != nil {
		b.Log(err)
		return
		
	}
	for n := 0; n < b.N; n++ {
		// Add 128 pool objects across 4 cache lines
		var handleIdx int = 0
		for i := 0; i < 64; i++ {
			for j := 0; j < 4; j++ {
				var err error
				handles[handleIdx], err = Pool.Put(ObjectQ, uint16(j))
				if err != nil {
					b.Log(err)
					return
				}
				handleIdx++
			}
		}
		// Remove 255 pool objects - remove in reverse - reversal order may not matter
		for i := 255; i >= 0; i-- {
			Pool.Remove(handles[i])
		}
	}
	runtime.ReadMemStats(&m2)
	b.Logf("Total memory allocated: %v", m2.TotalAlloc - m1.TotalAlloc)
	b.Logf("Total number of mallocs: %v", m2.Mallocs - m1.Mallocs)
}

// This is a map based implementation of Pool
// minimum space: 3072 bytes. (4 lines * 32 entries * 24 bytes)
// measured space: 35192 bytes.
// measured time: 2346 ns/op.
// memory overhead: 1145%
func BenchmarkMap(b *testing.B) {
	type key struct {
		line  uint16
		entry uint16
	}

	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	m := make(map[key]Object)
	for n := 0; n < b.N; n++ {
		for i := 0; i < 64; i++ {
			for j := 0; j < 4; j++ {
				key := key{uint16(i), uint16(j)}
				infB := Object{i, j, 0xFFFFFFFF}
				m[key] = infB
			}
		}
	}

	runtime.ReadMemStats(&m2)
	b.Logf("Total memory allocated: %v", m2.TotalAlloc - m1.TotalAlloc)
	b.Logf("Total number of mallocs: %v", m2.Mallocs - m1.Mallocs)
}
