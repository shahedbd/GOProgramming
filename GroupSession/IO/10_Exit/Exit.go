package main

import "fmt"
import "os"

func main() {

	defer fmt.Println("Not Printed! :( ")

	os.Exit(3)
}
