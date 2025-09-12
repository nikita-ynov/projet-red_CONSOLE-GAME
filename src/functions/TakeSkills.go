package functions

import (
	structure "PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
	"strconv"
)

func TakeSkills(player *structure.Character) {
	var exit string
	var choice string

	// Nettoyage de l'écran
	fmt.Print("\033[H\033[2J")
	fmt.Println("====== LEARN SKILL MENU ======")

	// Construire la liste des Spell Books disponibles
	skillsAvailable := []structure.Inventory{}
	for _, item := range player.Inventory {
		if item.UniqueObj > 0 { // Filtrer uniquement les Spell Books
			skillsAvailable = append(skillsAvailable, item)
		}
	}

	// Vérifier si aucun sort n’est disponible
	if len(skillsAvailable) == 0 {
		fmt.Println("You don't have any Spell Book to use")
		fmt.Print("Enter any key to exit: ")
		fmt.Scan(&exit)
		return
	}

	// Affichage des Spell Books disponibles
	fmt.Println("Choose what Spell Book to learn:")
	for i, skill := range skillsAvailable {
		fmt.Printf("%d: %s\n", i+1, skill.Name)
	}
	fmt.Println("0: Exit")

	// Lecture du choix utilisateur
	for {
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == "0" {
			return // Sortie
		}

		// Conversion du choix en entier
		selectedIndex, err := strconv.Atoi(choice)
		if err != nil || selectedIndex < 1 || selectedIndex > len(skillsAvailable) {
			fmt.Println("Invalid choice, try again.")
			continue
		}

		// Récupération de l’objet Inventory choisi
		chosenInventory := skillsAvailable[selectedIndex-1]

		// Apprentissage du sort via SpellBook
		SpellBook(player, chosenInventory)

		// Suppression du Spell Book de l’inventaire
		utils.RemoveObj(player, chosenInventory)

		fmt.Printf("You learned: %s\n", chosenInventory.Name)
		break
	}

	fmt.Print("Enter any key to exit: ")
	fmt.Scan(&exit)
}
