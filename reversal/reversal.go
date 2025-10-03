package main

import "fmt"

func main() {
	xs := []int{1, 3, 5, 7, 9}
	fmt.Println("before reversal", xs)
	//		{9,7,5,3,1
	//		left ,right
	for i := 0; i < len(xs)/2; i++ {
		j := len(xs) - 1 - i
		xs[i], xs[j] = xs[j], xs[i]
	}
	fmt.Println("after reversal", xs)

	for left, right := 0, len(xs)-1; left < right; left, right = left+1, right-1 {
		xs[left], xs[right] = xs[right], xs[left]
	}
	fmt.Println("after reversal", xs)
}
