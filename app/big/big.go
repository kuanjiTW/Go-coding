package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	var intA int = math.MaxInt64
	intA++
	bigA := big.NewInt(int64(math.MaxInt64))
	bigA.Add(bigA, big.NewInt(1))
	fmt.Println(math.MaxInt64)
	fmt.Println(intA)
	fmt.Println(bigA)

}
