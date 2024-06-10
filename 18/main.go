package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter struct {
	sync.RWMutex
	Value atomic.Uint32
}

func NewCounter() *Counter {
	return &Counter{Value: atomic.Uint32{}}
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.Value.Add(1)
}

func (c *Counter) GetValue() uint32 {
	c.RLock()
	defer c.RUnlock()
	return c.Value.Load()
}

func main() {
	counter := NewCounter()
	for i := 0; i < 1000; i++ {
		go counter.Inc()
	}

	time.Sleep(time.Millisecond * 200)
	fmt.Println(counter.GetValue())
}
