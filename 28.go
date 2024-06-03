package main

import (
	"fmt"
	"sync"
)

func printer[T any](wg *sync.WaitGroup, item1, item2 T) {
	defer wg.Done()
	fmt.Println(item1, item2)
}

func main() {
	fmt.Println("28th File")

	var wg sync.WaitGroup
	wg.Add(3) //Add the number of goroutine you are waiting for

	go printer(&wg, "soumik", "chakraborty")
	go printer(&wg, "is", "a")
	go printer(&wg, "fullstack", "dev")

	wg.Wait()

	fmt.Println("All messages printed")

}
