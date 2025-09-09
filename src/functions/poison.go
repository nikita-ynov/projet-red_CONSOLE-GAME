package functions

import (
	structure "PROJETRED/structure"
	"fmt"
	"time"
)

func PoisonPot(player *structure.Character) {
	for i := 0; i < 3; i++ {
		fmt.Println("Damage of poison : 10")
		fmt.Println("current Health :")
		fmt.Println(player.CurrentHp - 10)
		time.Sleep(1 * time.Second)
	}
}
