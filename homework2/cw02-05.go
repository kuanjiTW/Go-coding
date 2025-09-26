/*
  (20 points)
  Rewrite the program such that no if statement is used.
  The behavior of the program should remain unchanged
  after the rewrite.
*/

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter an integer: ")
	switch cnt, err := fmt.Scan(&n); {
	case err != nil || cnt != 1:
		fmt.Println("not an integer")
	case n < 0:
		fmt.Println("negative")
	default:
		fmt.Println("non-negative")
	}
}
