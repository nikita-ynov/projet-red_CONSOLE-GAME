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
