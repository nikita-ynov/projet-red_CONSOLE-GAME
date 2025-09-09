package main

import (
	functions "PROJETRED/functions"
	"fmt"
)

func main() {
	fmt.Print("\033[H\033[2J")

	fmt.Println("====== START GAME ======")

	player := functions.InitCharacter()
	var exit bool = true

	for exit {
		fmt.Print("\033[H\033[2J")
		menuChoice := functions.Menu()
		switch menuChoice {
		case 1:
			functions.DisplayInfo(player)
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
