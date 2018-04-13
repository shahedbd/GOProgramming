package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
	done <- true
}

// func hello(done chan bool) {
//     fmt.Println("Hello world goroutine")
//     done <- true
// }
func main() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")

	// done := make(chan bool)
	// go hello(done)
	// <-done
	// fmt.Println("main function")
}
