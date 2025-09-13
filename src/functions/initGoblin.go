package functions

import (
	structure "PROJETRED/structure"
)

func InitGoblin(name string, maxHp, spawnHp, damage int) structure.Monster {
	return structure.Monster{
		Name:      name,
		HpMax:     maxHp,
		CurrentHp: spawnHp,
		Damage:    damage,
	}
}
