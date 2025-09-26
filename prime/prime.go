package main

import "fmt"

func IsPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int = 10753 - 1
	for ; !IsPrime(n); n-- {
	}
	fmt.Println("Previous prime number is smaller than 10753: ", n)
}
