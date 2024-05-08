package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Booleans")

	testBool1 := true

	var testBool2 bool = false

	fmt.Println(testBool1, testBool2)

	fmt.Println(2+2, 2-2, 2*3, 2/2, 2%2)

	var z = 10
	z++

	var text string = "Hello world"

	text1 := "Hello developrs"

	intro := `hello world this is
	a multiline sring`

	const concat = "Concat output" + " Hello" + " world"

	fmt.Println(z, text, text1, intro, concat)

	fmt.Printf("%c", text1[0])
	fmt.Println(len(text1))

	fmt.Println(strings.Compare(text, text1))

	fmt.Println(strings.Contains(text, text1))

	fmt.Println(strings.ToLower(intro))
	fmt.Println(strings.ToUpper(intro))

}
