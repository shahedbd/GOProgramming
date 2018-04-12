package main

import "os"
import "strings"
import "fmt"

func main() {

	os.Setenv("MyEnVar", "ID505")

	fmt.Println("My Environment Variables:", os.Getenv("MyEnVar"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println("********************************")

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
