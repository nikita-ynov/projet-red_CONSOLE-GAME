package functions

import (
	"PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
)

func Takepot(player *structure.Character) {
	var choice int

	fmt.Print("\033[H\033[2J")
	fmt.Println("\033[1;36m====== USE POTION ======\033[0m")

	// Filter usable potions
	usableIndexes := []int{}
	for i, item := range player.Inventory {
		if item.ChangeHp > 0 || item.ChangeMana > 0 {
			usableIndexes = append(usableIndexes, i)
		}
	}

	if len(usableIndexes) == 0 {
		fmt.Println("\033[1;33mYou don't have any potions to use!\033[0m")
		utils.Exit()
		return
	}

	// Display menu
	fmt.Println("\nChoose a potion to use:")
	fmt.Println("0. Exit")
	for idx, invIndex := range usableIndexes {
		item := player.Inventory[invIndex]
		if item.ChangeHp > 0 {
			fmt.Printf("%d. %s [❤️ +%d HP]\n", idx+1, item.Name, item.ChangeHp)
		} else if item.ChangeMana > 0 {
			fmt.Printf("%d. %s [💧 +%d Mana]\n", idx+1, item.Name, item.ChangeMana)
		}
	}

	// Player choice
	for {
		fmt.Print("\nEnter your choice: ")
		fmt.Scan(&choice)
		if choice == 0 {
			return
		}
		if choice >= 1 && choice <= len(usableIndexes) {
			break
		}
		fmt.Println("Invalid choice, try again.")
	}

	// Get the actual index in player.Inventory
	selectedIndex := usableIndexes[choice-1]
	selected := player.Inventory[selectedIndex]

	// Apply effect
	if selected.ChangeHp > 0 {
		utils.AddHp(player, selected.ChangeHp)
		fmt.Printf("\033[1;32mYou used %s and recovered %d HP!\033[0m\n", selected.Name, selected.ChangeHp)
		fmt.Printf("Current HP: %d/%d\n", player.CurrentHp, player.HpMax)
	} else if selected.ChangeMana > 0 {
		utils.AddManna(player, selected.ChangeMana)
		fmt.Printf("\033[1;34mYou used %s and recovered %d Mana!\033[0m\n", selected.Name, selected.ChangeMana)
		fmt.Printf("Current Mana: %d/%d\n", player.CurrentMana, player.ManaMax)
	}

	// Remove only 1 quantity
	if player.Inventory[selectedIndex].Quantity > 1 {
		player.Inventory[selectedIndex].Quantity--
	} else {
		player.Inventory = append(player.Inventory[:selectedIndex], player.Inventory[selectedIndex+1:]...)
	}

	utils.Exit()
}
