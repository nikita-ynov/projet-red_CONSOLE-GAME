package functions

import (
	structure "PROJETRED/structure"
	"fmt"
)

func DisplayInfo(player *structure.Character) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("====== DISPLAY INFO - " + player.Name + " ======")
	fmt.Println("Class: " + player.Class)
	fmt.Printf("Lvl: %v\n", player.Lvl)
	fmt.Printf("Hp Max: %v\n", player.HpMax)
	fmt.Printf("Current Hp: %v\n", player.CurrentHp)
	fmt.Printf("Money: %v\n", player.Money)
	fmt.Print("SKILLS :\n")
	for i := 0; i < len(player.Skills); i++ {
		fmt.Println("------------")
		fmt.Println(player.Skills[i].Name)
		fmt.Print("Damage: ")
		fmt.Print(player.Skills[i].Dammage)
		fmt.Println("\n------------")
	}
	fmt.Printf("INVENTORY ITEMS: %v\n", len(player.Inventory))
	for _, item := range player.Inventory {
		fmt.Println("------------")
		fmt.Println(item.Name)
		fmt.Printf("Quantity: %v", item.Quantity)
		fmt.Printf("\n")
		if item.ChangeHp < 0 {
			fmt.Printf("Damage Hp: %v\n", item.ChangeHp)
		} else if item.ChangeHp > 0 {
			fmt.Printf("Health Hp: %v\n", item.ChangeHp)
		}
		fmt.Println("------------")
	}

	var choise int
	fmt.Print("Enter any key to close :   ")
	fmt.Scan(&choise)
}
