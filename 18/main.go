package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Counter struct {
	Value atomic.Uint32
}

func NewCounter() *Counter {
	return &Counter{Value: atomic.Uint32{}}
}

func (c *Counter) Inc() {
	c.Value.Add(1)
}

func (c *Counter) GetValue() uint32 {
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
