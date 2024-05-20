package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {

	multiply := func(x int, y int) int {
		return x * y
	}

	fmt.Println("Fourteenth File")
	fmt.Println(add(1, 2))
	fmt.Println(multiply(2, 3))
}
