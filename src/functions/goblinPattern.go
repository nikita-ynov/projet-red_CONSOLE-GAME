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
			fmt.Printf("%v Default Attack %v (Crit Damage: %v)", goblin.Name, player.Name, damage)
		} else {
			fmt.Printf("%v Default Attack %v (Damage: %v)", goblin.Name, player.Name, damage)
		}
		utils.RemoveHp(player, damage)
		fmt.Printf(" Player HP: %v\n", player.CurrentHp)
		if player.CurrentHp <= 0 {
			break
		}
		// Le joueur attaque le goblin
		time.Sleep(1 * time.Second)
		res := CharacterTurn()
		switch res {
		case "attack":
			utils.MonsterRemoveHp(&goblin, player.Dammage)
			fmt.Printf("%v Default Attack %v (Damage: %v) Goblin HP: %v\n", player.Name, goblin.Name, player.Dammage, goblin.CurrentHp)
		case "skill":
			skillName, skillDamage := TakeSkill(player)
			utils.MonsterRemoveHp(&goblin, skillDamage)
			fmt.Printf("%v use Skill: %v and Attack %v (Damage: %v) Goblin HP: %v\n", player.Name, skillName, goblin.Name, skillDamage, goblin.CurrentHp)
		case "health potion":
			Takepot(player)
		}
		if goblin.CurrentHp <= 0 {
			break
		}
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
