package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct{}

func (c cat) Speak() string {
	return "Meow Meow!"
}

type dog struct{}

func (d dog) Speak() string {
	return "Woof Woof!"
}

func main() {
	c := cat{}
	d := dog{}
	fmt.Println(c.Speak())
	fmt.Println(d.Speak())
}
