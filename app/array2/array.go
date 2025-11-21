package main

import "fmt"

func main() {
	//incr(5)
	//incr(3.14)
	//incr("Hello world")

	var x any = 5
	y := x.(int)
	fmt.Println(x, y)
}
