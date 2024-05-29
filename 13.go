package main

import "fmt"

func printString(str string) {
	fmt.Println(str)
}

func showNumbers(str string, numbers ...int) {
	fmt.Println(str, numbers)
}

func showChars(str ...string) {
	for index, item := range str {
		fmt.Println(index, item)
	}
}

func main() {
	fmt.Println("13th file")

	printString("soumik")
	showNumbers("smk", 1, 2, 3, 4, 5)
	showChars("smk1", "smk2", "smk3", "smk4")
}
