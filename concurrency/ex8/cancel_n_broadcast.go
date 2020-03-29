package main

import (
	"fmt"
	"time"
)

func main() {
	work := []int{1, 2, 3}
	start := time.Now() // START OMIT
	finished, cancel := make(chan struct{}), make(chan struct{})
	f := func(i int) {
		select {
		case <-time.After(100 * time.Millisecond):
		case <-cancel:
			fmt.Println("cancelled")
		}
		fmt.Println(i, time.Since(start))
		finished <- struct{}{}
	}
	for i := range work {
		go f(i)
	}
	go func() {
		close(cancel) // HL
	}()
	for range work {
		<-finished
	}
	fmt.Println("took ", time.Since(start)) // END OMIT
}
