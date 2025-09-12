package functions

import (
	structure "PROJETRED/structure"
)

func InitCharacter(name string, class string, maxHp, spawnHp, money int) structure.Character {
	objects := []structure.Inventory{}

	punch := []structure.Skills{{Name: "Punch", Dammage: 10}}

	player := structure.Character{
		Name:           name,
		Class:          class,
		Lvl:            0,
		HpMax:          maxHp,
		CurrentHp:      spawnHp,
		Money:          money,
		InventoryLimit: 10,
		Inventory:      objects,
		Skills:         punch,
	}
	return player
}
