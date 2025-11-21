package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var seed int64 = 123456789
	rand.Seed(seed)
	fmt.Println(seed)

}
