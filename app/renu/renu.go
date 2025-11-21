package main

import "fmt"

func main() {
	// 	username := "hello_你好_word"
	// 	for i := 0; i < len(username); i++ {
	// 		fmt.Print(username[i], " ")
	// 	}
	// 	fmt.Println()
	// 	rs := []rune(username)
	// 	for i := 0; i < len(rs); i++ {
	// 		fmt.Print(string(rs[i]), " ")
	// 	}

	str := "hi_你好"
	for i := 0; i < len(str); i++ {
		fmt.Printf("0x%x", str[i])
	}
	fmt.Println()

	rs := []rune(str)
	for i := 0; i < len(rs); i++ {
		fmt.Print(string(rs[i]), " ")
	}
	fmt.Println()

	for i := 0; i < len(rs); i++ {
		fmt.Print(rs[i], " ")
	}
	fmt.Println()

	for _, v := range str {
		fmt.Print(string(v), " ")
	}
	fmt.Println()

}
