package main

import "fmt"

func sendData(ch chan<- int) {
	ch <- 20
}

func main() {

	channel := make(chan int)
	go sendData(channel)

	value := <-channel

	fmt.Println("Received value is", value)

}
