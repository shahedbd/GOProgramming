package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)
}

//Notes:
/*
a, b := "second", "first"
fmt.Println(a, b) // Prints "second first"
b, a = a, b
fmt.Println(a, b) // Prints "first second"

*/
