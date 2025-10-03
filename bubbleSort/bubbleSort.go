// Bubble Sort Practice
package main

import "fmt"

func main() {

	num := []int{7, 3, 9, 5}

	for i := 0; i < len(num)-1; i++ {
		for j := 0; j < len(num)-1-i; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
	fmt.Println(num)
}
