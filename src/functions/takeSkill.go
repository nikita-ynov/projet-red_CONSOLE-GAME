package functions

import (
	"PROJETRED/structure"
	"fmt"
	"strconv"
)

func TakeSkill(player *structure.Character) (string, int) {
	var exit string
	var choice string

	// Nettoyage de l'écran
	fmt.Print("\033[H\033[2J")
	fmt.Println("====== USE SKILL ======")

	// Construire la liste des Skill disponibles
	skillsAvailable := []structure.Skills{}
	skillsAvailable = append(skillsAvailable, player.Skills...)

	// Vérifier si aucun sort n’est disponible
	if len(skillsAvailable) == 0 {
		fmt.Println("You don't have any Skill to use")
		fmt.Print("Enter any key to exit: ")
		fmt.Scan(&exit)
		return "", 0 //  return empty values
	}

	// Affichage des Skill disponibles
	fmt.Println("Choose what skill to use:")
	for i, skill := range skillsAvailable {
		fmt.Printf("%d: %s\n", i+1, skill.Name)
	}
	fmt.Println("0: Exit")

	// Lecture du choix utilisateur
	for {
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == "0" {
			return "", 0 //  return empty values if exit
		}

		// Conversion du choix en entier
		selectedIndex, err := strconv.Atoi(choice)
		if err != nil || selectedIndex < 1 || selectedIndex > len(skillsAvailable) {
			fmt.Println("Invalid choice, try again.")
			continue
		}

		chosenSkill := skillsAvailable[selectedIndex-1]
		return chosenSkill.Name, chosenSkill.Dammage
	}
}
