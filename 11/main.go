package main

import "fmt"

func main() {
	var setA = make(map[int]bool)
	var setB = make(map[int]bool)
	var intersection = make(map[int]bool)

	setA[1] = true
	setA[2] = true
	setA[3] = true
	setA[4] = true
	setA[5] = true
	setA[6] = true
	setA[7] = true
	setA[8] = true
	setA[9] = true
	setA[10] = true
	setB[2] = true
	setB[4] = true
	setB[6] = true
	setB[8] = true
	setB[10] = true
	setB[12] = true
	setB[14] = true
	setB[16] = true
	setB[18] = true
	setB[20] = true

	for k := range setA {
		if _, ok := setB[k]; ok {
			intersection[k] = true
		}
	}

	fmt.Println(intersection)
}
