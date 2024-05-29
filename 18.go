package main

import "fmt"

func main() {
	fmt.Println("18th File")

	x := 40
	var pointerTox *int = &x

	fmt.Println("Value of x:", x)
	fmt.Println("memory address of x:", &x)

	fmt.Println("Value pointed by pointex x", pointerTox)
	fmt.Println("Memory Addres stored in pointex x", *pointerTox)

}
