package main

import "fmt"

func BinarySearch(arr []int, target int) bool {
	if len(arr) == 0 {
		return false
	}

	base := arr[len(arr)/2]
	if target < base {
		return BinarySearch(arr[:len(arr)/2], target)
	} else if target > base {
		return BinarySearch(arr[len(arr)/2+1:], target)
	} else {
		return true
	}
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var target int
	fmt.Scan(&target)
	fmt.Println(BinarySearch(nums, target))
}
