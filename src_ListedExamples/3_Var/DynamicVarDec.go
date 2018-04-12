package main

import "fmt"

func main() {
	//Dynamic Type Declaration / Type Inference in Go
	var x float64
	x = 20.0
	y := "My string"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
}
