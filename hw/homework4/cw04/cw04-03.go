/*
  (25 points)

  Complete the two functions `insert` and `ppmap` to implement a
  sorted map. The struct `sortedmap` contains two fields:
  - keys: a sorted slice storing the sorted keys
  - kvmap: a map storing all the key-value pairs

  The function `insert` has three input parameters.
  - k: a key to be inserted
  - v: a value to be inserted in association with the given key
  - m: a sorted map
  Follow the instructions below to complete `insert`.
  1. Write an if statement to test whether the key `k` is contained
     in the map `m.kvmap`, and perform the following actions inside
     the if statement if the test fails:
     a) append the key `k` to the slice of keys `m.keys`, and
     b) sort the slice of keys using `sort.Ints(m.keys)`.
  2. After the if statement, insert the key-value pair to the
     map `m.kvmap`.

  Following the instructions below to complete `ppmap`.
  1. Write a `for ... := range ...` loop to iterate through the slice
     of keys `m.keys`.
  2. In each loop iteration, print the key and the value associated
     with the key in the sorted map.

  The expected output of this program is:
  ```
  1 => Alice
  2 => Charlie
  4 => Eason
  5 => Kevin
  8 => Bob
  9 => David
  ```
*/

package main

import (
	"fmt"
	"sort"
)

type sortedmap struct {
	keys  []int
	kvmap map[int]string
}

func insert(k int, v string, m *sortedmap) {
	// -- <code begin> --
	if _, exists := m.kvmap[k]; !exists {
		m.keys = append(m.keys, k)
		sort.Ints(m.keys)
	}
	m.kvmap[k] = v
	// -- <code end> --
}

func ppmap(m *sortedmap) {
	// -- <code begin> --
	for _, key := range m.keys {
		fmt.Printf("%d => %s\n", key, m.kvmap[key])
	}
	// -- <code end> --
}

func main() {
	m := sortedmap{[]int{}, map[int]string{}}
	insert(5, "Kevin", &m)
	insert(1, "Alice", &m)
	insert(8, "Bob", &m)
	insert(2, "Charlie", &m)
	insert(9, "David", &m)
	insert(4, "Eason", &m)
	ppmap(&m)
}
