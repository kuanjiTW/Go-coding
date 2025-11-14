/*
  (25 points)

  Complete the function `insert` to insert an integer to a sorted
  slice while keeping the resulting slice sorted. You are only
  allowed to add program statements between the following two lines.
    // --- <code begin> ---
    // --- <code end> ---

  Follow the instructions below to complete `insert`.
  1. Declare an integer variable `i` with the initial value 0.
  2. Write a for loop
     - with no initial statement,
     - with condition that holds if and only if `i < len(ns)` and
       `ns[i] <= m`,
     - with post-statement `i++`, and
     - with no loop body.
     This loop is used to find the index of the first element in
     `ns` that is greater than `m`. If there is no such element in
     `ns`, the index will be `len(ns)`.
  3. After the for loop, append the first i elements of `ns` to
     `ms` (without using any loop).
     (Hint: use the format `ns[j:k]` to read the elements of `ns`
            from index j (inclusive) to index k (exclusive).)
  4. Append `m` to `ms`.
  5. Append the last len(ns) - i elements of `ns` to `ms` (without
     using any loop).
*/

package main

import (
	"fmt"
	"math/rand"
)

func insert(m int, ns []int) []int {
	ms := []int{}
	// --- <code begin> ---
	i := 0
	for i < len(ns) && ns[i] <= m {
		i++
	}
	ms = append(ms, ns[:i]...)
	ms = append(ms, m)
	ms = append(ms, ns[i:]...)
	// --- <code end> ---
	return ms
}

func main() {
	ns := []int{}
	for i := 0; i < 10; i++ {
		m := rand.Intn(100)
		ns = insert(m, ns)
		fmt.Printf("Inserted %2d => %v\n", m, ns)
	}
}
