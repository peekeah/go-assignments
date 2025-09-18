package main

import "fmt"

func RunBank() {
	bank := make(map[string]*Account)

	bank["A101"] = &Account{Id: "A101", Name: "Kartik Sharma", City: "Bangalore", Age: 23, Balance: 45230.75}
	bank["A102"] = &Account{Id: "A102", Name: "Neha Verma", City: "Mumbai", Age: 32, Balance: 89320.10}
	bank["A103"] = &Account{Id: "A103", Name: "Rohan Singh", City: "Delhi", Age: 24, Balance: 15000.00}
	bank["A104"] = &Account{Id: "A104", Name: "Sneha Iyer", City: "Chennai", Age: 29, Balance: 67050.55}
	bank["A105"] = &Account{Id: "A105", Name: "Amit Patel", City: "Ahmedabad", Age: 35, Balance: 125430.90}

	sender := bank["A101"]

	// Get Balance
	fmt.Println(sender.GetBalance())

	// Deposit Balance
	sender.Deposit(123425)

	// Withdraw
	err := sender.Withdraw(2344223242)
	if err != nil {
		fmt.Println("err:", err)
	}

	// Transfer
	receiver := bank["A103"]
	err = sender.Transfer(receiver, 123245)
	if err != nil {
		fmt.Println("err:", err)
	}
}
