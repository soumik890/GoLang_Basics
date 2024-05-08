package main

import (
	"fmt"
)

func main() {
	fmt.Println("fifth file")
	password := "12345678"
	if len(password) > 7 {
		fmt.Println("The length is greater than 7")
	} else if len(password) > 8 {
		fmt.Println("Its length is greater than 8")
	} else {
		fmt.Println("The length is not greater than 7 ")
	}

	for i := 0; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	for i := 0; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}
