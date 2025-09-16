package functions

import (
	structure "PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
)

func Takepot(player *structure.Character) {
	var exit string
	var choice int

	fmt.Print("\033[H\033[2J")
	fmt.Println("====== TAKE HEALTH POTION ======")
	potions := []structure.Inventory{}
	for _, item := range player.Inventory {
		if item.ChangeHp > 0 || item.ChangeManna > 0 {
			potions = append(potions, item)
		}
	}

	if len(potions) == 0 { //verifie si l'inventaire est vide
		fmt.Printf("You don't have any potion for use\n")
		fmt.Print("Enter any key to exit :   ")
		fmt.Scan(&exit)
		return //si il est vide alors return
	}

	fmt.Println("Choise potion")
	fmt.Println("0. Exit")
	for i, potion := range potions {
		if potion.ChangeHp > 0 {
			fmt.Printf("%d. %s (HP %+d) \n", i+1, potion.Name, potion.ChangeHp)
		} else if potion.ChangeManna > 0 {
			fmt.Printf("%d. %s (Manna %+d) \n", i+1, potion.Name, potion.ChangeManna)
		}
	}

	for choice < 1 || choice > len(potions) {
		fmt.Print("Enter your choice :   ")
		fmt.Scan(&choice)
		if choice == 0 {
			return
		}
	}

	if potions[choice-1].ChangeHp > 0 {
		utils.AddHp(player, potions[choice-1].ChangeHp)
	} else if potions[choice-1].ChangeManna > 0 {
		utils.AddManna(player, potions[choice-1].ChangeManna)
	}

	utils.RemoveObj(player, potions[choice-1])
	fmt.Print("\033[H\033[2J")
	fmt.Println("Used potion: " + potions[choice-1].Name)

	if potions[choice-1].ChangeHp > 0 {
		fmt.Printf("You has now %v/%v HP \n", player.CurrentHp, player.HpMax)
	} else if potions[choice-1].ChangeManna > 0 {
		fmt.Printf("You has now %v/%v Manna \n", player.CurrentManna, player.MannaMax)
	}

	fmt.Print("Enter any key to exit :   ")
	fmt.Scan(&exit)
}
