// Add a new variable scope such that the program outputs:
//   11
//   6
//   22
// The scope should be as small as possible.

package main

import "fmt"

var x = 10

func main() {
	x = x + 1
	fmt.Println(x)
	x := 6
	fmt.Println(x)
	funcA()
}

func funcA() {
	x = x * 2
	fmt.Println(x)
}
