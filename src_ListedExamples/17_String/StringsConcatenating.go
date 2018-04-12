package main

import (
	"fmt"
	"strings"
)

func main() {
	greetings := []string{"Hello", "world!"}
	fmt.Println(strings.Join(greetings, " "))
}
