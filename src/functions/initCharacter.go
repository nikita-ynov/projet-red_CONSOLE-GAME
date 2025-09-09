package functions

import (
	structure "PROJETRED/structure"
	"fmt"
)

func InitCharacter(u structure.Character, obj []structure.Inventory) structure.Character {
	// // Example
	// var obj1 structure.Inventory = structure.Inventory{Name: "Bubble Gum", ChangeHp: 2, Quantity: 3}
	// fmt.Println(obj1)
	// // {Bubble Gum 2} {Phone 1}

	var player structure.Character = structure.Character{
		Name:      u.Name,
		Class:     u.Class,
		Lvl:       u.Lvl,
		HpMax:     u.HpMax,
		CurrentHp: u.CurrentHp,
		Inventory: obj,
	}
	fmt.Println(player)
	return player
}
