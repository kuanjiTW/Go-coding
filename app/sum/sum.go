package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0
	for i, v := range os.Args {
		if i == 0 {
			continue
		}
		if n, err := strconv.Atoi(v); err == nil {
			sum += n
		}

	}
	fmt.Println(sum)
}
