package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args), os.Args)

	argsWithProg := os.Args
	fmt.Println(argsWithProg)
}
