package functions

import (
	structure "PROJETRED/structure"
	"fmt"
)

func DisplayInfo(player structure.Character) {
	fmt.Println("====== DISPLAY INFO - " + player.Name + " ======")
	fmt.Println("Class: " + player.Class)
	fmt.Printf("Lvl: %v", player.Lvl)
	fmt.Printf("\n")
	fmt.Printf("Hp Max: %v", player.HpMax)
	fmt.Printf("\n")
	fmt.Printf("Current Hp: %v", player.CurrentHp)
	fmt.Printf("\n")
	fmt.Println("INVENTORY ITEMS:")
	for _, item := range player.Inventory {
		fmt.Println("------------")
		fmt.Println(item.Name)
		fmt.Printf("Quantity: %v", item.Quantity)
		fmt.Printf("\n")
		if item.ChangeHp < 0 {
			fmt.Printf("Damage Hp: %v", item.ChangeHp)
		} else {
			fmt.Printf("Health Hp: %v", item.ChangeHp)
		}
		fmt.Printf("\n")
		fmt.Println("------------")
	}
}
