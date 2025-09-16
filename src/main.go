package main

import (
	functions "PROJETRED/functions"
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
			functions.Takepot(&player)
		case 3:
			functions.Merchant(&player)
		case 4:
			functions.TrainingFight(&player)
		case 5:
			functions.Forgeron(&player)
		case 6:
			fmt.Print("====== GOODBYE ======")
			exit = false
		}

	}
}
