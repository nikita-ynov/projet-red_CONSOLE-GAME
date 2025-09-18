package functions

import (
	structure "PROJETRED/structure"
)

func InitCharacter(name string, class string, maxHp, spawnHp, money int) structure.Character {
	objects := []structure.Item{}
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
		Damage:                  -100,
		Initiative:              10,
		InventoryLimit:          10,
		Inventory:               objects,
		Skills:                  punch,
		Equipment:               equipements,
		AchievementsMenuVisible: false, // menu caché au départ
		Achievements: []structure.Achievement{
			{Name: "ez 10 Goblins!", Description: "Slay 10 goblins in a row without dying", Unlocked: false},
			{Name: "Confirmed Hero", Description: "Reach level 10", Unlocked: false},
			{Name: "Full Armor", Description: "Equip a complete armor set", Unlocked: false},
			{Name: "Poverty", Description: "End up with only 1 lonely coin in your purse", Unlocked: false},
			{Name: "The Noob Player", Description: "Die 3 times in a row", Unlocked: false},
			{Name: "Achievements Hunter Honor", Description: "Complete 5 achievements", Unlocked: false},
			// tout les succées dispo
		},
	}

	return player
}
