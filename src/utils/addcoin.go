package utils

import "PROJETRED/structure"

func AddCoin(moneyPlayer *structure.Character) {
	moneyPlayer.Money += moneyPlayer.Money + 10

}
