package main

import "fmt"

func main() {
	var x float64
	var y int
	x = 20.0
	y = 5
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)

	fmt.Print(x - float64(y))
}
