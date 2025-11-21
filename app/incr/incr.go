package main

import "fmt"

func incr_decr(n int) (x int, y int) {
	x = n + 1
	y = n - 1
	return
}

func sum(ns ...int) int {
	s := 0
	for _, n := range ns {
		s += n
	}
	return s
}

func main() {
	x, y := incr_decr(5)
	fmt.Println(x, y)

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(2, 4, 6, 8, 10))
}
