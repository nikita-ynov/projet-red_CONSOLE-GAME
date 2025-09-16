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

		// Items disponibles
		merchantItems := []structure.MerchantItems{
			{Name: "Life Potion", ChangeHp: 50, Quantity: 1, Price: 5},
			{Name: "Mana Potion", ChangeManna: 50, Quantity: 1, Price: 5},
			{Name: "Snus", Quantity: 1, Price: 1},
			{Name: "Skin Troll", Quantity: 1, Price: 7},
			{Name: "Wild Boar Leather", Quantity: 1, Price: 3},
			{Name: "Crow Feather", Quantity: 1, Price: 1},
			{Name: "Wolf Fur", Quantity: 1, Price: 1},
		}

		if player.InventoryLimit < 30 {
			merchantItems = append(merchantItems, structure.MerchantItems{
				Name: "Upgrade Inventory (+10 slots)", Quantity: 1, Price: 30,
			})
		}
		if checkSkill(*player, "Fire Ball") {
			merchantItems = append(merchantItems, structure.MerchantItems{
				Name: "Fire Ball", ChangeHp: -20, Quantity: 1, Price: 10,
			})
		}

		// Affichage du marchand
		fmt.Println("\033[1;33m====== MERCHANT ======\033[0m")
		fmt.Printf("Your Money: \033[1;32m%d 💰\033[0m\n", player.Money)
		fmt.Println("0. Exit")
		for index, item := range merchantItems {
			icon := ""
			if item.ChangeHp > 0 {
				icon = fmt.Sprintf("❤️ %+d HP", item.ChangeHp)
			} else if item.ChangeHp < 0 {
				icon = fmt.Sprintf("🔥 %d Damage", item.ChangeHp)
			} else if item.ChangeManna > 0 {
				icon = fmt.Sprintf("💧 %+d Mana", item.ChangeManna)
			}

			if icon != "" {
				fmt.Printf("%d. %-25s [%s] - %d💰\n", index+1, item.Name, icon, item.Price)
			} else {
				fmt.Printf("%d. %-25s - %d💰\n", index+1, item.Name, item.Price)
			}
		}

		// Choix du joueur
		var choice int
		for {
			fmt.Print("\nEnter your choice: ")
			_, err := fmt.Scan(&choice)
			if err != nil {
				fmt.Println("Invalid input, try again.")
				var discard string
				fmt.Scanln(&discard) // vider buffer
				continue
			}

			if choice == 0 {
				return
			}
			if choice < 0 || choice > len(merchantItems) {
				fmt.Println("Invalid choice, try again.")
				continue
			}
			break
		}

		selected := merchantItems[choice-1]

		// Vérification argent
		if selected.Price > player.Money {
			fmt.Print("\033[1;31mYou don't have enough money to buy this item!\033[0m\n")
			utils.Exit()
			continue
		}

		// Vérification inventaire
		if utils.InventoryIsAtMaxCapacity(player) {
			fmt.Print("\033[1;31mYour inventory is full!\033[0m\n")
			utils.Exit()
			continue
		}

		// Achat
		utils.RemoveMoney(player, selected.Price)
		switch selected.Name {
		case "Upgrade Inventory (+10 slots)":
			utils.UpgradeInvenorySlot(player, 10)
		case "Fire Ball":
			utils.AddSkill(player, structure.Skills{Name: selected.Name, Damage: selected.ChangeHp})
		default:
			utils.AddObj(player, structure.Inventory{
				Name:        selected.Name,
				ChangeHp:    selected.ChangeHp,
				ChangeManna: selected.ChangeManna,
				Quantity:    selected.Quantity,
			})
		}

		// Confirmation
		fmt.Printf("\n\033[1;32mYou bought: %s ✅\033[0m\n", selected.Name)
		fmt.Print("0. Exit\n1. Buy more\nEnter your choice: ")
		fmt.Scan(&exit)
	}
}
