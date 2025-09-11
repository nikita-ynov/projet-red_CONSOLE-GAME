package functions

import (
	structure "PROJETRED/structure"
	"fmt"
	"time"
)

func PoisonPot(player *structure.Character) { //  fonction qui realise l'effet de poison de Potion de poison
	for i := 0; i < 3; i++ {
		player.CurrentHp -= 10
		fmt.Println("Damage of poison : 10") //  inflige  10 degats
		fmt.Println("current Health :")      //affiche
		fmt.Println(player.CurrentHp)        //les Hp actuels
		time.Sleep(1 * time.Second)          //attend 1 s pour refaire l'effet
	}
}
