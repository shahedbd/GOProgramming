package main

//Senders have the ability
//to close the channel to notify receivers that
// no more data will be sent on the channel
import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 5; i++ {
		chnl <- i
	}
	close(chnl)
}
func main() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}
