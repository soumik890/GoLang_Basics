package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("29th File")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message form channel 1"
	}()
	go func() {
		time.Sleep(4 * time.Second)
		ch2 <- "Message form channel 2"
	}()

	select {

	case msg1 := <-ch1:
		fmt.Println("Received from channel 1:", msg1)

	case msg2 := <-ch2:
		fmt.Println("Received from channel 2:", msg2)

	case <-time.After(5 * time.Second):
		fmt.Println("no communication")

	}

}
