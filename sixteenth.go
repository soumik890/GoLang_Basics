package main

import "fmt"

func addName(name string, callback func(string)) {
	callback(name)
}

func main() {
	fmt.Println("Sixteenth File")

	addName("Soumik", func(s string) {
		fmt.Println("This is a callback")
	})

}
