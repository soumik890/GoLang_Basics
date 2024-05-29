package main

import "fmt"

func greet() {
	defer fmt.Println("Hello world deffered") // delayed execution
	fmt.Println("Non Deffered")
}

func main() {
	fmt.Println("17th File")

	greet()

}
