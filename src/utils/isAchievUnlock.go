package utils

import (
	structure "PROJETRED/structure"
)

// PlayerHasAchievement retourne true si le joueur possède le succès passé en paramètre
func PlayerHasAchievement(player structure.Character, achievementName string) bool {
	for _, a := range player.Achievements {
		if a.Name == achievementName && a.Unlocked {
			return true
		}
	}
	return false
}
