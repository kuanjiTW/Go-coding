package main

import "fmt"

/*
func main(){
	sum := 0
	for i := 1; i <= 10; i++{
		if i%2 == 0 {
			fmt.Println("even")
			sum += i
		}
	}
	fmt.Println(sum)
}
*/

func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		fmt.Println("i = ", i)
		if i%2 == 0 {
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}
