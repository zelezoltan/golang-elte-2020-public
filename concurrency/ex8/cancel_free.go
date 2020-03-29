package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() // OMIT
	finished := make(chan struct{})
	go func() {
		time.Sleep(100 * time.Millisecond)
		finished <- struct{}{}
	}()
	<-finished
	fmt.Println("took ", time.Since(start)) // OMIT
}
