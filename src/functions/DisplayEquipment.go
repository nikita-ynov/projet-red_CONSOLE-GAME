package functions

import (
	"PROJETRED/structure"
	"fmt"
)

func DisplayEquipment(equipment structure.Equipment) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("====== DISPLAY EQUIPMENT ======")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Helmet Emplacement  :  ", equipment.Helmet)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Breastplate Emplacement  :  ", equipment.Breastplate)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Legwarmer Emplacement  :  ", equipment.Legwarmer)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("----------------------")
}
