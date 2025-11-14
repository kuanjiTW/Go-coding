package main

import "fmt"

func main() {

	/*
			s1 := []int{1, 3, 5, 7}
			s2 := []int{2, 4, 6, 8}
			s1 = append(s1, 9)
			fmt.Println(s1)
			s1 = append(s1, s2...)
			fmt.Println(s1)

		s1 := []int{1, 3, 5, 7, 9, 11, 13}
		s2 := make([]int, 4, 10)
		copy(s2, s1)
		fmt.Println(s2)

		s3 := []int{2, 4, 6, 8}
		copy(s3, s1)
		fmt.Println(s3)
	*/
	s1 := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}}
	s2 := make([][]int, 4)
	s3 := make([][]int, 4)
	copy(s2, s1)

	for i := 0; i < min(len(s1), len(s2)); i++ {
		s3[i] = make([]int, 3)
		copy(s3[i], s1[i])
	}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	s1[0][1] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

}
