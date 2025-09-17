package functions

import (
	"PROJETRED/structure"
	"fmt"
)

func DisplayAchievements(player *structure.Character) {
	fmt.Print("\033[H\033[2J") // Clear console
	fmt.Println("\033[1;36m====== ACHIEVEMENTS ======\033[0m")

	// Afficher les succès
	for _, ach := range player.Achievements {
		status := "Locked"
		if ach.Unlocked {
			status = "Unlocked"
		}
		fmt.Printf("- %s: %s [%s]\n", ach.Name, ach.Description, status)
	}

	fmt.Println("\n==========================")
	fmt.Print("Press Enter to return to main menu...")
	fmt.Scanln() // attend que le joueur appuie sur Enter
	fmt.Scanln() // nécessaire pour consommer le \n restant
}
