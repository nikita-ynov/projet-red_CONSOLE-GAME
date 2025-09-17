package functions

import (
	structure "PROJETRED/structure"
)

func InitCharacter(name string, class string, maxHp, spawnHp, money int) structure.Character {
	objects := []structure.Inventory{}
	equipements := structure.Equipment{}
	punch := []structure.Skills{{Name: "Punch", Damage: -15}}

	player := structure.Character{
		Name:                    name,
		Class:                   class,
		Lvl:                     1,
		HpMax:                   maxHp,
		CurrentHp:               spawnHp,
		ManaMax:                 100,
		CurrentMana:             100,
		Money:                   money,
		Damage:                  -50,
		Initiative:              10,
		InventoryLimit:          10,
		Inventory:               objects,
		Skills:                  punch,
		Equipment:               equipements,
		AchievementsMenuVisible: false, // menu caché au départ
		Achievements: []structure.Achievement{
			{Name: "First 10 Goblins", Description: "Kill 10 goblins without dying", Unlocked: false},
			// tout les succées dispo
		},
	}

	return player
}
