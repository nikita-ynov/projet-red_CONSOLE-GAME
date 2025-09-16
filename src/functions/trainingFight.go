package functions

import (
	"PROJETRED/structure"
	"fmt"
	"time"
)

func TrainingFight(player *structure.Character) {

	monster := InitGoblin(
		"Crawmerax",
		100,
		100,
		-5,
	)

	var isPlayerTurn bool = true

	fmt.Print("\033[H\033[2J")

	fmt.Println("====== WELCOME TO THE TRAINING ======")

	time.Sleep(2 * time.Second)

	fmt.Print("\033[H\033[2J")

	for player.CurrentHp > 0 || monster.CurrentHp > 0 {

		if isPlayerTurn {

			CharacterTurn()

			isPlayerTurn = !isPlayerTurn
		} else {
			GoblinPattern(player, &monster)

			isPlayerTurn = !isPlayerTurn
		}

	}

	if player.CurrentHp <= 0 {
		fmt.Println("====== YOU DIE ======")
		fmt.Scan("Press any key")
	} else {
		fmt.Println("====== YOU KILL THE MONSTER ======")
		time.Sleep(2 * time.Second)
	}

}
