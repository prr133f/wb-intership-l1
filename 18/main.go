package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sync.RWMutex
	Value uint
}

func NewCounter() *Counter {
	return &Counter{Value: 0}
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.Value++
}

func (c *Counter) GetValue() uint {
	c.RLock()
	defer c.RUnlock()
	return c.Value
}

func main() {
	counter := NewCounter()
	for i := 0; i < 1000; i++ {
		go counter.Inc()
	}

	time.Sleep(time.Millisecond * 200)
	fmt.Println(counter.GetValue())
}
