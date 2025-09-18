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

		itemsToRemove := []structure.Item{}

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

		fmt.Println("====== BLACK SMITH ======")
		fmt.Println("0. Exit")

		// Affichage des ressources actuelles
		fmt.Println("\nYour current resources:")
		fmt.Printf("- Crow Feather: %d\n", CrowFeather)
		fmt.Printf("- Wolf Fur: %d\n", WolfFur)
		fmt.Printf("- Wild Boar Leather: %d\n", WildBoarLeather)
		fmt.Printf("- Skin Troll: %d\n", TrollSkin)
		fmt.Printf("- Gold: %d\n\n", player.Money)

		// Afficher les objets disponibles avec leurs coûts
		for index, item := range BlacksmithItems {
			fmt.Printf("%v. %s (Cost: %d gold)\n", index+1, item.Name, item.Price)

			// Ressources nécessaires
			if item.CrowFeather > 0 {
				color := "\033[32m"
				if CrowFeather < item.CrowFeather {
					color = "\033[31m" // rouge si insuffisant
				}
				fmt.Printf("%s   - %dx Crow Feather\033[0m\n", color, item.CrowFeather)
			}
			if item.WolfFur > 0 {
				color := "\033[32m"
				if WolfFur < item.WolfFur {
					color = "\033[31m"
				}
				fmt.Printf("%s   - %dx Wolf Fur\033[0m\n", color, item.WolfFur)
			}
			if item.WildBoarLeather > 0 {
				color := "\033[32m"
				if WildBoarLeather < item.WildBoarLeather {
					color = "\033[31m"
				}
				fmt.Printf("%s   - %dx Wild Boar Leather\033[0m\n", color, item.WildBoarLeather)
			}
			if item.TrollSkin > 0 {
				color := "\033[32m"
				if TrollSkin < item.TrollSkin {
					color = "\033[31m"
				}
				fmt.Printf("%s   - %dx Skin Troll\033[0m\n", color, item.TrollSkin)
			}
			color := "\033[32m"
			if player.Money < item.Price {
				color = "\033[31m"
			}
			fmt.Printf("%s   - %d gold\033[0m\n\n", color, item.Price)
		}

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(BlacksmithItems) {
			fmt.Println("Invalid choice.")
			continue
		}

		selected := BlacksmithItems[choice-1]

		// Vérification des ressources
		if selected.TrollSkin > TrollSkin || selected.CrowFeather > CrowFeather ||
			selected.WildBoarLeather > WildBoarLeather || selected.WolfFur > WolfFur {
			fmt.Println("\033[31mYou don't have the resources yet.\033[0m")
			utils.Exit()
			return
		}

		// Vérification argent
		if player.Money < selected.Price {
			fmt.Println("\033[31mYou don't have enough money.\033[0m")
			utils.Exit()
			return
		}

		// Retirer argent
		utils.RemoveMoney(player, selected.Price)

		// Retirer les ressources nécessaires
		if selected.CrowFeather > 0 {
			itemsToRemove = append(itemsToRemove, structure.Item{Name: "Crow Feather", Quantity: selected.CrowFeather})
		}
		if selected.WolfFur > 0 {
			itemsToRemove = append(itemsToRemove, structure.Item{Name: "Wolf Fur", Quantity: selected.WolfFur})
		}
		if selected.WildBoarLeather > 0 {
			itemsToRemove = append(itemsToRemove, structure.Item{Name: "Wild Boar Leather", Quantity: selected.WildBoarLeather})
		}
		if selected.TrollSkin > 0 {
			itemsToRemove = append(itemsToRemove, structure.Item{Name: "Skin Troll", Quantity: selected.TrollSkin})
		}

		for _, item := range itemsToRemove {
			utils.RemoveObj(player, item)
		}

		// Protection selon type d’objet
		protection := 0
		switch selected.Name {
		case "Explorer hat":
			protection = 5
		case "Explorer tunic":
			protection = 15
		case "Explorer boots":
			protection = 10
		}

		// Ajouter équipement
		utils.AddObj(player, structure.Item{
			Name:       selected.Name,
			Quantity:   1,
			Protection: protection,
		})

		// Message d’achat en jaune
		fmt.Printf("\n\033[33m====== YOU BOUGHT: %s! ======\033[0m\n", selected.Name)

		// Afficher ressources restantes
		fmt.Println("\nResources left:")
		fmt.Printf("- Crow Feather: %d\n", CrowFeather-selected.CrowFeather)
		fmt.Printf("- Wolf Fur: %d\n", WolfFur-selected.WolfFur)
		fmt.Printf("- Wild Boar Leather: %d\n", WildBoarLeather-selected.WildBoarLeather)
		fmt.Printf("- Skin Troll: %d\n", TrollSkin-selected.TrollSkin)
		fmt.Printf("- Gold: %d\n", player.Money)

		fmt.Print("\n0. Exit\n1. Buy more\nEnter your choice: ")
		fmt.Scan(&exit)
	}
}
