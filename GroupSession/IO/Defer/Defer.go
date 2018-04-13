package main

import "fmt"

func main() {
	defer fmt.Println("World")

	fmt.Println("Hello")
}

//A defer statement defers the execution of a function until the surrounding function returns.
