package utils

import (
	"PROJETRED/structure"
)

func RemoveObj(player *structure.Character, item structure.Item) {
	arr := []structure.Item{}
	for _, element := range player.Inventory {
		if element.Name == item.Name { // comparaison sur le nom uniquement
			// réduire la quantité
			newQuantity := element.Quantity - item.Quantity
			if newQuantity > 0 {
				element.Quantity = newQuantity
				arr = append(arr, element)
			}
			// si quantité <= 0 → on ne le remet pas (donc supprimé)
		} else {
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

func RemoveMana(player *structure.Character, n int) bool {
	player.CurrentMana += n

	if player.CurrentMana > player.ManaMax {
		player.CurrentMana = player.ManaMax
	}
	if player.CurrentMana <= 0 {
		player.CurrentMana = 0
		return true
	}

	return false
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
