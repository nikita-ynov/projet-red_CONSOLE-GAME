package functions

import (
	structure "PROJETRED/structure"
	"fmt"
)

func DisplayInfo(player *structure.Character, equipmentplayer structure.Equipment) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("====== DISPLAY INFO - " + player.Name + " ======")
	fmt.Println("Class: " + player.Class)
	fmt.Printf("Current Hp: %v/%v\n", player.CurrentHp, player.HpMax)
	fmt.Printf("Lvl: %v\n", player.Lvl)
	fmt.Printf("Expireience: %v/%v\n", player.Exp, player.Lvl*10)
	fmt.Printf("Money: %v\n", player.Money)
	fmt.Printf("SKILLS %v:\n", len(player.Skills))
	for i := 0; i < len(player.Skills); i++ {
		fmt.Println("------------")
		fmt.Println(player.Skills[i].Name)
		fmt.Printf("Damage: %v\n", player.Skills[i].Damage)
		fmt.Println("------------")
	}
	fmt.Printf("INVENTORY ITEMS: %v/%v\n", len(player.Inventory), player.InventoryLimit)
	for _, item := range player.Inventory {
		fmt.Println("------------")
		fmt.Println(item.Name)
		fmt.Printf("Quantity: %v", item.Quantity)
		fmt.Printf("\n")
		if item.ChangeHp < 0 {
			fmt.Printf("Damage Hp: %v\n", item.ChangeHp)
		} else if item.ChangeHp > 0 {
			fmt.Printf("Health Hp: %v\n", item.ChangeHp)
		} else if item.UniqueObj > 0 {
			fmt.Println("(can be use one time in the game)")
		}
		fmt.Println("------------")
	}

	var choice string

	fmt.Println("-----------------------")
	fmt.Println("1 - Open equipment menu")
	fmt.Println("other key - exit")
	fmt.Print("enter your choice : ")
	fmt.Scan(&choice)
	if choice == "1" {
		DisplayEquipment(equipmentplayer)
	}
}
