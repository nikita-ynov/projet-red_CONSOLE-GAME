package functions

import (
	"fmt"
)

func Menu() int {
	fmt.Print("====== MENU ======\n")
	fmt.Print("1 - Display Player Info\n")
	fmt.Print("2 - Take Potion\n")
	fmt.Print("3 - Merchant\n")
	fmt.Print("4 - Exit\n")

	var maxChoise int = 4
	var choise int
	for choise < 1 || choise > maxChoise {
		fmt.Print("Enter your choice :   ")
		fmt.Scan(&choise)
	}

	return choise
}
