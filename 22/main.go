package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b, c big.Int
	fmt.Scan(&a, &b)

	fmt.Printf("%v\n", c.Add(&a, &b))
	fmt.Printf("%v\n", c.Sub(&a, &b))
	fmt.Printf("%v\n", c.Mul(&a, &b))
	fmt.Printf("%v\n", c.Div(&a, &b))
}
