package gomap

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	cmap "github.com/orcaman/concurrent-map/v2"
)

func BenchmarkMapWithLock(b *testing.B) {
	for _, item := range testcase {
		b.Run(item.name, func(b *testing.B) {
			rand.Seed(time.Now().UnixNano())
			var lock sync.RWMutex
			m := make(map[string]int)
			for i := 0; i < int(item.msize); i++ {
				m[strconv.Itoa(i)] = i
			}
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				wg := sync.WaitGroup{}
				// read
				for i := 0; i < int(item.read); i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						for j := 0; j < int(item.times); j++ {
							tmp := rand.Int() % int(item.msize)
							lock.RLock()
							_ = m[strconv.Itoa(tmp)]
							lock.RUnlock()
						}
					}()
				}
				// write
				for i := 0; i < int(item.write); i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						for j := 0; j < int(item.times); j++ {
							tmp := rand.Int() % int(item.msize)
							lock.Lock()
							m[strconv.Itoa(tmp)] = tmp + 1
							lock.Unlock()
						}
					}()
				}
				wg.Wait()
			}
		})
	}
}

func BenchmarkSyncMap(b *testing.B) {
	for _, item := range testcase {
		b.Run(item.name, func(b *testing.B) {
			rand.Seed(time.Now().UnixNano())
			m := sync.Map{}
			for i := 0; i < int(item.msize); i++ {
				m.Store(strconv.Itoa(i), i)
			}
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				wg := sync.WaitGroup{}
				// read
				for i := 0; i < int(item.read); i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						for j := 0; j < int(item.times); j++ {
							tmp := rand.Int() % int(item.msize)
							_, _ = m.Load(strconv.Itoa(tmp))
						}
					}()
				}
				// write
				for i := 0; i < int(item.write); i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						for j := 0; j < int(item.times); j++ {
							tmp := rand.Int() % int(item.msize)
							m.Store(strconv.Itoa(tmp), tmp+1)
						}
					}()
				}
				wg.Wait()
			}
		})
	}
}

func BenchmarkCurMap(b *testing.B) {
	for _, item := range testcase {
		b.Run(item.name, func(b *testing.B) {
			rand.Seed(time.Now().UnixNano())
			m := cmap.New[int]()
			for i := 0; i < int(item.msize); i++ {
				m.Set(strconv.Itoa(i), i)
			}
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				wg := sync.WaitGroup{}
				// read
				for i := 0; i < int(item.read); i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						for j := 0; j < int(item.times); j++ {
							tmp := rand.Int() % int(item.msize)
							_, _ = m.Get(strconv.Itoa(tmp))
						}
					}()
				}
				// write
				for i := 0; i < int(item.write); i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						for j := 0; j < int(item.times); j++ {
							tmp := rand.Int() % int(item.msize)
							m.Set(strconv.Itoa(tmp), tmp+1)
						}
					}()
				}
				wg.Wait()
			}
		})
	}
}
