package functions

import (
	"PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
)

func Forgeron(player *structure.Character) {
	var exit int = 1
	for exit == 1 {
		fmt.Print("\033[H\033[2J")

		// Compter les ressources dans l’inventaire
		WolfFur := 0
		WildBoarLeather := 0
		TrollSkin := 0
		CrowFeather := 0

		for _, item := range player.Inventory {
			switch item.Name {
			case "Crow Feather":
				CrowFeather += item.Quantity
			case "Wolf Fur":
				WolfFur += item.Quantity
			case "Wild Boar Leather":
				WildBoarLeather += item.Quantity
			case "Skin Troll":
				TrollSkin += item.Quantity
			}
		}

		itemsToRemove := []structure.Inventory{}

		BlacksmithItems := []structure.BlacksmithItems{
			{
				Name:            "Explorer hat",
				Price:           5,
				CrowFeather:     1,
				WildBoarLeather: 1,
			},
			{
				Name:      "Explorer tunic",
				Price:     5,
				WolfFur:   2,
				TrollSkin: 1,
			},
			{
				Name:            "Explorer boots",
				Price:           5,
				WolfFur:         1,
				WildBoarLeather: 1,
			},
		}

		var choice int
		maxChoice := len(BlacksmithItems)

		fmt.Println("====== BLACK SMITH ======")
		fmt.Println("0. Exit")
		for index, item := range BlacksmithItems {
			fmt.Printf("%v. %s\n", index+1, item.Name)
		}

		for choice < 1 || choice > maxChoice {
			fmt.Print("Enter your choice :   ")
			fmt.Scan(&choice)
			if choice == 0 {
				return
			}
		}

		selected := BlacksmithItems[choice-1]

		// Vérification des ressources
		if selected.TrollSkin > TrollSkin || selected.CrowFeather > CrowFeather ||
			selected.WildBoarLeather > WildBoarLeather || selected.WolfFur > WolfFur {
			fmt.Print("\033[H\033[2J")
			fmt.Println("You don't have the resources yet")
			utils.Exit()
			return
		}

		// Vérification argent
		if player.Money < selected.Price {
			fmt.Print("\033[H\033[2J")
			fmt.Println("You don't have enough money")
			utils.Exit()
			return
		}

		// Vérification capacité inventaire
		if utils.InventoryIsAtMaxCapacity(player) {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Your inventory is full")
			utils.Exit()
			return
		}

		// Retirer argent
		utils.RemoveMoney(player, selected.Price)

		// Retirer les ressources nécessaires
		if selected.CrowFeather > 0 {
			itemsToRemove = append(itemsToRemove, structure.Inventory{
				Name:     "Crow Feather",
				ChangeHp: 0,
				Quantity: selected.CrowFeather,
			})
		}
		if selected.WolfFur > 0 {
			itemsToRemove = append(itemsToRemove, structure.Inventory{
				Name:     "Wolf Fur",
				ChangeHp: 0,
				Quantity: selected.WolfFur,
			})
		}
		if selected.WildBoarLeather > 0 {
			itemsToRemove = append(itemsToRemove, structure.Inventory{
				Name:     "Wild boar Leather",
				ChangeHp: 0,
				Quantity: selected.WildBoarLeather,
			})
		}
		if selected.TrollSkin > 0 {
			itemsToRemove = append(itemsToRemove, structure.Inventory{
				Name:     "Skin Troll",
				ChangeHp: 0,
				Quantity: selected.TrollSkin,
			})
		}

		for _, item := range itemsToRemove {
			utils.RemoveObj(player, item)
		}

		// Définir la protection selon le type d’équipement
		protection := 0
		switch selected.Name {
		case "Explorer hat":
			protection = 5
		case "Explorer tunic":
			protection = 10
		case "Explorer boots":
			protection = 5
		}

		// Ajouter l’équipement dans l’inventaire avec Protection
		utils.AddObj(player, structure.Inventory{
			Name:       selected.Name,
			Quantity:   1,
			ChangeHp:   0,
			Protection: protection,
		})

		fmt.Print("\033[H\033[2J")
		fmt.Printf("====== YOU BOUGHT : %s! ======\n", selected.Name)
		fmt.Print("0. Exit\n1. Buy more\nEnter your choice :   ")
		fmt.Scan(&exit)
	}
}
