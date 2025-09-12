package functions

import (
	"PROJETRED/structure"
	"PROJETRED/utils"
)

func SpellBook(player *structure.Character, nameSpell string) {
	// définition des sorts disponibles
	spells := []structure.Skills{
		{Name: "Fire Ball", Dammage: 10},
	}

	// chercher le sort demandé
	for _, spell := range spells {
		if spell.Name == nameSpell {
			// ajouter le sort au joueur
			player.Skills = append(player.Skills, spell)
			utils.RemoveObj(player, structure.Inventory{})
		}
	}
}
