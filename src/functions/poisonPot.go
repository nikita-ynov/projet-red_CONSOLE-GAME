package functions

import (
	"PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
	"time"
)

func PoisonPot(player *structure.Character, potion structure.Inventory) {
	var exit string
	fmt.Print("\033[H\033[2J")
	fmt.Print("====== POISON POT ======\n")
	for i := 0; i < 3; i++ {
		check := utils.RemoveHp(player, potion.ChangeHp)
		fmt.Println(check)
		if check {
			utils.IsWasted(player)
			break
		}
		fmt.Printf("You lost %vHP...\n ", potion.ChangeHp)
		time.Sleep(1 * time.Second)
	}
	fmt.Print("Enter any key to exit :   ")
	fmt.Scan(&exit)
}
