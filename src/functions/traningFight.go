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
	//stock la liste des armes de l'inventaire
	arr := []*structure.Weapon{}

	//récupere la liste des armes dans l'inventaire
	for _, value := range player.Inventory {
		if value.IsWeapon {
			toAdd := structure.Weapon{
				Name:        value.Name,
				Damage:      value.ChangeHp,
				LvlRequired: 0,
			}
			arr = append(arr, &toAdd)
		}
	}
	//verifie quel attaque a été choisie par le player
	res, weapon := CharacterTurn(arr)
	switch res {
	case "melee":
		utils.MonsterRemoveHp(goblin, player.Damage)
		fmt.Printf("\033[1;32m%s attacks %s! (%d damage)\033[0m\n", player.Name, goblin.Name, player.Damage)
		fmt.Printf("Goblin HP: %d/%d\n", goblin.CurrentHp, goblin.HpMax)
	case "spell":
		if player.CurrentMana >= 10 {
			skillName, skillDamage := TakeSkill(player)
			utils.MonsterRemoveHp(goblin, skillDamage)
			utils.RemoveMana(player, -10)
			fmt.Printf("\033[1;34m%s uses skill '%s'! (%d damage)\033[0m\n", player.Name, skillName, skillDamage)
			fmt.Printf("Goblin HP: %d/%d | Mana: %d/%d\n", goblin.CurrentHp, goblin.HpMax, player.CurrentMana, player.ManaMax)

		} else {
			fmt.Println("\033[1;33mNot enough Mana to use a skill!\033[0m")
			characterAttack(player, goblin)
		}
	case "health potion":
		Takepot(player)
	case "weapon":
		utils.MonsterRemoveHp(goblin, weapon.Damage)
		fmt.Printf("\033[1;32m%s attack %s using %s (%d damage)\033[0m\n", player.Name, goblin.Name, weapon.Name, weapon.Damage)
		fmt.Printf("Goblin HP: %d/%d\n", goblin.CurrentHp, goblin.HpMax)
	}
}

func TrainingFight(player *structure.Character) {
	goblin := InitGoblin("Training Goblin", 100, 100, -5)
	fmt.Println("\033[1;32m====== START TRAINING ======\033[0m")

	for i := 0; goblin.CurrentHp > 0 && player.CurrentHp > 0; i++ {
		time.Sleep(1 * time.Second)
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
		player.Deathinarow += 1
		utils.IsWasted(player)
		if player.Deathinarow == 3 {
			UnlockAchievement(player, "The Noob Player", "Die 3 times in a row")
		}
	} else if goblin.CurrentHp <= 0 {
		player.Deathinarow = 0
		randomDamage := rand.Intn(50)

		var randomWeapon = GetRandomWeapon()

		finalDamage := int(randomWeapon.Coef * float32(randomDamage))

		fmt.Print("====== YOU WIN ======\n")

		utils.AddExp(player, 5)
		fmt.Println(" + 5xp")
		utils.AddCoin(player)
		fmt.Println(" + 10 coins")

		fmt.Printf("====== GOBLIN DROP %s ======\n", randomWeapon.Name)
		fmt.Printf("%s damage : %d\n", randomWeapon.Name, finalDamage)
		utils.AddObj(
			player,
			structure.Inventory{
				Name:       randomWeapon.Name,
				ChangeHp:   -finalDamage,
				ChangeMana: 0,
				Quantity:   1,
				IsWeapon:   true,
			},
		)

	}
	// Succès : tuer 10 goblins sans mourir
	player.GoblinsKilledWithoutDying++
	if player.GoblinsKilledWithoutDying >= 10 {
		UnlockAchievement(player, "ez 10 Goblins!", "Slay 10 goblins in a row without dying")
		player.GoblinsKilledWithoutDying = 0 // reset pour un prochain cycle

	}
	if player.Lvl == 10 {
		UnlockAchievement(player, "Confirmed Hero", "Reach level 10")
	}
	utils.Exit()
}
