package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var (
		counter int32 = 0
	)
	var workers = make(chan struct{}, 10)
	var wg sync.WaitGroup
	var n int = 1000
	wg.Add(n)
	for i := 1; i <= n; i++ {
		workers <- struct{}{}
		go func() {
			atomic.AddInt32(&counter, 1)
			defer func() {
				<-workers
				wg.Done()
			}()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
