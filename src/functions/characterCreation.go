package functions

import (
	structure "PROJETRED/structure"
	"fmt"
)

func CharacterCreation() structure.Character {
	fmt.Println("===== CHARACTER CREATION =====")
	var name string
	var class string

	fmt.Print("Enter your name :   ")
	fmt.Scan(&name)
	name = StartWithUppercase(name)

	mainCharacter := InitCharacter(name, class)

	return mainCharacter
}

func StartWithUppercase(name string) string {
	//verifie que le string n'est pas vide
	if len(name) == 0 {
		return ""
	}

	newName := ""
	first := name[0]
	//verifie si la premiere lettre du pseudo est en maj et l'a met en maj si necessaire
	if first >= 'a' && first <= 'z' {
		newName += string(first - 32)
	} else {
		newName += string(first)
	}
	//parcourt tout le string et met en minuscule si necessaire
	for i := 1; i < len(name); i++ {
		c := name[i]
		if c >= 'A' && c <= 'Z' {
			newName += string(c + 32)
		} else {
			newName += string(c)
		}
	}

	return newName
}
