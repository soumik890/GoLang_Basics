package main

import "fmt"

func main() {
	fmt.Println("Sixth Doc")

	i := 0
	input := 1

	for i < 10 {
		fmt.Println(i, "Prining the value of i")
		i++
	}

	switch input {

	case 1:
		fmt.Println("1 is Input")

	case 2:
		fmt.Println("2 is Input")

	case 3:
		fmt.Println("3 is Input")

	default:
		fmt.Println("Default Triggered")

	}

}
