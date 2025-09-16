package functions

import (
	structure "PROJETRED/structure"
)

func InitCharacter(name string, class string, maxHp, spawnHp, money int) structure.Character {
	objects := []structure.Inventory{}
	equipements := []structure.Equipment{}
	punch := []structure.Skills{{Name: "Punch", Damage: -10}}

	player := structure.Character{
		Name:            name,
		Class:           class,
		Lvl:             1,
		HpMax:           maxHp,
		CurrentHp:       spawnHp,
		Money:           money,
		Damage:          -5,
		Initiative:      10,
		InventoryLimit:  10,
		Inventory:       objects,
		Skills:          punch,
		PlayerEquipment: equipements,
	}

	return player
}
