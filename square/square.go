package main

import "fmt"

func main() {
	// draw a square
	// for i := 0; i < 4; i++ {
	// 	for j := 0; j < 2; j++ {
	// 		fmt.Print("* ")
	// 	}
	// 	fmt.Println()

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i, "x", j, "=")
			if i*j < 10 {
				fmt.Print(" ")
			}
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}
}
