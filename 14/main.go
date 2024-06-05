package main

import "fmt"

func main() {
	var value interface{}

	switch value.(type) {
	case int:
		fmt.Println("Value is type of int")
	case string:
		fmt.Println("Value is type of string")
	case bool:
		fmt.Println("Value is type of bool")
	case chan interface{}:
		fmt.Println("Value is type of chan")
	default:
		fmt.Println("Value is unknown type")
	}
}
