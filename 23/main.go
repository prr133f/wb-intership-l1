package main

import "fmt"

func removeFromSlice[T any](slice []T, idx int) []T {
	return append(slice[:idx+1], slice[idx+2:]...)
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var target int
	fmt.Scan(&target)
	fmt.Println(removeFromSlice(nums, target-1))
}
