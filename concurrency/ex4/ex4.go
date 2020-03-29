package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	// TODO: wait for all results
	results := make(map[string]int)
	mu := sync.Mutex{}
	for _, path := range Files() {
		go func(path string) {
			hash := Hash(path)
			mu.Lock()
			defer mu.Unlock()
			results[path] = hash
		}(path)
	}
	// END OMIT

	fmt.Println(results)
	fmt.Println("took ", time.Since(start))
}

func Hash(path string) int {
	time.Sleep(100 * time.Millisecond)
	return len(path) // *not* collision free hash
}

func Files() []string {
	return []string{"ex1/ex1.go", "ex2/cksum.go", "ex3/ex3.go"}
}
