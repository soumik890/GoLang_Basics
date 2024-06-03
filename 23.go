package main

import "fmt"

func printSlice[T any](s []T) {
	for index, item := range s {
		fmt.Println(index, item)
	}
}

func main() {
	fmt.Println("23rd file")

	arr1 := []int{100, 200, 300, 400, 500, 600}

	arr2 := []string{"soumik", "chakraborty"}

	printSlice(arr1)
	printSlice(arr2)

}
