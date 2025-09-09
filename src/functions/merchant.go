package functions

import (
	"PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
)

func Merchant(player structure.Character) {

	merchantItems := []structure.MerchantItems{
		{
			Name:     "Life Potion",
			ChangeHp: 50,
			Quantity: 1,
			Price:    0,
		}, {
			Name:     "Kill Potion",
			ChangeHp: -50,
			Quantity: 1,
			Price:    0,
		}, {
			Name:     "Snus",
			ChangeHp: 0,
			Quantity: 1,
			Price:    0,
		},
	}

	fmt.Println("====== MERCHANT ======")
	fmt.Println("0. Exit")
	for index, merchantItem := range merchantItems {
		if merchantItem.ChangeHp > 0 {
			fmt.Printf("%d. %s (HP %+d) - %d$\n", index+1, merchantItem.Name, merchantItem.ChangeHp, merchantItem.Price)
		} else if merchantItem.ChangeHp < 0 {
			fmt.Printf("%d. %s (HP %-d) - %d$\n", index+1, merchantItem.Name, merchantItem.ChangeHp, merchantItem.Price)
		} else {
			fmt.Printf("%d. %s - %d$\n", index+1, merchantItem.Name, merchantItem.Price)
		}
	}

	var choice int
	maxChoice := len(merchantItems)

	for choice < 1 || choice > maxChoice {
		fmt.Print("Enter your choice :   ")
		fmt.Scan(&choice)
		if choice == 0 {
			fmt.Println("Leaving merchant...")
			return
		}
	}

	// Example action when buying:
	selected := merchantItems[choice-1]
	utils.AddObj(&player, structure.Inventory{
		Name:     selected.Name,
		ChangeHp: selected.ChangeHp,
		Quantity: selected.Quantity,
	})

	fmt.Printf("You bought %s!\n", selected.Name)
}
