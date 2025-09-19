package main

import "fmt"

func RunEcommCart() {
	cartItems := []Item{
		{Name: "Book", Price: 200},
		{Name: "Pen", Price: 50},
		{Name: "Laptop", Price: 800},
		{Name: "Headphones", Price: 300},
		{Name: "Chair", Price: 400},
		{Name: "Table", Price: 600},
		{Name: "Mobile Phone", Price: 1200},
		{Name: "Keyboard", Price: 150},
		{Name: "Mouse", Price: 100},
		{Name: "Monitor", Price: 900},
	}

	cart := InitializeCart()

	for i := range cartItems {
		cart.AddItem(cartItems[i])
	}

	total := cart.TotalPrice()
	fmt.Println("Total of cart item is:", total)
}
