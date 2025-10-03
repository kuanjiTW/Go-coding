package main

import "fmt"

func main() {
	xs := []int{11, 7, 1, 3, 9, 10, 5}
	for i := 0; i < len(xs); i++ {
		for j := i + 1; j < len(xs); j++ {
			if xs[j] < xs[i] {
				xs[i], xs[j] = xs[j], xs[i]
			}
		}
	}
	fmt.Println(xs)

	for i := 0; i < len(xs); i++ {
		smallest := i
		for j := i + 1; j < len(xs); j++ {
			if xs[j] < xs[smallest] {
				smallest = j
			}
		}
	}
	fmt.Println(xs)
}
