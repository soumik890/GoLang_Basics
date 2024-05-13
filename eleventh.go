package main

import "fmt"

func main() {
	fmt.Println("Eleventh file")

	type Person struct {
		name    string
		age     int
		married bool
	}

	var userOne Person

	userOne.name = "soumik"
	userOne.age = 31
	userOne.married = false

	fmt.Println(userOne)

}
