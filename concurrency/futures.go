package main

import (
	"fmt"
	"time"
)

func translate(word string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		time.Sleep(100 * time.Millisecond)
		ch <- word
	}()
	return ch
}

func main() {
	fmt.Println(<-translate("hello"), <-translate("world"))
}
