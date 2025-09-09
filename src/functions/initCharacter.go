package functions

import (
	structure "PROJETRED/structure"
)

func InitCharacter() structure.Character {
	fmt.Println("===== INIT CHARACTER =====")

	var name string

	fmt.Print("Enter your name :   ")
	fmt.Scan(&name)

	objects := []structure.Inventory{{Name: "Bubble Gum", ChangeHp: 2, Quantity: 3}}

	player := structure.Character{
		Name:      name,
		Class:     "homme",
		Lvl:       0,
		HpMax:     1000,
		CurrentHp: 100,
		Inventory: objects,
	}
	return player
}
