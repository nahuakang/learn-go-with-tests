package sync

// Counter type
type Counter struct {
	value int
}

// Inc is a Counter method to increase value by 1
func (c *Counter) Inc() {
	c.value++
}

// Value is a Counter method to return the value
func (c *Counter) Value() int {
	return c.value
}
