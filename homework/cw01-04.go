// This file contains an incomplete program. The function shr1 shifts
// the values pointed to by x, y, and z to the right by one position.
// For example, if the values of x, y, z are 3, 5, 7 respectively,
// then  after shr1, thevalues of x, y, z become 7, 3, 5 respectively.
// Complete the function shr1 such that this program outputs:
//   7 3 5
//   9 4 1

// package main

// import "fmt"

// func shr1(x *int, y *int, z *int) {
// 	// -- <code begin> --
// 	*x, *y, *z = *z, *x, *y
// 	// -- <code end> --
// }

// func main() {
// 	x, y, z := 3, 5, 7
// 	shr1(&x, &y, &z)
// 	fmt.Println(x, y, z)

// 	x, y, z = 4, 1, 9
// 	shr1(&x, &y, &z)
// 	fmt.Println(x, y, z)
// }
