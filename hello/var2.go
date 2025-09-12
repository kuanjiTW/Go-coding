package main

import "fmt"

func main() {

	var price float64 = 25
	var units float64 = 20
	var total float64 = price * units
	var discount float64 = 0.8
	total = total * discount

	fmt.Println(total)

	var n int = 25
	n /= 2
	n += 5

	fmt.Println(n)

	n = 25
	fmt.Println(n)
	fmt.Printf("n:%v\n", n)
	fmt.Printf("n:%+v\n", n)
	fmt.Printf("n:%#v\n", n)
	fmt.Printf("n:%T\n", n)
	fmt.Printf("n:%d\n", n)
	fmt.Printf("n:%s\n", n)

	a := 20
	m := 17

	fmt.Printf("the value of n is %v and the type is %T", a, m)

	// n2 := 20
	// m2 := "hello"

	// users
}
