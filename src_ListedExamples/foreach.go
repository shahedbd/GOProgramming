package main

import "fmt"

func main() {

	arrayOne := [3]string{"Apple", "Mango", "Banana"}

	for index, element := range arrayOne {
		fmt.Println(index, element)
	}

}
