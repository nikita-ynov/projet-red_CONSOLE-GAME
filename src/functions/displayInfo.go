package functions

import (
	structure "PROJETRED/structure"
	"PROJETRED/utils"
	"fmt"
)

func DisplayInfo(player *structure.Character) {
	fmt.Print("\033[H\033[2J")

	// En-tête avec couleur
	fmt.Printf("\033[1;36m====== CHARACTER INFO - %s ======\033[0m\n", player.Name)
	fmt.Printf("Class        : %s\n", player.Class)
	fmt.Printf("Level        : %d\n", player.Lvl)
	fmt.Printf("Experience   : %d / %d\n", player.Exp, player.Lvl*10)

	// HP avec barre visuelle
	hpBar := progressBar(player.CurrentHp, player.HpMax, 20, "♥")
	fmt.Printf("HP           : %d / %d %s\n", player.CurrentHp, player.HpMax, hpBar)

	// Argent
	fmt.Printf("Money        : %d 💰\n", player.Money)

	// Compétences
	fmt.Printf("\n\033[1;33m-- Skills (%d) --\033[0m\n", len(player.Skills))
	if len(player.Skills) == 0 {
		fmt.Println("   (none)")
	} else {
		for _, skill := range player.Skills {
			fmt.Printf("   • %s (Damage: %d)\n", skill.Name, skill.Damage)
		}
	}

	// Inventaire
	fmt.Printf("\n\033[1;32m-- Inventory (%d/%d) --\033[0m\n", len(player.Inventory), player.InventoryLimit)
	if len(player.Inventory) == 0 {
		fmt.Println("   (empty)")
	} else {
		for _, item := range player.Inventory {
			if !item.IsWeapon {
				fmt.Printf("   • %s x%v", item.Name, item.Quantity)
				if item.ChangeHp < 0 {
					fmt.Printf("  [⚔ %v HP]\n", item.ChangeHp)
				} else if item.ChangeHp > 0 {
					fmt.Printf("  [❤️ +%v HP]\n", item.ChangeHp)
				} else {
					fmt.Println()
				}
			} else {
				fmt.Printf("   • %s x%v", item.Name, item.Quantity)
				if item.ChangeHp < 0 {
					fmt.Printf("  [⚔ %v DMG]\n", item.ChangeHp)
				} else if item.ChangeHp > 0 {
					fmt.Printf("  [❤️ +%v DMG]\n", item.ChangeHp)
				} else {
					fmt.Println()
				}
			}

		}
	}

	fmt.Println("\n===============================")
	utils.Exit()
}

// Petite fonction pour dessiner une barre de progression
func progressBar(current, max, width int, symbol string) string {
	if max == 0 {
		return ""
	}
	ratio := float64(current) / float64(max)
	filled := int(ratio * float64(width))
	bar := ""
	for i := 0; i < width; i++ {
		if i < filled {
			bar += symbol
		} else {
			bar += "-"
		}
	}
	return fmt.Sprintf("[%s]", bar)
}
