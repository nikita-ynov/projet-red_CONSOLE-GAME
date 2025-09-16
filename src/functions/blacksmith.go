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

		WolfFur := 0
		WildBoarLeather := 0
		TrollSkin := 0
		CrowFeather := 0

		//boucle pour compter le nombre de ressources dans l'inventaire du joueur
		for _, item := range player.Inventory {
			switch item.Name {

			case "Crow Feather":
				CrowFeather += item.Quantity
			case "Wolf Fur":
				WolfFur += item.Quantity
			case "Wild boar leather":
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
			}, {
				Name:      "Explorer tunic",
				Price:     5,
				WolfFur:   2,
				TrollSkin: 1,
			}, {
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

		//verifie que le joueur possede bien les ressources necessaires
		if selected.TrollSkin > TrollSkin {
			fmt.Println("You don't have the ressources yet")
			fmt.Printf("You have %v TrollSkin\n", TrollSkin)
			utils.Exit()
			return
		}
		if selected.CrowFeather > CrowFeather {
			fmt.Println("You don't have the ressources yet")
			fmt.Printf("You have %v CrowFeather\n", CrowFeather)
			utils.Exit()
			return
		}
		if selected.WildBoarLeather > WildBoarLeather {
			fmt.Println("You don't have the ressources yet")
			fmt.Printf("You have %v WildBoarLeather\n", WildBoarLeather)
			utils.Exit()
			return
		}
		if selected.WolfFur > WolfFur {
			fmt.Println("You don't have the ressources yet")
			fmt.Printf("You have %v WolfFur\n", WolfFur)
			utils.Exit()
			return
		}

		//verifie que le joueur possede 5 pieces d'or
		if player.Money < 6 {
			fmt.Println("You don't have enough money")
			utils.Exit()
			return
		}
		//verifie que l'ajout de l'équipement n'excede pas la capacité max de l'inventaire
		if utils.InventoryIsAtMaxCapacity(player) {
			fmt.Println("You're inventory is full")
			utils.Exit()
			return
		}

		utils.RemoveMoney(player, 5)

		if selected.CrowFeather > 0 {
			itemsToRemove = append(itemsToRemove, structure.Inventory{
				Name:     "Crow Featherl",
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
				Name:     "Wild boar leather",
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

		utils.AddObj(player, structure.Inventory{
			Name:     selected.Name,
			ChangeHp: 0,
			Quantity: 1,
		})

		fmt.Print("\033[H\033[2J")
		fmt.Printf("====== YOU BOUGHT : %s! ======\n", selected.Name)

		fmt.Print("0. Exit\n1. Buy more\nEnter your choice :   ")
		fmt.Scan(&exit)
	}
}
