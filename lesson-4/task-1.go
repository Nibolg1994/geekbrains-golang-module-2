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
	wg.Add(1000)
	for i := 1; i <= 1000; i++ {
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
