package main

import "fmt"

/* global variable declaration */
var g int

func main() {
	/* local variable declaration */
	var a, b int

	/* actual initialization */
	a = 10
	b = 20
	g = a + b

	fmt.Printf("value of a = %d, b = %d and g = %d\n", a, b, g)
}
