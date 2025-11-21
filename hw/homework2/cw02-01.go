/*
  (20 points)
  Write a program that fulfills the following requirements.
  - Ask for an integer from the standard input (using fmt.Scan).
    A prompt may be printed before asking for the input integer.
  - Print the string "multiple of 3" if the input integer is a
    multiple of 3 but not a multiple of 7.
  - Print the string "multiple of 7" if the input integer is a
    multiple of 7 but not a multiple of 3.
  - Print the string "bang!" if the input integer is both a
    multiple of 3 and a multiple of 7.
*/

package main

import "fmt"

func main() {
	var (
		v1 int
	)
	fmt.Println("enter a integer:")
	fmt.Scan(&v1)
	if v1%3 == 0 && v1%7 != 0 {
		fmt.Println("multiple of 3")
	} else if v1%3 != 0 && v1%7 == 0 {
		fmt.Println("multiple of 7")
	} else if v1%3 == 0 && v1%7 == 0 {
		fmt.Println("bang!")
	} else {
		fmt.Println("not a multiple of 3 or 7")
	}
}
