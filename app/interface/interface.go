package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct {
	name string
	age  int
}
type dog struct{}
type person struct {
	name string
}

func (c cat) Speak() string {
	return "Meow Meow"
}
func (c cat) GetName() string {
	return c.name
}
func (c cat) GetAge() int {
	return c.age
}
func (c *cat) SetName(name string) {
	c.name = name
}

func (d dog) Speak() string {
	return "Woof Woof"
}

func (p person) Speak() string {
	return "Hello, my name is " + p.name
}

type I int

func (i I) Speak() string {
	return "I am an integer"
}

func saySomething(ss ...Speaker) {
	for _, s := range ss {
		fmt.Println(s.Speak())
	}
}

func main() {
	n := 10
	fmt.Println(I(n).Speak())
	fmt.Println()

	c := cat{}
	d := dog{}
	p := person{name: "Kevin"}
	saySomething(c, d, p)
	fmt.Println()

	c = cat{"Alice", 3}
	c.SetName("Bob")
	fmt.Println(c.GetName())
	// c1.Greeting = func() string {
	// 	return "meow meow meow , meoooooooow"
	// }
	// fmt.Println("c1:")
	// fmt.Println(c1.Speak())
	// fmt.Println(c1.Greeting())

	// c2 := cat{}
	// c2.Greeting = func() string {
	// 	return "meow~~~~~~~~~~~~~~~~"
	// }
	// fmt.Println("c2:")
	// fmt.Println(c2.Speak())
	// fmt.Println(c2.Greeting())

}
