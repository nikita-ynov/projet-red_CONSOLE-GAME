package functions

import (
	"PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
)

func DisplayEquipment(player *structure.Character) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("====== DISPLAY EQUIPMENT ======")
	fmt.Println("")

	// Helmet
	if player.Equipment.Helmet.Name == "" {
		fmt.Println("Helmet Emplacement  :  (empty)")
	} else {
		fmt.Printf("Helmet Emplacement  :  %s (+%d HEALTH)\n", player.Equipment.Helmet.Name, player.Equipment.Helmet.Protection)
	}

	fmt.Println("")

	// Breastplate
	if player.Equipment.Breastplate.Name == "" {
		fmt.Println("Breastplate Emplacement  :  (empty)")
	} else {
		fmt.Printf("Breastplate Emplacement  :  %s (+%d HEALTH)\n", player.Equipment.Breastplate.Name, player.Equipment.Breastplate.Protection)
	}

	fmt.Println("")

	// Legwarmer
	if player.Equipment.Legwarmer.Name == "" {
		fmt.Println("Legwarmer Emplacement  :  (empty)")
	} else {
		fmt.Printf("Legwarmer Emplacement  :  %s (+%d HEALTH)\n", player.Equipment.Legwarmer.Name, player.Equipment.Legwarmer.Protection)
	}

	fmt.Println("\n----------------------")

	// Compter les items équipables
	numItemEquipable := 0
	for _, item := range player.Inventory {
		if item.Name == "Explorer hat" || item.Name == "Explorer tunic" || item.Name == "Explorer boots" {
			numItemEquipable++
		}
	}

	fmt.Printf("You have %d items equipable\n", numItemEquipable)
	var choice int
	if numItemEquipable > 0 {
		fmt.Println("Do you want to equip one ?")
		fmt.Println("1 - YES")
		fmt.Println("0 - NO")
		fmt.Scan(&choice)

		if choice == 1 {
			EquipPlayer(player) // Appelle ton menu d’équipement
		}

	} else {

		for choice != 0 {
			fmt.Println("0 - Back to Player Stats :")
			fmt.Scan(&choice)
		}
		return
	}
	utils.Exit()
}
