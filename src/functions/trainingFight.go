package functions

import (
	"PROJETRED/structure"
	"fmt"
	"time"
)

func TrainingFight(player *structure.Character) {

	monster := &structure.Monster{
		Name:       "Crawmerax",
		HpMax:      100,
		CurrentHp:  100,
		Damage:     10,
		Initiative: 0,
	}

	var isPlayerTurn bool = true

	fmt.Print("\033[H\033[2J")

	fmt.Println("====== WELCOME TO THE TRAINING ======")

	time.Sleep(2 * time.Second)

	fmt.Print("\033[H\033[2J")

	for player.CurrentHp > 0 || monster.CurrentHp > 0 {

		if isPlayerTurn {

			CharacterTurn()
			// fmt.Println("====== CHOOSE YOUR ATTACK ======")
			// for i := range player.Skills {
			// 	fmt.Println("------------")
			// 	fmt.Println(i+1, ".", player.Skills[i].Name)
			// 	fmt.Println("------------")
			// }

			// var choice int
			// fmt.Print("Choose the spell you want to use : ")
			// fmt.Scan(&choice)

			// if choice < 1 || choice > len(player.Skills) {
			// 	fmt.Println("Invalid choice")
			// 	return
			// }

			// selected := player.Skills[choice-1]

			// utils.MonsterRemoveHp(monster, selected.Damage)

			// fmt.Printf("MONSTER HAS NOW %d HP\n", monster.CurrentHp)

			// time.Sleep(2 * time.Second)

			isPlayerTurn = !isPlayerTurn
		} else {
			GoblinPattern(player)
			// fmt.Println("====== MONSTER IS ATTACKING YOU ======")

			// time.Sleep(2 * time.Second)

			// utils.RemoveHp(player, inflictedDamage)

			// fmt.Printf("YOU HAVE %d HP\n", player.CurrentHp)

			isPlayerTurn = !isPlayerTurn
		}

	}

	if player.CurrentHp <= 0 {
		fmt.Println("====== YOU DIE ======")
	} else {
		fmt.Println("====== YOU KILL THE MONSTER ======")
	}

}
