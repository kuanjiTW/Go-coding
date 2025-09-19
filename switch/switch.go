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
	fmt.Print("Enter your role (0-3): ")
	fmt.Scan(&r)
	//solution A
	switch r {
	case Administrator, PowerUser:
		fmt.Println("Access allowed")
	case NormalUser, Guest:
		fmt.Println("Access denied")
	default:
		fmt.Println("Access denied")
	}

	//solution B
	// if r == Administrator || r == PowerUser {
	// 	fmt.Println("Access allowed")
	// } else if r == NormalUser || r == Guest {
	// 	fmt.Println("Access denied")
	// } else {
	// 	fmt.Println("Access denied")
	// }

	var amount int = 7
	if amount < 0 {

	}
}
