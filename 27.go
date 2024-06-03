package main

import (
	"fmt"
)

func main() {

	unbufferedChannel := make(chan int)      //unbuffered channel with no capacity
	bufferedChannel := make(chan string, 20) //buffered channel with capacity 10

	go func() {
		unbufferedChannel <- 20
	}()

	go func() {
		bufferedChannel <- "soumik"
		bufferedChannel <- "chakraborty"
		close(bufferedChannel)
	}()

	valueUnbuffered := <-unbufferedChannel
	valueBuffered := <-bufferedChannel

	fmt.Println("value from unbuffered channel", valueUnbuffered)

	for index, item := range valueBuffered {
		fmt.Println("value from bufferd channel", index, item)
	}

	for item := range bufferedChannel {
		fmt.Println("value from bufferd channel", item)
	}

}
