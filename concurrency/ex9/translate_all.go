package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"sync"
	"time"
)

func translate(word string) string {
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // HL
	return word
}

func all(words ...string) map[string]string {
	work := make(map[string]string, len(words))
	mu := &sync.Mutex{}
	// TODO: return within 100ms, even without all results
	// START OMIT
	eg, _ := errgroup.WithContext(context.Background())
	for _, word := range words {
		word := word
		eg.Go(func() error {
			translated := translate(word)
			mu.Lock()
			defer mu.Unlock()
			work[word] = translated
			return nil
		})
	}
	eg.Wait()
	// END OMIT
	return work
}

func main() {
	start := time.Now()
	rand.Seed(start.UnixNano())
	fmt.Println(all("hello", "world"))
	fmt.Println("took ", time.Since(start))
}
