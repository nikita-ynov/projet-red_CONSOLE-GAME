package functions

import (
	structure "PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
)

func Takepot(player *structure.Character) {
	fmt.Print("\033[H\033[2J")
	newInventory := []structure.Inventory{} //nouvelle inventaire vide
	hasDone := false                        //permet d'utiliser 1 seule fois le spell
	haslifePotion := false                  //renvoie true si le player a une life potion
	var itemUsed structure.Inventory        //ajout de l'item a utiliser ici

	if len(player.Inventory) == 0 { //verifie si l'inventaire est vide
		fmt.Printf("Empty inventory \n")
		var exit string
		fmt.Print("Enter any key to exit :   ")
		fmt.Scan(&exit)
		return //si il est vide alors return
	}

	for _, item := range player.Inventory { //parcourt l'inventaire du joueur

		if item.Name == "Life Potion" { //verifie si il existe un item Life Potion
			haslifePotion = true
		}

		if item.ChangeHp > 0 && !hasDone {
			if player.CurrentHp+50 > player.HpMax { //verifie que l'ajout de la potion n'excede pas les pv max
				player.CurrentHp = player.HpMax //si c'est le cas alors les pv du joueur = pv max
			} else {
				utils.AddHP(player, 50) //appelle de la func permettant d'ajouter des pv au joueur
			}
			hasDone = true
			itemUsed = item
		}
		if item != itemUsed {
			newInventory = append(newInventory, item)
		}

	}

	if !haslifePotion { //return s'il n'a pas de life potion
		fmt.Printf("Player has no lifepotion in its inventory \n")
		var exit string
		fmt.Print("Enter any key to exit :   ")
		fmt.Scan(&exit)
	}

	fmt.Printf("Player has now %v HP \n", player.CurrentHp)
	var exit string
	fmt.Print("Enter any key to exit :   ")
	fmt.Scan(&exit)
	player.Inventory = newInventory
}
