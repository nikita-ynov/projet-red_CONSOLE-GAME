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
