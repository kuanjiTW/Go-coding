package main

import "fmt"

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	n := 25
	p := new(int)
	*p = 30

	q := &n
	*q = 3
	fmt.Println(p, *p, q, *q, n)

	a, b := 4, 8
	swap(&a, &b)
	fmt.Println(a, b)

}
