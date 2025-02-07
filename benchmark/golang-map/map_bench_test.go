package gomap

import (
	"strconv"
	"sync"
	"testing"

	cmap "github.com/orcaman/concurrent-map/v2"
)

const (
	mapsize   = 5000
	rwcnt     = mapsize
	allread   = rwcnt
	moreread  = 100
	halfread  = 2
	morewrite = 100
	allwrite  = rwcnt
)

func initMap() map[string]interface{} {
	m := make(map[string]interface{})
	for i := range mapsize {
		m[strconv.Itoa(i)] = 0
	}
	return m
}

func initSyncMap() *sync.Map {
	m := &sync.Map{}
	for i := range mapsize {
		m.Store(strconv.Itoa(i), 0)
	}
	return m
}

func initCMap() *cmap.ConcurrentMap[string, interface{}] {
	m := cmap.New[interface{}]()
	for i := range mapsize {
		m.Set(strconv.Itoa(i), 0)
	}
	return &m
}

func BenchmarkLockMapAllRead(b *testing.B) {
	var lk sync.RWMutex
	m := initMap()
	var wg sync.WaitGroup
	wg.Add(b.N)

	b.ResetTimer()
	for range b.N {
		go rwLockMap(m, &lk, &wg, allread, false)
	}
	wg.Wait()
}

func BenchmarkSyncMapAllRead(b *testing.B) {
	m := initSyncMap()
	b.ResetTimer()
	var wg sync.WaitGroup
	wg.Add(b.N)

	for range b.N {
		go rwSyncMap(m, &wg, allread, false)
	}
	wg.Wait()
}

func BenchmarkCMapAllRead(b *testing.B) {
	m := initCMap()
	b.ResetTimer()
	var wg sync.WaitGroup
	wg.Add(b.N)

	for range b.N {
		go rwCMap(m, &wg, allread, false)
	}
	wg.Wait()
}

func BenchmarkLockMapMoreRead(b *testing.B) {
	var lk sync.RWMutex
	m := initMap()
	var wg sync.WaitGroup
	wg.Add(b.N)

	b.ResetTimer()
	for range b.N {
		go rwLockMap(m, &lk, &wg, moreread, false)
	}
	wg.Wait()
}

func BenchmarkSyncMapMoreRead(b *testing.B) {
	m := initSyncMap()
	b.ResetTimer()
	var wg sync.WaitGroup
	wg.Add(b.N)

	for range b.N {
		go rwSyncMap(m, &wg, moreread, false)
	}
	wg.Wait()
}

func BenchmarkCMapMoreRead(b *testing.B) {
	m := initCMap()
	b.ResetTimer()
	var wg sync.WaitGroup
	wg.Add(b.N)

	for range b.N {
		go rwCMap(m, &wg, moreread, false)
	}
	wg.Wait()
}

func BenchmarkLockMapHalfRead(b *testing.B) {
	var lk sync.RWMutex
	m := initMap()
	var wg sync.WaitGroup
	wg.Add(b.N)

	b.ResetTimer()
	for range b.N {
		go rwLockMap(m, &lk, &wg, halfread, false)
	}
	wg.Wait()
}

func BenchmarkSyncMapHalfRead(b *testing.B) {
	m := initSyncMap()
	b.ResetTimer()
	var wg sync.WaitGroup
	wg.Add(b.N)

	for range b.N {
		go rwSyncMap(m, &wg, halfread, false)
	}
	wg.Wait()
}

func BenchmarkCMapHalfRead(b *testing.B) {
	m := initCMap()
	b.ResetTimer()
	var wg sync.WaitGroup
	wg.Add(b.N)

	for range b.N {
		go rwCMap(m, &wg, halfread, false)
	}
	wg.Wait()
}

func BenchmarkLockMapMoreWrite(b *testing.B) {
	var lk sync.RWMutex
	m := initMap()
	var wg sync.WaitGroup
	wg.Add(b.N)

	b.ResetTimer()
	for range b.N {
		go rwLockMap(m, &lk, &wg, morewrite, true)
	}
	wg.Wait()
}

func BenchmarkSyncMapMoreWrite(b *testing.B) {
	m := initSyncMap()
	b.ResetTimer()
	var wg sync.WaitGroup
	wg.Add(b.N)

	for range b.N {
		go rwSyncMap(m, &wg, morewrite, true)
	}
	wg.Wait()
}

func BenchmarkCMapMoreWrite(b *testing.B) {
	m := initCMap()
	b.ResetTimer()
	var wg sync.WaitGroup
	wg.Add(b.N)

	for range b.N {
		go rwCMap(m, &wg, morewrite, true)
	}
	wg.Wait()
}

func BenchmarkLockMapAllWrite(b *testing.B) {
	var lk sync.RWMutex
	m := initMap()
	var wg sync.WaitGroup
	wg.Add(b.N)

	b.ResetTimer()
	for range b.N {
		go rwLockMap(m, &lk, &wg, allwrite, true)
	}
	wg.Wait()
}

func BenchmarkSyncMapAllWrite(b *testing.B) {
	m := initSyncMap()
	b.ResetTimer()
	var wg sync.WaitGroup
	wg.Add(b.N)

	for range b.N {
		go rwSyncMap(m, &wg, allwrite, true)
	}
	wg.Wait()
}

func BenchmarkCMapAllWrite(b *testing.B) {
	m := initCMap()
	b.ResetTimer()
	var wg sync.WaitGroup
	wg.Add(b.N)

	for range b.N {
		go rwCMap(m, &wg, allwrite, true)
	}
	wg.Wait()
}

func rwLockMap(m map[string]interface{}, lk *sync.RWMutex, wg *sync.WaitGroup, rcnt int, flag bool) {
	defer wg.Done()
	for i := range rwcnt {
		if xor(i%rcnt == 0, flag) {
			lk.Lock()
			m[strconv.Itoa(mapsize+i%mapsize)] = 0
			lk.Unlock()
		} else {
			lk.RLock()
			_ = m[strconv.Itoa(i%mapsize)]
			lk.RUnlock()
		}
	}
}

func rwSyncMap(m *sync.Map, wg *sync.WaitGroup, rcnt int, flag bool) {
	defer wg.Done()
	for i := range rwcnt {
		if xor(i%rcnt == 0, flag) {
			m.Store(strconv.Itoa(mapsize+i%mapsize), 0)
		} else {
			_, _ = m.Load(strconv.Itoa(i % mapsize))
		}
	}
}

func rwCMap(m *cmap.ConcurrentMap[string, interface{}], wg *sync.WaitGroup, rcnt int, flag bool) {
	defer wg.Done()
	for i := range rwcnt {
		if xor(i%rcnt == 0, flag) {
			m.Set(strconv.Itoa(mapsize+i%mapsize), 0)
		} else {
			_, _ = m.Get(strconv.Itoa(i % mapsize))
		}
	}
}

func xor(a bool, b bool) bool {
	return (a || b) && !(a && b)
}
