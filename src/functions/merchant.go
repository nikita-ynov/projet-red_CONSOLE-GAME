package functions

import (
	"PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
)

func checkSkill(player structure.Character, name string) bool {
	for _, el := range player.Skills {
		if el.Name == name {
			return false
		}
	}
	return true
}

func Merchant(player *structure.Character) {
	var exit int = 1
	for exit == 1 {

		fmt.Print("\033[H\033[2J")
		merchantItems := []structure.MerchantItems{
			{
				Name:        "Life Potion",
				ChangeHp:    50,
				ChangeManna: 0,
				Quantity:    1,
				Price:       0,
			},
			{
				Name:        "Manna Potion",
				ChangeHp:    0,
				ChangeManna: 50,
				Quantity:    1,
				Price:       0,
			}, {
				Name:        "Kill Potion",
				ChangeHp:    -50,
				ChangeManna: 0,
				Quantity:    1,
				Price:       0,
			}, {
				Name:        "Snus",
				ChangeHp:    0,
				ChangeManna: 0,
				Quantity:    1,
				Price:       0,
			},
			{
				Name:        "Skin Troll",
				ChangeHp:    0,
				ChangeManna: 0,
				Quantity:    1,
				Price:       7,
			},
			{
				Name:        "Wild boar leather",
				ChangeHp:    0,
				ChangeManna: 0,
				Quantity:    1,
				Price:       3,
			},
			{
				Name:        "Crow Featherl",
				ChangeHp:    0,
				ChangeManna: 0,
				Quantity:    1,
				Price:       1,
			},
			{
				Name:        "Wolf Fur",
				ChangeHp:    0,
				ChangeManna: 0,
				Quantity:    1,
				Price:       1,
			},
		}

		if player.InventoryLimit < 30 {
			merchantItems = append(merchantItems, structure.MerchantItems{
				Name:        "Upgrade Inventory (+10 slot)",
				ChangeHp:    0,
				ChangeManna: 0,
				Quantity:    1,
				Price:       30,
			})
		}

		if checkSkill(*player, "Fire Ball") {
			merchantItems = append(merchantItems, structure.MerchantItems{
				Name:        "Fire Ball", // skill
				ChangeHp:    -20,
				ChangeManna: 0,
				Quantity:    1,
				Price:       4,
			})
		}

		fmt.Println("====== MERCHANT ======")
		fmt.Println("0. Exit")
		for index, merchantItem := range merchantItems {
			if merchantItem.ChangeHp > 0 {
				fmt.Printf("%d. %s (HP %+d) - %d$\n", index+1, merchantItem.Name, merchantItem.ChangeHp, merchantItem.Price)
			} else if merchantItem.ChangeHp < 0 {
				fmt.Printf("%d. %s (HP %-d) - %d$\n", index+1, merchantItem.Name, merchantItem.ChangeHp, merchantItem.Price)
			} else if merchantItem.ChangeManna > 0 {
				fmt.Printf("%d. %s (Manna %+d) - %d$\n", index+1, merchantItem.Name, merchantItem.ChangeManna, merchantItem.Price)
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

		if utils.InventoryIsAtMaxCapacity(player) {
			fmt.Print("You've reached your 10 items inventory limit\n")
			fmt.Print("Press any key to go back to MENU")
			fmt.Scan(&exit)
			return
		}

		utils.RemoveMoney(player, selected.Price)
		if selected.Name == "Upgrade Inventory (+10 slot)" {
			utils.UpgradeInvenorySlot(player, 10)
		} else if selected.Name == "Fire Ball" {
			utils.AddSkill(player, structure.Skills{
				Name:   selected.Name,
				Damage: selected.ChangeHp,
			})
		} else {
			utils.AddObj(player, structure.Inventory{
				Name:        selected.Name,
				ChangeHp:    selected.ChangeHp,
				ChangeManna: selected.ChangeManna,
				Quantity:    selected.Quantity,
				UniqueObj:   selected.UniqueObj,
			})
		}
		fmt.Print("\033[H\033[2J")
		fmt.Printf("====== YOU BOUGHT : %s! ======\n", selected.Name)

		fmt.Print("0. Exit\n1. Buy more\nEnter your choice :   ")
		fmt.Scan(&exit)
	}
}
