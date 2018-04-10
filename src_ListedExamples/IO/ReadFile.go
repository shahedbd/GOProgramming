package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	ReadMe, err := ioutil.ReadFile("MyFiles\\ReadMe.txt")
	check(err)
	fmt.Print(string(ReadMe))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
