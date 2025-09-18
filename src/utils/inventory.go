package utils

import (
	"PROJETRED/structure"
	"fmt"
)

func InventoryIsAtMaxCapacity(player *structure.Character) bool {
	if len(player.Inventory) < player.InventoryLimit {
		return false
	} else {
		return true
	}
}

func UpgradeInvenorySlot(player *structure.Character, n int) {
	if player.InventoryLimit+n <= 30 {
		player.InventoryLimit += n
	}
}

func IsitemInInventory(inventory []structure.Skills, item structure.MerchantItems) bool {

	for _, itemToCheck := range inventory {
		if itemToCheck.Name == item.Name && item.IsSpell {
			return true
		}
	}

	return false
}

func OpenInventory(player *structure.Character) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("====== INVENTORY ======")

	if len(player.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")

		return
	}

	for i, item := range player.Inventory {
		icon := ""
		if item.IsWeapon {
			icon = fmt.Sprintf("⚔ %d DMG", item.ChangeHp)
		} else if item.ChangeHp > 0 {
			icon = fmt.Sprintf("❤️ %+d HP", item.ChangeHp)
		} else if item.ChangeHp < 0 {
			icon = fmt.Sprintf("🔥 %d DMG", item.ChangeHp)
		} else if item.ChangeMana > 0 {
			icon = fmt.Sprintf("💧 %+d Mana", item.ChangeMana)
		}

		fmt.Printf("%d. %s x%d [%s]\n", i+1, item.Name, item.Quantity, icon)
	}

	fmt.Println("0. Exit")
	fmt.Print("\nChoose an item number to remove (0 to exit): ")

	var choice int
	fmt.Scan(&choice)
	if choice <= 0 || choice > len(player.Inventory) {

		return
	}

	// Supprimer une unité ou tout l'item si quantity = 1
	selected := player.Inventory[choice-1]
	if selected.Quantity > 1 {
		player.Inventory[choice-1].Quantity--
		fmt.Printf("Removed 1 %s from inventory.\n", selected.Name)
	} else {
		player.Inventory = append(player.Inventory[:choice-1], player.Inventory[choice-1+1:]...)
		fmt.Printf("Removed %s from inventory.\n", selected.Name)
	}

}
