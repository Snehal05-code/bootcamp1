package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}

	var wg sync.WaitGroup

	// Run 1000 goroutines
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc("somekey")
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value:", c.Value("somekey"))
}
