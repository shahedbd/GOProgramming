package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Type x for exit!")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		abc := scanner.Text()
		if abc == "x" {
			os.Exit(3)
		}
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

//A line filter is a common type of program that reads input on stdin,
//processes it, and then prints some derived result to stdout.
