package main

import "fmt"

// func main() {
// 	m := map[int]string{10: "1", 2: "warn", 5: "1.2.1"}
// 	for k, v := range m {
// 		fmt.Println(k, " => ", v)
// 	}
// }

func main() {
	items := []string{"Charger", "Laptop", "Tablet"}
	for i, item := range items {
		fmt.Println(i, " => ", item)
	}
	for i := 0; i < len(items); i++ {
		fmt.Println(i, " => ", items[i])
	}
}
