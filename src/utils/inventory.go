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

func UpgradeInvenorySlot(player *structure.Character, n int) {
	if player.InventoryLimit+n <= 30 {
		player.InventoryLimit += n
	}
}
