package utils

import (
	"PROJETRED/structure"
)

func InventoryIsAtMaxCapacity(inventory *[]structure.Inventory) bool {
	if len(*inventory) < 10 {
		return false
	} else {
		return true
	}
}
