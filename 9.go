package main

import "fmt"

func main() {
	fmt.Println("nineth file")
	num1 := make([]int, 5, 8)

	num1[0] = 1
	num1[1] = 10
	num1[2] = 20
	num1[3] = 30
	num1[4] = 30

	fmt.Println(num1)
	fmt.Println("Length of  num1 is", len(num1))
	fmt.Println("Capacity of num1 is", cap(num1))

	ppl := []string{"soumik", "piyali", "soma", "santosh"}

	for index, value := range ppl {
		fmt.Println(index, value)
	}

}
