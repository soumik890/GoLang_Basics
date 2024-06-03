package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 { // while loop in golang
		sum += sum
	}
	fmt.Println(sum)
}
