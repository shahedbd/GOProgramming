package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("MyFiles\\WriteMe.txt", d1, 0644)
	check(err)

	//Display it
	ReadMe, err := ioutil.ReadFile("MyFiles\\WriteMe.txt")
	check(err)
	fmt.Print(string(ReadMe))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
