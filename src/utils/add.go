package utils

import (
	"PROJETRED/structure"
)

func AddObj(player *structure.Character, item structure.Inventory) {
	for i, element := range player.Inventory {
		if element.Name == item.Name {
			player.Inventory[i].Quantity += item.Quantity
			return
		}
	}
	player.Inventory = append(player.Inventory, item)
}

func AddSkill(player *structure.Character, skill structure.Skills) {
	for _, element := range player.Skills {
		if element.Name == skill.Name {
			return
		}
	}
	player.Skills = append(player.Skills, skill)
}

func AddHp(player *structure.Character, n int) {
	if n+player.CurrentHp > player.HpMax || n > player.HpMax {
		player.CurrentHp = player.HpMax
	} else {
		player.CurrentHp += n
	}
}

func AddManna(player *structure.Character, n int) {
	if n+player.CurrentMana > player.ManaMax || n > player.ManaMax {
		player.CurrentMana = player.ManaMax
	} else {
		player.CurrentMana += n
	}
}

func AddExp(player *structure.Character, exp int) {
	player.Exp += exp

	if exp >= 10*player.Lvl {
		player.Lvl += 1
	}
}
