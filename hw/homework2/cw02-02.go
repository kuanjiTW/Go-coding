/*
  (20 points)
  A year is a leap year if it meets one of the following two conditions:
  - It is divisible by 400.
  - OR, it is divisible by 4 but not by 100.

  Add an if statement (possibly with else-if statements) between
  `// <code begin>` and `// <code end>` with a proper Boolean condition.
  Assign true to is_leap_year in the body of the if statement (possibly also
  in the bodies of the else-if statements) such that is_leap_year is
  true at the end of isLeapYear if and only if the input year is a leap year.
  The output of this program should be:

    === Leap Year Tests ===
    2000: true
    1900: false
    2024: true
    2023: false

  Note: The precedence of && is higher than ||. For example, `b1 || b2 && b3`
        is interpreted as `b1 || (b2 && b3)`.
*/

package main

import "fmt"

func isLeapYear(year int) bool {
	var is_leap_year bool
	// <code begin>
	if year%400 == 0 {
		is_leap_year = true
	} else if year%4 == 0 && year%100 != 0 {
		is_leap_year = true
	} else {
		is_leap_year = false
	}
	// <code end>
	return is_leap_year
}

func main() {
	fmt.Println("=== Leap Year Tests ===")
	fmt.Printf("2000: %v\n", isLeapYear(2000))
	fmt.Printf("1900: %v\n", isLeapYear(1900))
	fmt.Printf("2024: %v\n", isLeapYear(2024))
	fmt.Printf("2023: %v\n", isLeapYear(2023))
}
