package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
	mu    *sync.Mutex
}

func newCounter() counter {
	newCount := counter{}
	newCount.count = 0
	newCount.mu = new(sync.Mutex)
	return newCount
}

func (c *counter) increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func main() {
	myCounter := newCounter()
	countEnd := 20
	wg := new(sync.WaitGroup)
	wg.Add(countEnd)

	for i := 0; i < countEnd; i++ {
		go func(c *counter) {
			defer wg.Done()
			c.increment()
		}(&myCounter)

	}
	wg.Wait()

	fmt.Println(myCounter.count)
}
