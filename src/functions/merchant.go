package functions

import (
	"PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
)

func Merchant(player *structure.Character) {
	var exit int = 1
	for exit == 1 {

		fmt.Print("\033[H\033[2J")
		merchantItems := []structure.MerchantItems{
			{
				Name:     "Life Potion",
				ChangeHp: 50,
				Quantity: 1,
				Price:    3,
			}, {
				Name:     "Kill Potion",
				ChangeHp: -50,
				Quantity: 1,
				Price:    7,
			},
			{
				Name:     "SpellBook : Fire Ball",
				ChangeHp: 0,
				Quantity: 1,
				Price:    25,
			},
			{
				Name:     "Wolf fur",
				ChangeHp: 0,
				Quantity: 1,
				Price:    4,
			},
			{
				Name:     "Skin Troll",
				ChangeHp: 0,
				Quantity: 1,
				Price:    7,
			},
			{
				Name:     "Wild boar leather",
				ChangeHp: 0,
				Quantity: 1,
				Price:    3,
			},
			{
				Name:     "Crow Featherl",
				ChangeHp: 0,
				Quantity: 1,
				Price:    1,
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
				return
			}
		}

		selected := merchantItems[choice-1]

		if selected.Price > player.Money {
			fmt.Print("\033[H\033[2J")
			fmt.Print("You don't have enough money for buy this item.\n")
			fmt.Print("Enter any key to quit :   ")
			fmt.Scan(&exit)
			return

		}

		if InventoryIsAtMaxCapacity(&player.Inventory) {
			fmt.Print("You've reached your 10 items inventory limit\n")
			fmt.Print("Press any key to go back to MENU")
			fmt.Scan(&exit)
			return
		}

		utils.RemoveMoney(player, selected.Price)
		utils.AddObj(player, structure.Inventory{
			Name:     selected.Name,
			ChangeHp: selected.ChangeHp,
			Quantity: selected.Quantity,
		})
		fmt.Print("\033[H\033[2J")
		fmt.Printf("====== YOU BOUGHT : %s! ======\n", selected.Name)

		fmt.Print("0. Exit\n1. Buy more\nEnter your choice :   ")
		fmt.Scan(&exit)
	}
}
