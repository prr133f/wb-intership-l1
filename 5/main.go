package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)
	wg := new(sync.WaitGroup)

	ch := make(chan int)
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*time.Duration(n)))
	defer cancel()
	wg.Add(2)

	go func(c context.Context, ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
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
	}(ctx, ch, wg)

	go func(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case ch <- rand.Int():
				time.Sleep(time.Millisecond * 100)
			}
		}
	}(ctx, ch, wg)

	<-ctx.Done()
	wg.Wait()
}
