package main

import (
	"math/rand"
	"sync"
	"testing"
)

func BenchmarkMutexRead(b *testing.B) {
	var (
		counter int
	)
	b.Run("BenchmarkMutexRead", func(b *testing.B) {
		var lock sync.Mutex
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(10)%10 == 0 {
					lock.Lock()
					counter++
					lock.Unlock()
				} else {
					lock.Lock()
					_ = counter
					lock.Unlock()
				}
			}
		})
	})
}

func BenchmarkMutexReadWrite(b *testing.B) {
	var (
		counter int
	)
	b.Run("BenchmarkMutexReadWrite", func(b *testing.B) {
		var lock sync.Mutex
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(10)%2 == 0 {
					lock.Lock()
					counter++
					lock.Unlock()
				} else {
					lock.Lock()
					_ = counter
					lock.Unlock()
				}
			}
		})
	})
}

func BenchmarkMutexWrite(b *testing.B) {
	var (
		counter int
	)
	b.Run("BenchmarkMutexWrite", func(b *testing.B) {
		var lock sync.Mutex
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(10)%10 == 0 {
					lock.Lock()
					counter++
					lock.Unlock()
				} else {
					lock.Lock()
					_ = counter
					lock.Unlock()
				}
			}
		})
	})
}

func BenchmarkRWMutexReadWrite(b *testing.B) {
	var (
		counter int
	)
	b.Run("BenchmarkRWMutexReadWrite", func(b *testing.B) {
		var lock sync.RWMutex
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(10)%2 == 0 {
					lock.Lock()
					counter++
					lock.Unlock()
				} else {
					lock.RLock()
					_ = counter
					lock.RUnlock()
				}
			}
		})
	})
}

func BenchmarkRwMutexWrite(b *testing.B) {
	var (
		counter int
	)
	b.Run("BenchmarkRwMutexWrite", func(b *testing.B) {
		var lock sync.RWMutex
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(10)%10 != 0 {
					lock.Lock()
					counter++
					lock.Unlock()
				} else {
					lock.RLock()
					_ = counter
					lock.RUnlock()
				}
			}
		})
	})
}

func BenchmarkRwMutexRead(b *testing.B) {
	var (
		counter int
	)
	b.Run("BenchmarkRwMutexRead", func(b *testing.B) {
		var lock sync.RWMutex
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(10)%10 == 0 {
					lock.Lock()
					counter++
					lock.Unlock()
				} else {
					lock.RLock()
					_ = counter
					lock.RUnlock()
				}
			}
		})
	})
}
