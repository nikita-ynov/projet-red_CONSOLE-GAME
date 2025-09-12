package functions

import (
	"PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
	"time"
)

func GoblinPattern(player *structure.Character) {
	var exit string
	fmt.Print("====== GOBLIN PATTERN ======\n")
	goblin := InitGoblin("Traning Goblin", 100, 100, 5)
	fmt.Print("====== START TRAINING ======\n")
	for i := 0; goblin.CurrentHp == 0 || player.CurrentHp == 0; i++ {
		time.Sleep(1 * time.Second)

		if i%6 == 0 {
			fmt.Printf("%v Attack %v (Crit Damage: %v)", goblin.Name, player.Name, goblin.Damage*2)
			utils.RemoveHp(player, goblin.Damage*2)
		} else {
			fmt.Printf("%v Attack %v (Damage: %v)", goblin.Name, player.Name, goblin.Damage)
			utils.RemoveHp(player, goblin.Damage)
		}

		time.Sleep(1 * time.Second)
		fmt.Printf("%v Attack %v (Damage: %v)", player.Name, goblin.Name, 10)
		utils.MonsterRemoveHp(&goblin, 10)

	}
	if player.CurrentHp <= 0 {
		utils.IsWasted(player)
	} else if goblin.CurrentHp <= 0 {
		fmt.Print("====== YOU ARE WIN ======\n")
	}

	fmt.Print("Enter any key to exit :   ")
	fmt.Scan(&exit)
}
