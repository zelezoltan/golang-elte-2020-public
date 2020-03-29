package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() // OMIT
	finished, cancel := make(chan struct{}), make(chan struct{})
	f := func() {
		select {
		case <-time.After(100 * time.Millisecond):
		case <-cancel:
		}
		finished <- struct{}{}
	}
	go f() // HL
	go f() // HL
	go func() { cancel <- struct{}{} }()
	<-finished                              // HL
	<-finished                              // HL
	fmt.Println("took ", time.Since(start)) // OMIT
}
