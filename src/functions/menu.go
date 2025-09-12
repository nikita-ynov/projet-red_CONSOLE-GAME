package functions

import (
	"fmt"
)

func Menu() int {
	fmt.Print("====== MENU ======\n")
	fmt.Print("1 - Display Player Info\n")
	fmt.Print("2 - Inventory Player\n")
	fmt.Print("3 - Take Health Potion\n")
	fmt.Print("4 - Merchant\n")
	fmt.Print("5 - Test funcs\n")
	fmt.Print("6 - Exit\n")

	var maxChoice int = 5
	var choice int
	for choice < 1 || choice > maxChoice {
		fmt.Print("Enter your choice :   ")
		fmt.Scan(&choice)
	}

	return choice
}
