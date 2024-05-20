package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func (P Person) employee() {
	fmt.Println("I am", P.firstName, P.lastName)

}

func main() {
	func() {
		fmt.Println("fifteenth file auto call annonymous function")
	}()

	smk := Person{
		firstName: "Soumik",
		lastName:  "Chakraborty",
	}

	smk.employee()

}
