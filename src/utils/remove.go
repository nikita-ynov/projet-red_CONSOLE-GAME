package utils

import "PROJETRED/structure"

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
