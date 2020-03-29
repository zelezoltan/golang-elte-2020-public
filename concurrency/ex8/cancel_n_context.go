package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	work := []int{1, 2, 3}
	start := time.Now() // START OMIT
	finished := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background()) // HL
	f := func(ctx context.Context, i int) {                 // HL
		select {
		case <-time.After(100 * time.Millisecond):
		case <-ctx.Done(): // HL
			fmt.Println(ctx.Err()) // HL
		}
		fmt.Println(i, time.Since(start))
		finished <- struct{}{}
	}
	for i := range work {
		go f(ctx, i) // HL
	}
	go func() {
		cancel() // HL
	}()
	for range work {
		<-finished
	}
	fmt.Println("took ", time.Since(start)) // END OMIT
}
