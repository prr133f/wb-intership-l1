package main

import "fmt"

func main() {
	var strs = []string{"cat", "cat", "dog", "cat", "tree"}
	var set = make(map[string]bool, len(strs))

	for _, el := range strs {
		set[el] = true
	}

	fmt.Println(set)
}
