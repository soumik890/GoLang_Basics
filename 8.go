package main

import "fmt"

func main() {
	fmt.Println("Eighth File")
	slice := make([]string, 3, 5)
	fmt.Println(slice)
	fmt.Println("Length", len(slice))
	fmt.Println("Capacity", cap(slice))

	num := []int{10, 20, 30, 40}
	fmt.Println(num)
	fmt.Println(num[:])
	fmt.Println(num[0:2])

	num = append(num, 50, 60, 70, 80)
	fmt.Println(num)
}
