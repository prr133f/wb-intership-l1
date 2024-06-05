package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)

	ch := make(chan int)
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*time.Duration(n)))
	defer cancel()

	go func(c context.Context, ch chan int) {
		for {
			select {
			case <-c.Done():
				return
			case data, ok := <-ch:
				if !ok {
					return
				}
				fmt.Println(data)
			}
		}
	}(ctx, ch)

	go func(c context.Context, ch chan int) {
		for {
			select {
			case <-c.Done():
				return
			case ch <- rand.Int():
				time.Sleep(time.Millisecond * 100)
			}
		}
	}(ctx, ch)

	<-ctx.Done()
}
