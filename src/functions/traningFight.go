package functions

import (
	"PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
	"math/rand"
	"time"
)

func goblinAttack(i int, goblin structure.Monster, player *structure.Character) {
	// Goblin attaque
	damage := goblin.Damage
	if i != 0 && i%3 == 0 {
		damage *= 2
		fmt.Printf("\033[1;31m%s attacks %s! Critical hit! (%d damage)\033[0m\n", goblin.Name, player.Name, damage)
	} else {
		fmt.Printf("\033[1;31m%s attacks %s! (%d damage)\033[0m\n", goblin.Name, player.Name, damage)
	}
	utils.RemoveHp(player, damage)
	fmt.Printf("Player HP: %d/%d\n", player.CurrentHp, player.HpMax)
}

func characterAttack(player *structure.Character, goblin *structure.Monster) {

	arr := []*structure.Weapon{}

	//récupere la liste des armes dans l'inventaire
	for _, value := range player.Inventory {
		if value.Name == "Excalibur" {
			toAdd := structure.Weapon{
				Name:        value.Name,
				Damage:      value.ChangeHp,
				LvlRequired: 0,
			}
			arr = append(arr, &toAdd)
		}
	}

	res, weapon := CharacterTurn(arr)
	switch res {
	case "attack":
		utils.MonsterRemoveHp(goblin, player.Damage)
		fmt.Printf("\033[1;32m%s attacks %s! (%d damage)\033[0m\n", player.Name, goblin.Name, player.Damage)
		fmt.Printf("Goblin HP: %d/%d\n", goblin.CurrentHp, goblin.HpMax)
	case "skill":
		if player.CurrentHp > 10 {
			skillName, skillDamage := TakeSkill(player)
			utils.MonsterRemoveHp(goblin, skillDamage)
			utils.RemoveMana(player, -10)
			fmt.Printf("\033[1;34m%s uses skill '%s'! (%d damage)\033[0m\n", player.Name, skillName, skillDamage)
			fmt.Printf("Goblin HP: %d/%d | Mana: %d/%d\n", goblin.CurrentHp, goblin.HpMax, player.CurrentManna, player.MannaMax)

		} else {
			fmt.Println("\033[1;33mNot enough Mana to use a skill!\033[0m")
			characterAttack(player, goblin)
		}
	case "health potion":
		Takepot(player)
	case "weapon":
		fmt.Println("weapon")
		utils.MonsterRemoveHp(goblin, weapon.Damage)
		fmt.Printf("\033[1;32m%s attack %s using %s (%d damage)\033[0m\n", player.Name, goblin.Name, weapon.Name, weapon.Damage)
		fmt.Printf("Goblin HP: %d/%d\n", goblin.CurrentHp, goblin.HpMax)
	}
}

func TrainingFight(player *structure.Character) {
	fmt.Print("\033[H\033[2J")
	goblin := InitGoblin("Training Goblin", 100, 100, -5)
	fmt.Println("\033[1;32m====== START TRAINING ======\033[0m")

	for i := 0; goblin.CurrentHp > 0 && player.CurrentHp > 0; i++ {
		time.Sleep(3 * time.Second)
		if player.Initiative > goblin.Initiative {
			characterAttack(player, &goblin)
			if goblin.CurrentHp <= 0 {
				break
			}
			time.Sleep(1 * time.Second)

			goblinAttack(i, goblin, player)
			if player.CurrentHp <= 0 {
				break
			}
		} else {
			goblinAttack(i, goblin, player)
			if player.CurrentHp <= 0 {
				break
			}

			time.Sleep(1 * time.Second)

			characterAttack(player, &goblin)
			if goblin.CurrentHp <= 0 {
				break
			}
		}
	}

	// Résultat
	if player.CurrentHp <= 0 {
		utils.IsWasted(player)
	} else if goblin.CurrentHp <= 0 {
		weaponDrop := structure.Weapon{
			Name:        "Excalibur",
			Damage:      rand.Intn(50),
			LvlRequired: player.Lvl,
		}
		fmt.Print("====== YOU WIN ======\n")
		utils.AddExp(player, 5)
		fmt.Printf("====== GOBLIN DROP %s ======\n", weaponDrop.Name)
		fmt.Printf("%s damage : %d\n", weaponDrop.Name, weaponDrop.Damage)
		utils.AddObj(
			player,
			structure.Inventory{
				Name:        weaponDrop.Name,
				ChangeHp:    -weaponDrop.Damage,
				ChangeManna: 0,
				Quantity:    1,
			},
		)

	}

	utils.Exit()
}
