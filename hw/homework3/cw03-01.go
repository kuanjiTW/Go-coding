/*
  (20 points)
  Write a program that prints all even numbers between 0 and 20
  (both inclusive) in descending order using a for-loop. The
  fewer `fmt.Println` (or other similar printing functions) calls
  you use, the more points you get.
*/

package main

import "fmt"

func main() {
	for i := 20; i >= 0; i -= 2 {
		fmt.Println(i)
	}
}
