package utils

import (
	"PROJETRED/structure"
	"fmt"
)

func IsWasted(player *structure.Character) {
	fmt.Print("\033[H\033[2J")
	fmt.Print("====== YOU ARE DEAD ======\n")
	player.CurrentHp = player.HpMax / 2
	fmt.Printf("You are rebirth with %v/%v HP\n", player.HpMax/2, player.HpMax)
	player.Deaths++

}
