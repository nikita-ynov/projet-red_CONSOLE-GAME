package functions

import (
	structure "PROJETRED/structure"
	"fmt"
	"strings"
)

func InitCharacter(name string, class string) structure.Character {
	fmt.Println("===== INIT CHARACTER =====")

	// var class string
	// var name string

	// fmt.Print("Enter your name :   ")
	// fmt.Scan(&name)

	for strings.ToUpper(class) != "M" && strings.ToUpper(class) != "F" {
		fmt.Print("Enter your sex (M - Man or F - Female) :   ")
		fmt.Scan(&class)
	}
	switch strings.ToUpper(class) {
	case "M":
		class = "Man"
	case "F":
		class = "Female"
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
