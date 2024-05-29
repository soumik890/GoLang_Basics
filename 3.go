package main

import (
	"fmt"
	"strings"
)

func main() {
	var hello = "Hello"
	fmt.Println(hello)
	fmt.Println(strings.Count(hello, "l"))
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(2 == 2)
	fmt.Println(2 != 2)
	fmt.Println(2 < 2)
	fmt.Println(2 > 2)
	fmt.Println(2 >= 2)
	fmt.Println(2 <= 2)
}
