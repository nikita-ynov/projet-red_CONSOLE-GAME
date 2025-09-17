package functions

import (
	"PROJETRED/structure"
	"fmt"
)

func CharacterTurn(weapons []*structure.Weapon) (string, structure.Weapon) {

	var choice int
	fmt.Println("\033[1;36m======= PLAYER TURN =======\033[0m")
	fmt.Println("1. Melee ⚔")
	fmt.Println("2. Use Spell 🔥")
	fmt.Println("3. Take Health/Mana Potion 💧")
	for i := range weapons {
		fmt.Printf("%d. Attack using %s (%d damage)\n", i+4, weapons[i].Name, weapons[i].Damage)
	}

	maxChoice := len(weapons) + 4
	var selectedWeapon structure.Weapon
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
			if choice > 3 {
				selectedWeapon = *weapons[choice-4]
				fmt.Println("You choose", selectedWeapon.Name)
			}
			break
		}
		fmt.Println("\033[1;33mInvalid choice! Please select a valid option.\033[0m")
	}

	switch choice {
	case 1:
		return "melee", selectedWeapon
	case 2:
		return "spell", selectedWeapon
	case 3:
		return "health potion", selectedWeapon
	default:
		return "weapon", selectedWeapon
	}
}
