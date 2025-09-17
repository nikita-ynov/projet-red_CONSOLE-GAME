package functions

import (
	"PROJETRED/structure"
	"fmt"
)

func Menu(player *structure.Character) int {

	fmt.Print("\033[H\033[2J") // Clear console
	fmt.Println("\033[1;36m====== MAIN MENU ======\033[0m")
	fmt.Println("1 - Display Player Info")
	fmt.Println("2 - Use Health/Mana Potion")
	fmt.Println("3 - Visit Merchant")
	fmt.Println("4 - Training Battle")
	fmt.Println("5 - Blacksmith")
	fmt.Println("6 - Exit")
	if player.AchievementsMenuVisible {
		fmt.Println("7 - Achievements")
	}
	var choice int
	for {
		fmt.Print("\nEnter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("\033[1;31mInvalid input! Please enter a number.\033[0m")
			var discard string
			fmt.Scanln(&discard) // clear input buffer
			continue
		}
		if choice >= 1 && choice <= 7 {
			break
		}
		fmt.Println("\033[1;31mInvalid choice! Please choose a valid option.\033[0m")
	}

	return choice
}
