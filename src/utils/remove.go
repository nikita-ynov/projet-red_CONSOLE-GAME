package utils

import (
	"PROJETRED/structure"
)

func RemoveObj(player *structure.Character, item structure.Inventory) {
	arr := []structure.Inventory{}
	for _, element := range player.Inventory {

		if element.Quantity > 1 && element == item {
			element.Quantity -= 1
			arr = append(arr, element)
		} else if element != item {
			arr = append(arr, element)
		}
	}
	player.Inventory = arr
}

func RemoveHp(player *structure.Character, n int) bool {
	player.CurrentHp += n

	// clamp entre 0 et MaxHp
	if player.CurrentHp > player.HpMax {
		player.CurrentHp = player.HpMax
	}
	if player.CurrentHp <= 0 {
		player.CurrentHp = 0
		return true // mort
	}

	return false // encore en vie
}

func MonsterRemoveHp(monster *structure.Monster, n int) bool {
	monster.CurrentHp += n

	// clamp entre 0 et MaxHp
	if monster.CurrentHp > monster.HpMax {
		monster.CurrentHp = monster.HpMax
	}
	if monster.CurrentHp <= 0 {
		monster.CurrentHp = 0
		return true // mort
	}

	return false // encore en vie
}
func RemoveMoney(player *structure.Character, n int) {
	if player.Money-n >= 0 {
		player.Money -= n
	} else {
		player.Money = 0
	}
}
