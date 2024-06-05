package main

import (
	"fmt"
	"practice/obj"
)

func main() {
	p1 := obj.NewPoint(0, 0)
	p2 := obj.NewPoint(3, 4)
	fmt.Println(p1.Distance(p2))
}
