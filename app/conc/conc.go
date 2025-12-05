package main

import "fmt"

func hello() {
	fmt.Println("Hello, World!")
}

func main() {
	fmt.Println("Start")
	go hello()
	fmt.Println("End")
}
