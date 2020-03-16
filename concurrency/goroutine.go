package main

import (
	"fmt"
	"time"
)

func say(greeting string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(greeting)
}

func main() {
	go say("hello")
	go say("world")
	time.Sleep(200 * time.Millisecond)
}
