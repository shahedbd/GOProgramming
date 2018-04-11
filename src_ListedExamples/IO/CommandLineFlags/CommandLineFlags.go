package main

import "flag"
import "fmt"

func main() {

	MyStringFlagPtr := flag.String("MyStringFlag", "My String Flag", "A string")
	numbPtr := flag.Int("numb", 42, "An int")
	boolPtr := flag.Bool("fork", false, "A bool")

	var svar string
	flag.StringVar(&svar, "svar5", "bar", "a string var")
	flag.Parse()

	fmt.Println("My String Flag:", *MyStringFlagPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
