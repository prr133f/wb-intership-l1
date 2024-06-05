package main

import (
	"fmt"
	"sync"
)

func main() {
	var ans int
	nums := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	for _, num := range nums {
		wg.Add(1)
		go func(n int) {
			ans += n * n
			wg.Done()
		}(num)
	}

	wg.Wait()

	fmt.Println(ans)
}
