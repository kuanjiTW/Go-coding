/*
(20 points)
Write a program that fulfills the following requirements.
  - Declare distinct constants: Monday, Tuesday, Wednesday,
    Thursday, Friday, Saturday, Sunday.
  - Initialize a variable with some day of the week.
  - Write a switch statement without the default case and inside
    the switch statement,
  - print "working day" if the variable is a weekday, or
  - print "happy day" if the variable is a weekend day.
*/
package main

import "fmt"

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	var day int
	fmt.Println("enter a day of the week (0 for Monday, 6 for Sunday):")
	fmt.Scan(&day)
	switch day {
	case Monday, Tuesday, Wednesday, Thursday, Friday:
		fmt.Println("working day")
	case Saturday, Sunday:
		fmt.Println("happy day")
	default:
		fmt.Println("not a valid day")
	}
}
