package main

import (
	"fmt"
)

func foo(c <-chan string) {
	fmt.Println("foo", <-c)
}

func bar(c <-chan string) {
	fmt.Println("bar", <-c)
}

func main() {
	c := make(chan string)
	go foo(c)
	go bar(c)
	c <- "nigga"
	c <- "2nigga"
}
