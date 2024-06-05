package main

import "fmt"

func main() {
	var n, i int64
	fmt.Scan(&n, &i)

	var mask int64 = 1 << i

	fmt.Printf("%064b\n", n^mask)
}
