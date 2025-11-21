package main

import "fmt"

func main() {
	var x1 int16 = 100
	fmt.Println(x1)
	x1 = -7
	fmt.Println(x1)
	// -7 (untyped int constant) as uint8 value in assignment(overflows)
	x1 = 1024
	fmt.Println(x1)
	//1024 (untyped int constant) as int8 value in assignment (overflows)
	var x32 float32 = 100.0 / 30.0
	var x64 float64 = 100.0 / 30.0
	fmt.Println(x32)
	fmt.Println(x64)
	x := 0.1
	y := 0.2
	if x+y == 0.3 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}

	var x2 int8 = 127
	// -128 to 127
	x += 1
	fmt.Println(x2)
	var y2 uint8 = 255
	// 0 to 255
	y2++
	fmt.Println(y2)
}
