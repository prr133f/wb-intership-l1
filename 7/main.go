package main

import (
	"fmt"
	"sync"
)

var sharedMap = make(map[string]any, 10)

func SetToSharedMap(key string, value any, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	sharedMap[key] = value
}

func main() {
	var mu sync.Mutex

	go SetToSharedMap("key1", "value1", &mu)
	go SetToSharedMap("key2", "value2", &mu)
	go SetToSharedMap("key1", 123, &mu)

	fmt.Println(sharedMap)
}
