package functions

import (
	"PROJETRED/structure"
	"fmt"
)

// RemoveEquipmentPlayer retire un équipement et le remet dans l’inventaire
func RemoveEquipmentPlayer(player *structure.Character, slot string) {
	switch slot {
	case "helmet":
		if player.Equipment.Helmet.Name != "" {
			// Retirer la protection des HP max
			player.HpMax -= player.Equipment.Helmet.Protection

			player.Inventory = append(player.Inventory, structure.Inventory{
				Name:       player.Equipment.Helmet.Name,
				Protection: player.Equipment.Helmet.Protection,
				Quantity:   1,
			})
			player.Equipment.Helmet = structure.Helmet{}
		}
	case "breastplate":
		if player.Equipment.Breastplate.Name != "" {
			player.HpMax -= player.Equipment.Breastplate.Protection

			player.Inventory = append(player.Inventory, structure.Inventory{
				Name:       player.Equipment.Breastplate.Name,
				Protection: player.Equipment.Breastplate.Protection,
				Quantity:   1,
			})
			player.Equipment.Breastplate = structure.Breastplate{}
		}
	case "legwarmer":
		if player.Equipment.Legwarmer.Name != "" {
			player.HpMax -= player.Equipment.Legwarmer.Protection

			player.Inventory = append(player.Inventory, structure.Inventory{
				Name:       player.Equipment.Legwarmer.Name,
				Protection: player.Equipment.Legwarmer.Protection,
				Quantity:   1,
			})
			player.Equipment.Legwarmer = structure.Legwarmer{}
		}
	}
}

// SwitchEquipmentPlayer gère le changement d’équipement et remet l’ancien dans l’inventaire
func SwitchEquipmentPlayer(item structure.Inventory, player *structure.Character) {
	var oldName string
	var oldProtection int

	switch item.Name {
	case "Explorer hat":
		oldName = player.Equipment.Helmet.Name
		oldProtection = player.Equipment.Helmet.Protection

		player.Equipment.Helmet = structure.Helmet{
			Name:       item.Name,
			Protection: item.Protection,
		}
	case "Explorer tunic":
		oldName = player.Equipment.Breastplate.Name
		oldProtection = player.Equipment.Breastplate.Protection

		player.Equipment.Breastplate = structure.Breastplate{
			Name:       item.Name,
			Protection: item.Protection,
		}
	case "Explorer boots":
		oldName = player.Equipment.Legwarmer.Name
		oldProtection = player.Equipment.Legwarmer.Protection

		player.Equipment.Legwarmer = structure.Legwarmer{
			Name:       item.Name,
			Protection: item.Protection,
		}
	default:
		fmt.Println("Unknown equipment type.")
		return
	}

	// Mise à jour des HP max
	player.HpMax -= oldProtection   // retirer l’ancien bonus
	player.HpMax += item.Protection // ajouter le nouveau

	// Remettre l’ancien équipement dans l’inventaire si ce n’est pas vide
	if oldName != "" {
		player.Inventory = append(player.Inventory, structure.Inventory{
			Name:       oldName,
			Protection: oldProtection,
			Quantity:   1,
		})
		fmt.Printf("You switched %s with %s [+%d HP]\n", oldName, item.Name, item.Protection)
	} else {
		fmt.Printf("You equipped %s [+%d HP]\n", item.Name, item.Protection)
	}
}

// EquipPlayer permet au joueur d'équiper un objet de son inventaire
func EquipPlayer(player *structure.Character) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("====== EQUIP ITEM ======")

	// Filtrer les items équipables (Protection > 0)
	equipable := []structure.Inventory{}
	for _, item := range player.Inventory {
		if item.Protection > 0 {
			equipable = append(equipable, item)
		}
	}

	if len(equipable) == 0 {
		fmt.Println("No equipable items in your inventory.")
		fmt.Println("Press Enter to return...")
		fmt.Scanln()
		return
	}

	// Afficher la liste
	fmt.Println("Items equipable in your inventory:")
	for i, item := range equipable {
		fmt.Printf("%d. %s [+%d HP] (x%d)\n", i+1, item.Name, item.Protection, item.Quantity)
	}

	// Demander choix
	var choice int
	fmt.Print("Choose an item to equip (0 to cancel): ")
	fmt.Scan(&choice)

	if choice <= 0 || choice > len(equipable) {
		fmt.Println("Cancelled.")
		return
	}

	selected := equipable[choice-1]

	// Équiper avec SwitchEquipmentPlayer
	SwitchEquipmentPlayer(selected, player)

	// Retirer du stock (gestion de Quantity)
	for i := range player.Inventory {
		if player.Inventory[i].Name == selected.Name {
			if player.Inventory[i].Quantity > 1 {
				player.Inventory[i].Quantity--
			} else {
				player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			}
			break
		}
	}
	if player.Equipment.Breastplate.Name == "Explorer tunic" && player.Equipment.Helmet.Name == "Explorer hat" && player.Equipment.Legwarmer.Name == "Explorer boots" {
		UnlockAchievement(player, "Full Armor", "Equip a complete armor set")
	}
}
