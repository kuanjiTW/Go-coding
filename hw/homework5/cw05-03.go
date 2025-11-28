/*
  (30 points)

  Implement a `gen` function that returns another function,
  such that the returned function returns the next prime
  number (starting from 2) every time it is called.

  The expected output of this program is:
  ```
  01-th prime: 2
  02-th prime: 3
  03-th prime: 5
  04-th prime: 7
  05-th prime: 11
  06-th prime: 13
  07-th prime: 17
  08-th prime: 19
  09-th prime: 23
  10-th prime: 29
  ```

  Hints:
  - Use closure.
  - The function `IsPrime` returns `true` if the input number
    is prime, and otherwise returns `false`.
*/

package main

import (
	"fmt"
	"math/big"
)

func IsPrime(n int64) bool {
	return big.NewInt(n).ProbablyPrime(0)
}

func gen() func() int64 {
	// ----- <code begin> -----
	current := int64(1)
	return func() int64 {
		for {
			current++
			if IsPrime(current) {
				return current
			}
		}
	}
	// ----- <code end> -----
}

func main() {
	primes := gen()
	for i := 0; i < 10; i++ {
		fmt.Printf("%02d-th prime: %d\n", i+1, primes())
	}
}
