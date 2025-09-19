package functions

import (
	"PROJETRED/image"
	structure "PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
)

// Vérifie si le joueur possède déjà le skill
func checkSkill(player *structure.Character, name string) bool {
	for _, el := range player.Skills {
		if el.Name == name {
			return false
		}
	}
	return true
}

func Merchant(player *structure.Character) {

	exit := 1
	fmt.Println("")
	for exit == 1 {
		fmt.Print("\033[H\033[2J") // Clear console
		image.MerchantImage()

		// Items de base
		baseItems := []structure.MerchantItems{
			{Name: "Life Potion", ChangeHp: 50, Quantity: 1, Price: 5},
			{Name: "Mana Potion", ChangeMana: 50, Quantity: 1, Price: 5},
			{Name: "Snus", Quantity: 1, Price: 1},
			{Name: "Skin Troll", Quantity: 1, Price: 7},
			{Name: "Wild Boar Leather", Quantity: 1, Price: 3},
			{Name: "Crow Feather", Quantity: 1, Price: 1},
			{Name: "Wolf Fur", Quantity: 1, Price: 1},
			{Name: "Fire Ball", ChangeHp: -20, Quantity: 1, Price: 10, IsSpell: true},
		}

		merchantItems := []structure.MerchantItems{}

		// Filtrer les compétences déjà connues

		for _, item := range baseItems {
			if !utils.IsitemInInventory(player.Skills, item) {
				merchantItems = append(merchantItems, item)
			}
		}

		// Extension inventaire
		if player.InventoryLimit < 30 {
			merchantItems = append(merchantItems, structure.MerchantItems{
				Name: "Upgrade Inventory (+10 slots)", Quantity: 1, Price: 30,
			})
		}

		// Skills spéciaux liés aux succès
		achievementsToSkills := []struct {
			AchievementName string
			SkillName       string
			SkillDamage     int
			SkillPrice      int
		}{
			{"Full Armor", "Ice Spike", -30, 15},
			{"ez 10 Goblins!", "Thunder Strike", -25, 12},
			{"Confirmed Hero", "Earth Smash", -40, 20},
		}

		for _, mapping := range achievementsToSkills {
			// Vérifier si succès débloqué et skill non acquis
			unlocked := false
			for _, ach := range player.Achievements {
				if ach.Name == mapping.AchievementName && ach.Unlocked {
					unlocked = true
					break
				}
			}
			if unlocked && checkSkill(player, mapping.SkillName) {
				fmt.Printf("\033[1;33m🎉 Congrats! '%s' skill is now available!\033[0m\n", mapping.SkillName)
				merchantItems = append(merchantItems, structure.MerchantItems{
					Name:     mapping.SkillName,
					ChangeHp: mapping.SkillDamage,
					Quantity: 1,
					Price:    mapping.SkillPrice,
				})
			}
		}

		// Affichage marchand
		fmt.Println("\033[1;33m====== MERCHANT ======\033[0m")
		fmt.Printf("Your Money: \033[1;32m%d 💰\033[0m\n", player.Money)
		fmt.Println("0. Exit")
		for i, item := range merchantItems {
			icon := ""
			if item.ChangeHp > 0 {
				icon = fmt.Sprintf("❤️ %+d HP", item.ChangeHp)
			} else if item.ChangeHp < 0 {
				icon = fmt.Sprintf("🔥 %d Damage", item.ChangeHp)
			} else if item.ChangeMana > 0 {
				icon = fmt.Sprintf("💧 %+d Mana", item.ChangeMana)
			}

			if icon != "" {
				fmt.Printf("%d. %-25s [%s] - %d💰\n", i+1, item.Name, icon, item.Price)
			} else {
				fmt.Printf("%d. %-25s - %d💰\n", i+1, item.Name, item.Price)
			}
		}

		// Choix joueur
		var choice int
		for {
			fmt.Print("\nEnter your choice: ")
			_, err := fmt.Scan(&choice)
			if err != nil {
				fmt.Println("Invalid input, try again.")
				var discard string
				fmt.Scanln(&discard)
				continue
			}
			if choice == 0 {
				return
			}
			if choice < 0 || choice > len(merchantItems) {
				fmt.Println("Invalid choice, try again.")
				continue
			}
			break
		}

		selected := merchantItems[choice-1]

		// Vérification argent
		if selected.Price > player.Money {
			fmt.Print("\033[1;31mYou don't have enough money!\033[0m\n")
			utils.Exit()
			continue
		}

		// Vérification inventaire
		if utils.InventoryIsAtMaxCapacity(player) && selected.Name != "Fire Ball" && selected.Name != "Upgrade Inventory (+10 slots)" {
			fmt.Print("\033[1;31mYour inventory is full!\033[0m\n")
			utils.Exit()
			continue
		}

		// Achat
		utils.RemoveMoney(player, selected.Price)
		if player.Money == 1 {
			UnlockAchievement(player, "Poverty", "End up with only 1 lonely coin in your purse")
		}

		switch selected.Name {
		case "Upgrade Inventory (+10 slots)":
			utils.UpgradeInvenorySlot(player, 10)
		case "Fire Ball", "Ice Spike", "Thunder Strike", "Earth Smash":
			utils.AddSkill(player, structure.Skills{Name: selected.Name, Damage: selected.ChangeHp})
			fmt.Printf("\033[1;32mSkill '%s' learned! ✅\033[0m\n", selected.Name)
		default:
			utils.AddObj(player, structure.Item{
				Name:       selected.Name,
				ChangeHp:   selected.ChangeHp,
				ChangeMana: selected.ChangeMana,
				Quantity:   selected.Quantity,
			})
		}

		// Confirmation
		fmt.Printf("\n\033[1;32mYou bought: %s ✅\033[0m\n", selected.Name)
		fmt.Print("0. Exit\n1. Buy more\nEnter your choice: ")
		fmt.Scan(&exit)
	}
}
