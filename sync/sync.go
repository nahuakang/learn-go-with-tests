package sync

import "sync"

// Counter type
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter returns a new Counter
func NewCounter() *Counter {
	return &Counter{}
}

// Inc is a Counter method to increase value by 1
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value is a Counter method to return the value
func (c *Counter) Value() int {
	return c.value
}
