package main

import "fmt"

func printItem[T any](item1, item2 T) (T, T) {
	return item1, item2

}

func main() {
	fmt.Println("22nd file")
	num1, num2 := printItem(1, 2)
	str1, str2 := printItem("one", "two")
	bool1, bool2 := printItem(true, false)

	fmt.Println(num1, num2, str1, str2, bool1, bool2)

}
