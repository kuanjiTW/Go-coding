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
func main1() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	m := map[int]string{0: "alice", 3: "bob", 7: "charlie"}
	for k, v := range m {
		fmt.Println(k, "=>", v)
	}

	s := []int{5, 7, 3, 9}
	fmt.Println(s)
	for _, v := range s {
		fmt.Println("=>", v)
	}
}

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
	main1()
}
