package functions

import (
	"PROJETRED/data"
	"math/rand"
)

func GetRandomWeapon() data.CoefWeapon {

	// Somme des DropRates
	total := 0
	for _, w := range data.WeaponsData {
		total += w.DropRate
	}

	// Tirage aléatoire
	random := rand.Intn(total)

	// Sélection de l'arme
	cumulative := 0
	for _, w := range data.WeaponsData {
		cumulative += w.DropRate
		if random < cumulative {
			return w
		}
	}
	return data.WeaponsData[0]
}
