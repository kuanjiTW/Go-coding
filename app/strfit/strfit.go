package main

import "fmt"

func main() {
	var rstr string = `Hello\n\tworld`
	var istr string = "Hello\n\tworld"
	fmt.Println(rstr)
	fmt.Println(istr)
}
