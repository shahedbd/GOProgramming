package main

import (
	"fmt"
	"time"
)

func main() {
	go GoKeyWordfunc("World")
	GoKeyWordfunc("Hello")
}

func GoKeyWordfunc(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

//A goroutine is a lightweight thread managed by the Go runtime.
//go f(x, y, z)
