package utils

import (
	"PROJETRED/structure"
)

func InventoryIsAtMaxCapacity(player *structure.Character) bool {
	if len(player.Inventory) < player.InventoryLimit {
		return false
	} else {
		return true
	}
}
