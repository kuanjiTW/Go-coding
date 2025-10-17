/*
  (20 points)
  Rewrite the program into an equivalent one which uses an
  infinite for-loop to calculate the summation.
*/

package main

import "fmt"

func main() {
	sum := 0
	i := 0
	for {
		if i > 10 {
			break
		}
		sum += i
		i++
	}

	fmt.Println("sum =", sum)
}
