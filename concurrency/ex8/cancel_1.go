package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() // OMIT
	finished, cancel := make(chan struct{}), make(chan struct{})
	go func() {
		select {
		case <-time.After(100 * time.Millisecond): // HL
		case <-cancel: // HL
		}
		finished <- struct{}{}
	}()
	go func() { cancel <- struct{}{} }() // HL
	<-finished
	fmt.Println("took ", time.Since(start)) // OMIT
}
