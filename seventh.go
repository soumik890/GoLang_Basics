package main

import "fmt"

func main() {
	fmt.Println("Seventh File")

	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	var number [5]int

	number[0] = 20
	number[1] = 30
	number[2] = 40

	fmt.Println(number)

	var arr3 = [8]string{"js", "golang", "react", "node", "React-Natve", "ts"}

	for i := 0; i < len(arr3); i++ {
		fmt.Println(i, arr3[i])
	}

	fmt.Println(cap(arr3))
	fmt.Println(len(arr3))
}
