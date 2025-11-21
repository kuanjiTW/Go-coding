package main

import "fmt"

func main() {
	var arr [3]string = [3]string{"John", "Alice", "Bob"}
	var strs = [...]string{"Hello", "World"}
	var nums [4]int = [4]int{1, 2}
	var nums2 [5]int = [5]int{0: 1, 2: 2, 7, 9}
	fmt.Println(arr)
	fmt.Println(strs)
	fmt.Println(nums)
	fmt.Println(nums2)

	ns := [5]int{2, 4, 6, 8, 10}
	i := 0
	i, ns[i] = 4, 777
	fmt.Println(ns)

}
