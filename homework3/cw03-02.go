/*
  (20 points)
  Write a program that prints a rotated pyramid pattern.
  The height of the pyramid is obtained from a call to
  `fmt.Scan`. You may assume that the input height is always
  positive. As an example of rotated pyramid patterns, for
  a pyramid of height 4, the program should prints:
  ```
  Height: 4
  *
  **
  ***
  ****
  ***
  **
  *
  ```
  As another example, for a pyramid of height 6, the program
  should prints:
  ```
  Height: 6
  *
  **
  ***
  ****
  *****
  ******
  *****
  ****
  ***
  **
  *
  ```
  You may use `strings.Repeat(s, n)` to repeat a string `s`
  n times.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var height int
	fmt.Print("Height:")
	fmt.Scan(&height)
	for i := 1; i <= height; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
	for i := height - 1; i >= 1; i-- {
		fmt.Println(strings.Repeat("*", i))
	}
}
