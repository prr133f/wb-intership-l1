package main

import (
	"context"
	"fmt"
	"time"
)

func goWithCtx(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершение горутины с контекстом")
			return
		default:
			fmt.Println("Работает горутина с контекстом")
			time.Sleep(time.Millisecond * 200)
		}
	}
}

func goWithCtxDeadline(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершение горутины с дедлайном")
			return
		default:
			fmt.Println("Работает горутина с дедлайном")
			time.Sleep(time.Millisecond * 200)
		}
	}
}

func goWithChan(quit chan int) {
	for {
		select {
		case <-quit:
			fmt.Println("Завершение горутины с каналом")
			return
		default:
			fmt.Println("Работает горутина с каналом")
			time.Sleep(time.Millisecond * 200)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go goWithCtx(ctx)
	quit := make(chan int)
	go goWithChan(quit)
	ctx, cancelDeadline := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*500))
	defer cancelDeadline()
	go goWithCtxDeadline(ctx)

	time.Sleep(time.Millisecond * 1000)
	cancel()
	quit <- 1
	time.Sleep(time.Millisecond * 100)
}
