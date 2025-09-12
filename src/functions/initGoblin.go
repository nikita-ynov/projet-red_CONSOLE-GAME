package functions

import (
	structure "PROJETRED/structure"
)

func InitGoblin(name string, maxHp, spawnHp, damage int) structure.Moster {
	return structure.Moster{
		Name:      name,
		HpMax:     maxHp,
		CurrentHp: spawnHp,
		Damage:    damage,
	}
}
