package main

import "fmt"

func main() {
	m := map[int]string{5: "hello", 8: "world"}
	m[3] = "hi"
	m[4] = ""
	fmt.Println(m[3])
	fmt.Println(m[4])
	if _, e := m[4]; e {
		fmt.Println("m[4]:", m[4])
	}
	if k, e := m[4]; e {
		fmt.Println("m[4]:", k)
	}

}
