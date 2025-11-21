/*
(20 points)
The Fibonacci sequence is a number series where each number is
the sum of the two preceding ones. By definition, the first number
and the second number in the series are 0 and 1 respectively.
Below is a partial list of the Fibonacci sequence:

	0 1 1 2 3 5 8 13 21 34 ...

Write a program that prints the first n numbers in the Fibonacci
sequence. The number n is obtained from `fmt.Scan` and is assumed
to be positive.
Hints:
- Treat the first and second numbers as special cases.
- Use two variables to remember the last two numbers printed so far.
- The next number to be printed is the sum of the two variables.
- Update the two variables once the next number has been printed.
*/
package main

import "fmt"

func main() {
	var n int
	fmt.Print("n:")
	fmt.Scan(&n)

	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
}
