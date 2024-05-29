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
	fmt.Println(userOne.name, userOne.age, userOne.married)

	testVar := struct {
		first string
		last  string
	}{
		first: "soumik",
		last:  "chakraborty",
	}

	fmt.Println(testVar, "***************")

}
