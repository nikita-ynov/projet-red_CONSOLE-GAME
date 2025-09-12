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
	goblin := InitGoblin("Training Goblin", 100, 100, -5)
	fmt.Print("====== START TRAINING ======\n")

	for i := 0; goblin.CurrentHp > 0 && player.CurrentHp > 0; i++ {
		time.Sleep(3 * time.Second)

		// Goblin attaque
		damage := goblin.Damage
		if i != 0 && i%3 == 0 {
			damage *= 2
			fmt.Printf("%v Attack %v (Crit Damage: %v)", goblin.Name, player.Name, damage)
		} else {
			fmt.Printf("%v Attack %v (Damage: %v)", goblin.Name, player.Name, damage)
		}
		utils.RemoveHp(player, damage)
		fmt.Printf(" Player HP: %v\n", player.CurrentHp)

		// Le joueur attaque le goblin
		time.Sleep(1 * time.Second)
		utils.MonsterRemoveHp(&goblin, -10) // -10 = soin, +10 = dégâts, selon ta logique
		fmt.Printf("%v Attack %v (Damage: %v) Goblin HP: %v\n", player.Name, goblin.Name, -10, goblin.CurrentHp)
	}

	// Résultat
	if player.CurrentHp <= 0 {
		utils.IsWasted(player)
	} else if goblin.CurrentHp <= 0 {
		fmt.Print("====== YOU WIN ======\n")
	}

	fmt.Print("Enter any key to exit :   ")
	fmt.Scan(&exit)
}
