package main

import "fmt"

func main() {
	var setA = make(map[int]bool)
	var setB = make(map[int]bool)
	var intersection = make(map[int]bool)

	for i := 1; i <= 10; i++ {
		setA[i] = true
	}
	for i := 1; i <= 10; i++ {
		setB[i*2] = true
	}

	for k := range setA {
		if _, ok := setB[k]; ok {
			intersection[k] = true
		}
	}

	fmt.Println(intersection)
}
