package functions

import (
	structure "PROJETRED/structure"
)

func InitCharacter(name string, class string, maxHp int, spawnHp int) structure.Character {
	objects := []structure.Inventory{}

	punch := []structure.Skills{{Name: "Punch", Dammage: 10}}

	player := structure.Character{
		Name:      name,
		Class:     class,
		Lvl:       0,
		HpMax:     maxHp,
		CurrentHp: spawnHp,
		Inventory: objects,
		Skills:    punch,
	}
	return player
}
