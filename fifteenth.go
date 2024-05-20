package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func (P Person) employee() string {
	return (P.firstName + " " + P.lastName)

}

func main() {
	func() {
		fmt.Println("fifteenth file annonymous function")
	}()

	smk := Person{
		firstName: "Soumik",
		lastName:  "Chakraborty",
	}

	fmt.Println(smk.employee())

}
