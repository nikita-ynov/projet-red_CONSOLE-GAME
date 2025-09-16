package functions

import (
	"fmt"
)

func CharacterTurn() string {
	var choice int
	fmt.Println("\033[1;36m======= PLAYER TURN =======\033[0m")
	fmt.Println("1. Attack ⚔")
	fmt.Println("2. Use Skill 🔥")
	fmt.Println("3. Take Health/Mana Potion 💧")

	maxChoice := 3
	for {
		fmt.Print("\nEnter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("\033[1;33mInvalid input! Please enter a number.\033[0m")
			var discard string
			fmt.Scanln(&discard) // clear input buffer
			continue
		}

		if choice >= 1 && choice <= maxChoice {
			break
		}
		fmt.Println("\033[1;33mInvalid choice! Please select a valid option.\033[0m")
	}

	switch choice {
	case 2:
		return "skill"
	case 3:
		return "health potion"
	default:
		return "attack"
	}
}
