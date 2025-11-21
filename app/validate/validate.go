package main

import (
	"fmt"
)

func main() {
	input := 21
	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "is even")
	} else {
		fmt.Println(input, "is odd")
	}

	input = 5
	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "is even")
	} else {
		fmt.Println(input, "is odd") // Added missing else block
	}
}

func validate(input int) error {
	if input < 0 {
		return fmt.Errorf("input must be non-negative")
	}
	return nil
}
