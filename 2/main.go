package main

import (
	"fmt"
	"sync"
)

// Канал не сохраняет порядок чисел. Для сохранения использовать слайс
func main() {
	nums := []int{2, 4, 6, 8, 10}
	ans := make(chan int, len(nums))

	var wg sync.WaitGroup
	for _, num := range nums {
		wg.Add(1)
		go func(n int) {
			ans <- n * n
			wg.Done()
		}(num)
	}

	go func() {
		wg.Wait()
		close(ans)
	}()

	for el := range ans {
		fmt.Println(el)
	}
}
