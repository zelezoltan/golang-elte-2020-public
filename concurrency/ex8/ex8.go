package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	results := make(map[string]int)
	mu := sync.Mutex{}
	// MAIN-START OMIT
	eg, _ := errgroup.WithContext(context.Background()) // HL
	for _, path := range Files() {
		path := path
		eg.Go(func() error { // HL
			hash, err := Hash(path)
			if err != nil {
				return err // HL
			}
			mu.Lock()
			defer mu.Unlock()
			results[path] = hash
			return nil
		})
	}
	// MAIN-END OMIT
	if err := eg.Wait(); err != nil { // HL
		fmt.Printf("ERROR %s\n", err)
	} else {
		fmt.Println(results)
	}
	fmt.Println("took ", time.Since(start))
}

func Hash(path string) (int, error) {
	if len(path) == 12 {
		return 0, fmt.Errorf("cannot calculate hash for %q", path)
	}
	time.Sleep(100 * time.Millisecond)
	return len(path), nil
}

func Files() []string {
	return []string{"ex1/ex1.go", "ex2/cksum.go", "ex3/ex3.go"}
}
