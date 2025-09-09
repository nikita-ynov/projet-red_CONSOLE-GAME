package main

import (
	functions "PROJETRED/functions"
	structure "PROJETRED/structure"
	"fmt"
)

var character structure.Character

func main() {
	fmt.Println("START GAME")
	functions.InitCharacter()
}
