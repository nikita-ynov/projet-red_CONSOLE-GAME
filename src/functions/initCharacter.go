package functions

import (
	structure "PROJETRED/structure"
)

func InitCharacter(name string, class string, maxHp, spawnHp, money int) structure.Character {
	objects := []structure.Inventory{}
	equipements := structure.Equipment{}
	punch := []structure.Skills{{Name: "Punch", Damage: -15}}

	player := structure.Character{
		Name:           name,
		Class:          class,
		Lvl:            1,
		HpMax:          maxHp,
		CurrentHp:      spawnHp,
		ManaMax:        100,
		CurrentMana:    100,
		Money:          money,
		Damage:         -10,
		Initiative:     10,
		InventoryLimit: 10,
		Inventory:      objects,
		Skills:         punch,
		Equipment:      equipements,
	}

	return player
}
