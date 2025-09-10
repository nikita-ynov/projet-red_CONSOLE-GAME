package functions

import (
	"PROJETRED/structure"
	"fmt"
)

func IsWasted(player *structure.Character) {
	if player.CurrentHp <= 0 {
		fmt.Print("====== YOU ARE DEAD ======")
	}
}
