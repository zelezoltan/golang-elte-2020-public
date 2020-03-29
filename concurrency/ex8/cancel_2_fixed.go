package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() // OMIT
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
	go f(1)
	go f(2)
	go func() {
		cancel <- struct{}{}
		cancel <- struct{}{} // HL
	}()
	<-finished
	<-finished
	fmt.Println("took ", time.Since(start)) // END OMIT
}
