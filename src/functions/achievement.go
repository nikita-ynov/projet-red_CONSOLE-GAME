package functions

import (
	"PROJETRED/structure"
	"fmt"
)

func UnlockAchievement(player *structure.Character, name string, description string) {
	firstUnlock := false

	// Vérifie si le succès existe déjà
	found := false
	for i := range player.Achievements {
		if player.Achievements[i].Name == name {
			found = true
			if !player.Achievements[i].Unlocked {
				player.Achievements[i].Unlocked = true
				fmt.Printf("\033[1;33mAchievement unlocked: %s! - %s\033[0m\n", name, description)
				firstUnlock = true

			}
			break
		}
	}

	// Si le succès n’existe pas encore, on l’ajoute et le débloque
	if !found {
		player.Achievements = append(player.Achievements, structure.Achievement{
			Name:        name,
			Description: description,
			Unlocked:    true,
		})
		fmt.Printf("\033[1;33mAchievement unlocked: %s! - %s\033[0m\n", name, description)
		firstUnlock = true
	}

	// Affiche le menu Achievements la première fois qu’un succès est débloqué
	if firstUnlock {
		player.AchievementsMenuVisible = true
	}

}
