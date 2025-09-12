// Add a new variable scope such that the program outputs:
//   11
//   6
//   22
// The scope should be as small as possible.

package main

import "fmt"

func main() {
	x := 10
	x = x + 1
	fmt.Println(x)
	{
		x := 6
		fmt.Println(x)
	}
	x = x * 2
	fmt.Println(x)
}
