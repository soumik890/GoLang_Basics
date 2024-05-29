package main

import "fmt"

func panicTester(name string, age int) {

	if age > 60 {
		panic("Old man detected")
	}
}

func main() {
	fmt.Println("Hello 19th")

	panicTester("soumik", 80)

}
