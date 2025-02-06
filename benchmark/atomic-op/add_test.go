package atomicop

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	var sum uint64
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1e7; j++ {
			sum++
		}
	}
}

func BenchmarkAtomicAdd(b *testing.B) {
	var sum uint64
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1e7; j++ {
			atomic.AddUint64(&sum, 1)
		}
	}
}

func BenchmarkLockAdd(b *testing.B) {
	var sum uint64
	var lock sync.Mutex
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1e7; j++ {
			lock.Lock()
			sum++
			lock.Unlock()
		}
	}
}
