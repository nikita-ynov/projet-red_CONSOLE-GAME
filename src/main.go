package main

import (
	functions "PROJETRED/functions"
	"fmt"
)

func main() {
	fmt.Println("====== START GAME ======")

	player := functions.InitCharacter()
	var exit bool = true

	for exit {
		menuChoise := functions.Menu()
		switch menuChoise {
		case 1:
			functions.DisplayInfo(&player)
		case 2:
			functions.Takepot(player)
		case 3:
			functions.Merchant(player)
		case 4:
			fmt.Print("====== GOODBYE ======")
			exit = false
		}

	}
}
