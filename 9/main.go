package main

import (
	"fmt"
)

func Transfer(src chan int, dest chan int) {
	for v := range src {
		dest <- v
	}
	close(dest)
}

func main() {
	nums := []int{2, 4, 6, 8, 10, 12, 14, 16}
	in := make(chan int)
	out := make(chan int)

	go Transfer(in, out)

	done := make(chan int)
	go func(out chan int, done chan int) {
		for v := range out {
			fmt.Println(v)
		}
		done <- 1
	}(out, done)

	for _, num := range nums {
		in <- num
	}
	close(in)

	<-done
}
