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
	n = n / 2

	fmt.Println()
}
