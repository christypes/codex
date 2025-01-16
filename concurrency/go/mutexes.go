package main

import (
	"fmt"
	"sync"
)

// Struct packs the lock with the shared var nicely
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Method for Container struct
func (c *Container) inc(name string) {
	c.mu.Lock()
	// The defer ensures that the unlock happens not even if function errors.
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	var wg sync.WaitGroup
	c := Container{counters: map[string]int{"a": 0, "b": 0}}
	doInc := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	wg.Add(3) // Sets up 3 increments to run 3 goroutines
	go doInc("a", 10000)
	go doInc("a", 15000)
	go doInc("b", 10000)
	wg.Wait() // blocks the main() until all returns
	fmt.Println(c.counters)
}

// WaitGroup ... Counter that blocks until zero
// 	.Add(int): Increments by int
//	.Done(): Decrements by 1
//	.Wait(): Blocks until counter is zero
