package utils

import "PROJETRED/structure"

func UseSpell(player *structure.Character, spell structure.Spells) { //AJOUTER LES SORTS
	for _, element := range player.Inventory {
		if element.Name == spell.Name {
			RemoveObjPlayer(player, "SpellBook : Fire Ball")
			// enlever le spell de inventory et ajouter dans listspell
		}
	}

}
