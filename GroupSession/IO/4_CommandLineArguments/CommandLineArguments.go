package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args))

	argsWithProg := os.Args
	fmt.Println(argsWithProg)
}

//Command-line arguments are a common way to parameterize execution of programs.
