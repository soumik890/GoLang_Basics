package main

import "fmt"

type animal interface {
	GetInfo() string
}

type cat struct {
	name  string
	color string
}

func (c cat) GetInfo() string {
	fmt.Sprintln("name:%s, color:%s", c.name, c.color)
}

func main() {

	fmt.Println("20th file")
}
