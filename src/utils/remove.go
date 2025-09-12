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
	if player.CurrentHp-n > 0 {
		player.CurrentHp += n
		return false
	} else {
		player.CurrentHp = 0
		return true
	}
}

func RemoveMoney(player *structure.Character, n int) {
	if player.Money-n >= 0 {
		player.Money -= n
	} else {
		player.Money = 0
	}
}
