package functions

import (
	"PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
	"strconv"
)

func TakeSkill(player *structure.Character) (string, int) {
	var choice string

	// Clear console
	fmt.Print("\033[H\033[2J")
	fmt.Println("\033[1;36m====== USE SKILL ======\033[0m")

	// Build available skills list
	skillsAvailable := player.Skills

	// No skills available
	if len(skillsAvailable) == 0 {
		fmt.Println("\033[1;33mYou don't have any skills to use!\033[0m")
		utils.Exit()
		return "", 0
	}

	// Display skills
	fmt.Println("Choose a skill to use:")
	for i, skill := range skillsAvailable {
		fmt.Printf("%d: %s 🔥 (Damage: %d)\n", i+1, skill.Name, skill.Damage)
	}
	fmt.Println("0: Exit")

	// Read user choice
	for {
		fmt.Print("\nEnter your choice: ")
		fmt.Scan(&choice)

		if choice == "0" {
			return "", 0
		}

		selectedIndex, err := strconv.Atoi(choice)
		if err != nil || selectedIndex < 1 || selectedIndex > len(skillsAvailable) {
			fmt.Println("\033[1;33mInvalid choice! Please try again.\033[0m")
			continue
		}

		chosenSkill := skillsAvailable[selectedIndex-1]
		fmt.Printf("\033[1;34mYou used skill '%s'! Damage: %d 🔥\033[0m\n", chosenSkill.Name, chosenSkill.Damage)
		return chosenSkill.Name, chosenSkill.Damage
	}
}
