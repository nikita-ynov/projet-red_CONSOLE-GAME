package functions

import (
	structure "PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
)

func Takepot(player *structure.Character) {
	var choice int

	// Clear console and header
	fmt.Print("\033[H\033[2J")
	fmt.Println("\033[1;36m====== USE POTION ======\033[0m")

	// Filter usable potions
	potions := []structure.Inventory{}
	for _, item := range player.Inventory {
		if item.ChangeHp > 0 || item.ChangeMana > 0 {
			potions = append(potions, item)
		}
	}

	// No potions available
	if len(potions) == 0 {
		fmt.Println("\033[1;33mYou don't have any potions to use!\033[0m")
		utils.Exit()
		return
	}

	// Display potions menu
	fmt.Println("\nChoose a potion to use:")
	fmt.Println("0. Exit")
	for i, potion := range potions {
		if potion.ChangeHp > 0 {
			fmt.Printf("%d. %s [❤️ +%d HP]\n", i+1, potion.Name, potion.ChangeHp)
		} else if potion.ChangeMana > 0 {
			fmt.Printf("%d. %s [💧 +%d Mana]\n", i+1, potion.Name, potion.ChangeMana)
		}
	}
	// Player choice loop
	for choice < 1 || choice > len(potions) {
		fmt.Print("\nEnter your choice: ")
		fmt.Scan(&choice)
		if choice == 0 {
			return
		}
	}

	selected := potions[choice-1]

	// Apply potion effect
	if selected.ChangeHp > 0 {
		utils.AddHp(player, selected.ChangeHp)
		fmt.Printf("\033[1;32mYou used %s and recovered %d HP!\033[0m\n", selected.Name, selected.ChangeHp)
		fmt.Printf("Current HP: %d/%d\n", player.CurrentHp, player.HpMax)
	} else if selected.ChangeMana > 0 {
		utils.AddManna(player, selected.ChangeMana)
		fmt.Printf("\033[1;34mYou used %s and recovered %d Mana!\033[0m\n", selected.Name, selected.ChangeMana)
		fmt.Printf("Current Mana: %d/%d\n", player.CurrentMana, player.ManaMax)
	}

	// Remove potion from inventory
	utils.RemoveObj(player, selected)

	utils.Exit()
}
