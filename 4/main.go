package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var workers int
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Scan(&workers)

	data := make(chan int)

	for i := 0; i < workers; i++ {
		go func(data chan int) {
			for v := range data {
				fmt.Println(v)
			}
		}(data)
	}

	for {
		data <- rand.Int()
		time.Sleep(time.Millisecond * 200)
	}

}
