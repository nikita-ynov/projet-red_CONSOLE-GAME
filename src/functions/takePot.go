package functions

import structure "PROJETRED/structure"

func Takepot(player structure.Character, effect int) {

	newInventory := []structure.Inventory{}
	hasDone := false
	var itemUsed structure.Inventory

	for _, item := range player.Inventory {

		if item.ChangeHp > 0 && !hasDone {
			if player.CurrentHp+50 > player.HpMax {
				player.CurrentHp = player.HpMax
			} else {
				player.CurrentHp += 50
			}

			itemUsed = item

		}
		if item != itemUsed {
			newInventory = append(newInventory, item)
		}

	}

	player.Inventory = newInventory
}
