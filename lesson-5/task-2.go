package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	counter int64
}

func (s *Counter) add(i int) {
	s.Lock()
	s.counter++
	defer s.Unlock()
}

func main() {
	var c Counter = Counter{counter: 0}
	var wg sync.WaitGroup
	var n int = 2
	wg.Add(n)
	go func() {
		c.add(1)
		wg.Done()
	}()
	go func() {
		c.add(1)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(c.counter)
}
