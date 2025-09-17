package functions

import "PROJETRED/structure"

func InitAchievements() []structure.Achievement {
	return []structure.Achievement{
		{Name: "ez 10 Goblins!", Description: "Slay 10 goblins in a row without dying", Unlocked: false},
		{Name: "Confirmed Hero", Description: "Reach level 10", Unlocked: false},
		{Name: "Full Armor", Description: "Equip a complete armor set", Unlocked: false},
		{Name: "Poverty", Description: "End up with only 1 lonely coin in your purse", Unlocked: false},
		{Name: "The Noob Player", Description: "Die 3 times in a row", Unlocked: false},
		{Name: "Achievements Hunter Honor", Description: "Complete 5 achievements", Unlocked: false},
	}
}
