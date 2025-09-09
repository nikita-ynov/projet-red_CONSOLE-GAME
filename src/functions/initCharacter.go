package functions

import (
	structure "PROJETRED/structure"
	"fmt"
	"strings"
)

func InitCharacter() structure.Character {
	fmt.Println("===== INIT CHARACTER =====")

	var name string
	var class string

	fmt.Print("Enter your name :   ")
	fmt.Scan(&name)

	for strings.ToUpper(class) != "H" && strings.ToUpper(class) != "F" {
		fmt.Print("Enter your sex (H - Homme or F - Femme) :   ")
		fmt.Scan(&class)
	}
	switch strings.ToUpper(class) {
	case "H":
		class = "Homme"
	case "F":
		class = "Femme"
	}
	objects := []structure.Inventory{}

	player := structure.Character{
		Name:      name,
		Class:     class,
		Lvl:       0,
		HpMax:     1000,
		CurrentHp: 100,
		Inventory: objects,
	}
	return player
}
