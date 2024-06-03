package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func main() {
	go printNumbers()
	fmt.Println("26th File")
	time.Sleep(5 * time.Second)

}
