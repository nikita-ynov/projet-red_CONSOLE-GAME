package main

import (
	functions "PROJETRED/functions"
	"PROJETRED/structure"
	"fmt"
)

func main() {
	fmt.Print("\033[H\033[2J")

	fmt.Println("====== START GAME ======")

	player := functions.CharacterCreation()
	var exit bool = true

	for exit {
		fmt.Print("\033[H\033[2J")
		menuChoice := functions.Menu()
		switch menuChoice {
		case 1:
			functions.DisplayInfo(&player)
		case 2:
			functions.TakeSkills(&player) // pour apprendre des skills
		case 3:
			functions.Takepot(&player)
		case 4:
			functions.Merchant(&player)
		case 5:
			functions.GoblinPattern(&player)
		case 6:
			functions.PoisonPot(&player, structure.Inventory{
				Name:     "Test Potion",
				ChangeHp: -30,
				Quantity: 1,
			})
		case 7:
			fmt.Print("====== GOODBYE ======")
			exit = false
		}

	}
}
