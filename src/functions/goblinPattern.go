package functions

import (
	"PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
	"time"
)

func goblinAttack(i int, goblin structure.Monster, player structure.Character) {
	// Goblin attaque
	damage := goblin.Damage
	if i != 0 && i%3 == 0 {
		damage *= 2
		fmt.Printf("%v Default Attack %v (Crit Damage: %v)", goblin.Name, player.Name, damage)
	} else {
		fmt.Printf("%v Default Attack %v (Damage: %v)", goblin.Name, player.Name, damage)
	}
	utils.RemoveHp(&player, damage)
	fmt.Printf(" Player HP: %v\n", player.CurrentHp)
}

func characterAttack(player *structure.Character, goblin structure.Monster) {
	res := CharacterTurn()
	switch res {
	case "attack":
		utils.MonsterRemoveHp(&goblin, player.Damage)
		fmt.Printf("%v Default Attack %v (Damage: %v) Goblin HP: %v\n", player.Name, goblin.Name, player.Damage, goblin.CurrentHp)
	case "skill":
		if player.CurrentHp > 10 {
			skillName, skillDamage := TakeSkill(player)
			utils.MonsterRemoveHp(&goblin, skillDamage)
			utils.RemoveMana(player, 10)
			fmt.Printf("%v use Skill: %v and Attack %v (Damage: %v) Goblin HP: %v\n", player.Name, skillName, goblin.Name, skillDamage, goblin.CurrentHp)
		} else {
			fmt.Println("You don't have Manna")
			characterAttack(player, goblin)
		}
	case "health potion":
		Takepot(player)
	}
}

func GoblinPattern(player *structure.Character) {
	var exit string
	fmt.Print("====== GOBLIN PATTERN ======\n")
	goblin := InitGoblin("Training Goblin", 100, 100, -5)
	fmt.Print("====== START TRAINING ======\n")

	for i := 0; goblin.CurrentHp > 0 && player.CurrentHp > 0; i++ {
		time.Sleep(3 * time.Second)
		if player.Initiative > goblin.Initiative {
			characterAttack(player, goblin)
			if goblin.CurrentHp <= 0 {
				break
			}
			time.Sleep(1 * time.Second)

			goblinAttack(i, goblin, *player)
			if player.CurrentHp <= 0 {
				break
			}
		} else {
			goblinAttack(i, goblin, *player)
			if player.CurrentHp <= 0 {
				break
			}

			time.Sleep(1 * time.Second)

			characterAttack(player, goblin)
			if goblin.CurrentHp <= 0 {
				break
			}
		}
	}

	// Résultat
	if player.CurrentHp <= 0 {
		utils.IsWasted(player)
	} else if goblin.CurrentHp <= 0 {
		fmt.Print("====== YOU WIN ======\n")
		utils.AddExp(player, 5)
	}

	fmt.Print("Enter any key to exit :   ")
	fmt.Scan(&exit)
}
