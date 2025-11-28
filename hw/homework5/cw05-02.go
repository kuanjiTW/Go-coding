/*
  (30 points)

  Implement a `concat` function such that it concatenates
  all input lists (in the same order). The `concat` function
  must also satisfy all the following requirements.
  - It is a variadic function.
  - It is a polymorphic function with type parameters.
  - All the input lists and the returned list have the same type.

  For example, if the inputs are three integer lists
  `[]int{1, 2}`, `[]int{3, 4, 5}`, and `[]int{6, 7}`, then
  the output will be an integer list `[]int{1, 2, 3, 4, 5, 6, 7}`.

  If the inputs are two string lists `[]string{"ab", "cd"}` and
  `[]string{"ef", "gh"}`, then the output will be a string list
  `[]string{"ab", "cd", "ef", "gh"}`.

  Hint:
  - Use `append` to concatenate lists.
*/

package main

import "fmt"

// ----- <code begin> -----
func concat[T any](lists ...[]T) []T {
	var result []T
	for _, list := range lists {
		result = append(result, list...)
	}
	return result
}

// ----- <code end> -----

func main() {
	s1 := concat([]int{1, 2}, []int{3, 4, 5}, []int{6, 7})
	fmt.Println(s1) // expected output: [1 2 3 4 5 6 7]

	s2 := concat([]string{"abc", "de"}, []string{"f"}, []string{"gh", "i"})
	fmt.Println(s2) // expected output: [abc de f gh i]
}
