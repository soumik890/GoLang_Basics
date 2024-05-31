package main

import "fmt"

type animal interface {
	GetInfo() string
}

type Cat struct {
	name  string
	color string
}

func (c Cat) GetInfo() string {
	return fmt.Sprintln("name:",c.name, "color:", c.color)
}

func printAnimal(item animal) {
	fmt.Println(item.GetInfo())

}

func main() {

	fmt.Println("20th file")

	cat := Cat{name: "kitty", color: "orange"}

	printAnimal(cat)
}
