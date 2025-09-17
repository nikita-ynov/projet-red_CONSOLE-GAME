package functions

import (
	"PROJETRED/structure"
	"fmt"
)

func DisplayAchievements(player *structure.Character) {
	fmt.Print("\033[H\033[2J") // Clear console
	fmt.Println("\033[1;36m====== ACHIEVEMENTS ======\033[0m")

	for _, ach := range player.Achievements {
		if ach.Unlocked {
			// Jaune + check
			fmt.Printf("\033[1;33m- %s ✅ [%s]\033[0m\n", ach.Name, ach.Description)
		} else {
			// Rouge + lock
			fmt.Printf("\033[1;31m- %s 🔒 [%s]\033[0m\n", ach.Name, ach.Description)
		}
	}

	fmt.Println("\n==========================")
	fmt.Print("Press Enter to return to main menu...")
	fmt.Scanln() // attend que le joueur appuie sur Enter
	fmt.Scanln()
}
