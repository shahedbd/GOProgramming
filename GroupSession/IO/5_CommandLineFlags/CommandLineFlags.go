package main

import "flag"
import "fmt"

func main() {

	MyStringFlagPtr := flag.String("MyStringFlag", "Hell GO Lovers!", "A string")
	MyNumberFlagPtr := flag.Int("MyNumberFlag", 42, "An int")
	MyBooleanFlagPtr := flag.Bool("MyBooleanFlag", false, "A bool")

	var MyLocalFlag string
	flag.StringVar(&MyLocalFlag, "varName", "Local values", "a string var")
	flag.Parse()

	fmt.Println("My String Flag:", *MyStringFlagPtr)
	fmt.Println("My Number Flag:", *MyNumberFlagPtr)
	fmt.Println("My Boolean Flag:", *MyBooleanFlagPtr)
	fmt.Println("My Local Flag:", MyLocalFlag)
	fmt.Println("Flag Argument:", flag.Args())
}

//Command-line flags are a common way to specify options for command-line programs.
