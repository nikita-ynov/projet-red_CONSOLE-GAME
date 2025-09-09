package functions

import (
	structure "PROJETRED/structure"
)

func InitCharacter() structure.Character {
	// Example
	var obj1 structure.Inventory = structure.Inventory{Name: "Bubble Gum", ChangeHp: -2, Quantity: 3}

	var player structure.Character = structure.Character{
		Name:      "Tom",
		Class:     "SS+",
		Lvl:       1,
		HpMax:     100,
		CurrentHp: 70,
		Inventory: []structure.Inventory{obj1},
	}
	return player
}
