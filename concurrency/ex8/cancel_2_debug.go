package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() // OMIT
	finished, cancel := make(chan struct{}), make(chan struct{})
	f := func(i int) { // HL
		select {
		case <-time.After(100 * time.Millisecond):
		case <-cancel: // HL
			fmt.Println("cancelled") // HL
		}
		fmt.Println(i, time.Since(start)) // HL
		finished <- struct{}{}
	}
	go f(1) // HL
	go f(2) // HL
	go func() { cancel <- struct{}{} }()
	<-finished
	<-finished
	fmt.Println("took ", time.Since(start)) // END OMIT
}
