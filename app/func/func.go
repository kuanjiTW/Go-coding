package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	var f func(int, int) int = add
	r := f(5, 3)
	fmt.Println(r)
}
