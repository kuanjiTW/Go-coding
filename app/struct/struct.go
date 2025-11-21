package main

import "fmt"

func main() {
	type student struct {
		name      string
		birthyear int
		id        string
	}

	s1 := student{name: "Ayna", birthyear: 2000, id: "S001"}
	fmt.Println("student s1:", s1)
	s1.name = "Alice"
	s1.birthyear = 2001
	fmt.Println("name:", s1.name)
	fmt.Println("ID:", s1.id)
}
