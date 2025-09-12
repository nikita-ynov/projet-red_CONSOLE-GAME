package functions

import (
	"PROJETRED/structure"
)

func InventoryIsAtMaxCapacity(inventory *[]structure.Inventory) bool {
	if len(*inventory) < 1 {
		return false
	} else {
		return true
	}
}
