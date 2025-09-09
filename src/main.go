package main

import (
	functions "PROJETRED/functions"
	"fmt"
)

func main() {
	fmt.Println("START GAME")
	player := functions.InitCharacter()

	choise := functions.Menu()
	switch choise {
	case 1:
		functions.DisplayInfo(player)
	case 2:
		fmt.Print("takepot function")
	case 3:
		fmt.Print("====== GOODBYE ======")
	}
}
