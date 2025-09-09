package functions

import (
	structure "PROJETRED/structure"
	"fmt"
)

func InitCharacter() structure.Character {
	// Example
	var obj1 structure.Inventory = structure.Inventory{Name: "Bubble Gum", ChangeHp: 2, Quantity: 3}
	fmt.Println(obj1)
	// {Bubble Gum 2} {Phone 1}

	var player structure.Character = structure.Character{
		Name:      "Tom",
		Class:     "SS+",
		Lvl:       1,
		HpMax:     100,
		CurrentHp: 70,
		Inventory: []structure.Inventory{obj1},
	}
	fmt.Println(player)
	return player
}
