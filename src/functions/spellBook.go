package functions

import (
	"PROJETRED/structure"
)

func SpellBook(player *structure.Character, book structure.Inventory) {
	switch book.Name {
	case "Spell Book : Fire Ball":
		player.Skills = append(player.Skills, structure.Skills{Name: "Fire Ball", Dammage: 10})
	case "Ice Spike ":
		player.Skills = append(player.Skills, structure.Skills{Name: "Ice Spike", Dammage: 8})
	}
}
