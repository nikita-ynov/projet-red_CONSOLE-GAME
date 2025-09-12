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
			case "CrowFeather":
				CrowFeather++
			case "WolfFur":
				WolfFur++
			case "WildBoarLeather":
				WildBoarLeather++
			case "TrollSkin":
				TrollSkin++
			}

		}

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
			fmt.Print("Press any key to go back to the MENU :")
			fmt.Scan(&exit)
			return
		}
		if selected.CrowFeather > CrowFeather {
			fmt.Println("You don't have the ressources yet")
			fmt.Print("Press any key to go back to the MENU :")
			fmt.Scan(&exit)
			return
		}
		if selected.WildBoarLeather > WildBoarLeather {
			fmt.Println("You don't have the ressources yet")
			fmt.Print("Press any key to go back to the MENU :")
			fmt.Scan(&exit)
			return
		}
		if selected.WolfFur > WolfFur {
			fmt.Println("You don't have the ressources yet")
			fmt.Print("Press any key to go back to the MENU :")
			fmt.Scan(&exit)
			return
		}

		//verifie que le joueur possede 5 pieces d'or
		if player.Money < 6 {
			fmt.Println("You don't have enough money")
			return
		}
		//verifie que l'ajout de l'équipement n'excede pas la capacité max de l'inventaire
		if utils.InventoryIsAtMaxCapacity(&player.Inventory) {
			fmt.Println("You're inventory is full")
			return
		}

		utils.RemoveMoney(player, 5)
		utils.AddObj(player, structure.Inventory{
			Name:     selected.Name,
			ChangeHp: 0,
			Quantity: 1,
		})

		fmt.Print("\033[H\033[2J")
		fmt.Printf("====== YOU BOUGHT : %s! ======\n", selected.Name)

		fmt.Print("Press any key to leave")
		fmt.Scan(&exit)
	}
}
