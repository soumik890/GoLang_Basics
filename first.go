package main

import "fmt"

const (
	a = 10
	b = 20
	c = 30
)

func main() {
	var Test1 string = "Test1"
	Test2 := "Test2"
	var Test3, Test4, Test5, Test6 int = 3, 4, 5, 6
	Test7, Test8, Test9, Test10 := 7, 8, 9, 10
	TestNumber := 500
	var (
		Test11 int = 11
		Test12 int = 12
		Test13 int = 13
	)

	const constantVariable = "Immutable String"

	fmt.Println(Test1, Test2, Test3, Test4, Test5, Test6, Test7, Test8, Test9, Test10, TestNumber, Test11, Test12, Test13, constantVariable, a, b, c)
}
