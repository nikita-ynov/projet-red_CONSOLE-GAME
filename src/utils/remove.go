package utils

import "PROJETRED/structure"

func RemoveObjMerchant(item *structure.MerchantItems) { // RETIRE DES ITEMS DE MARCHANT
	item.MerchantQuantity -= 1
}

func RemoveObjPlayer(player *structure.Character, item structure.Inventory) { // RETIRE LES OBJETS
	newArray := []structure.Inventory{}
	for _, el := range player.Inventory {
		if el.Name != item.Name {
			newArray = append(newArray, el)
		}
	}
	player.Inventory = newArray

}
