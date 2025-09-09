package main

import (
	functions "PROJETRED/functions"
	"fmt"
)

func main() {
	fmt.Println("START GAME")
	player := functions.InitCharacter()
	fmt.Print(player)
	functions.DisplayInfo(player)
}
