package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	in := bufio.NewReader(os.Stdin)
	s, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	s = strings.TrimSpace(s)
	fmt.Println(reverseWords(s))
}
