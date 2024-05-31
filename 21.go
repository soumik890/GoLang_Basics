package main

import "fmt"

type Person interface {
	UserName() string
	Profession() string
}

type user1 struct {
	name   string
	gender string
}

type user2 struct {
	title string
	age   int
}

func (u user1) UserName() string {
	return "Soumik"
}

func (u user2) Profession() string {
	return "Backend Developer"
}

func main() {
	soumik := user1{
		name:   "soumik",
		gender: "m",
	}

	chakraborty := user2{
		title: "chakraborty",
		age:   10,
	}

	fmt.Println(soumik.UserName())
	fmt.Println(chakraborty.Profession())

}
