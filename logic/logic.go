package main

import "fmt"

func main() {
	var balance int = 1000
	var withdraw int
	fmt.Print("Withdraw amount: ")
	fmt.Scan(&withdraw)
	if withdraw < 0 {
		fmt.Println("The amount must be greater than 0")
	} else if withdraw >= balance {
		//the branch
		fmt.Println("The amount must be  smaller than or equal to balance")
	} else {
		//else branch
		balance -= withdraw
		fmt.Print("balance:", balance)
	}
}
