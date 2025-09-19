package main

import "fmt"

func main() {
	const (
		Administrator = iota
		PowerUser
		NormalUser
		Guest
	)
	var r = NormalUser
	switch r {
	case Administrator:
		fmt.Println("Access allowed")
	case PowerUser:
		fmt.Println("Access allowed")
	case NormalUser:
		fmt.Println("Access denied")
	case Guest:
		fmt.Println("Access denied")
		// default:
		// 	fmt.Println("Access denied")
	}
}
