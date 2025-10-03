// Bubble Sort Practice
package main

import "fmt"

func main() {

	xs := []int{7, 3, 9, 5}

	for i := 0; i < len(xs)-1; i++ {
		for j := 0; j < len(xs)-1-i; j++ {
			if xs[j] > xs[j+1] {
				xs[j], xs[j+1] = xs[j+1], xs[j]
			}
			fmt.Println(xs)
		}
	}
	fmt.Println(xs)
	// main2()
}

func main2() {
	xs := []int{7, 3, 9, 5}
	for i := 0; i < len(xs); i++ {
		// first := xs[i]
		for j := i + 1; j < len(xs); j++ {
			// second := xs[j]
			if xs[j] < xs[i] {
				xs[i], xs[j] = xs[j], xs[i]
			}
			fmt.Println(xs)
		}
	}
	fmt.Print(xs)

}
