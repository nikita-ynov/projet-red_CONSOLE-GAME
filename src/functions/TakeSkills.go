package functions

import (
	"PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
)

func TakeSkills(player *structure.Character, skill []structure.Inventory) {
	var exit string
	var choice int

	fmt.Print("\033[H\033[2J")
	fmt.Println("====== LEARN SKILL MENU ======")
	for _, item := range player.Inventory {
		if item.UniqueObj > 0 {
			skill = append(skill, item)
		}
	}

	if len(skill) == 0 { //verifie si l'inventaire est vide
		fmt.Printf("You don't have any Spell Book to use\n")
		fmt.Print("Enter any key to exit :   ")
		fmt.Scan(&exit)
		return //si il est vide alors return
	}

	fmt.Println("Choise what Spell Book to learn")
	fmt.Println("0. Exit")

	for choice < 1 || choice > len(skill) {
		fmt.Print("Enter your choice :   ")
		fmt.Scan(&choice)
		if choice == 0 {
			return
		}
	}

	utils.RemoveObj(player, structure.Inventory{})
	fmt.Print("\033[H\033[2J")
	fmt.Println("Used Spell Book: ")
	fmt.Print("Enter any key to exit :   ")
	fmt.Scan(&exit)
}
