package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Coder struct {
	Person
	yoe int
}

func main() {
	fmt.Println("12th File")

	test := Coder{
		Person: Person{
			name: "Soumik",
			age:  30,
		},
		yoe: 5,
	}

	fmt.Println(test)

}
