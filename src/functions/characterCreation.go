package functions

import (
	structure "PROJETRED/structure"
	"fmt"
)

func CharacterCreation() structure.Character {

	chooseClass := []string{"Human", "Elfe", "Dwarf"}

	var name string
	var choice int
	var class string
	var maxHp int
	var spawnHp int

	fmt.Print("Enter your name :   ")
	fmt.Scan(&name)
	name = StartWithUppercase(name)

	fmt.Print("Choose your Class : \n")
	fmt.Printf("1 - %v : \n", chooseClass[0])
	fmt.Printf("2 - %v : \n", chooseClass[1])
	fmt.Printf("3 - %v : \n", chooseClass[2])

	fmt.Print("Choose your class :  ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		maxHp = 100
		class = chooseClass[0]
	case 2:
		maxHp = 80
		class = chooseClass[1]
	case 3:
		maxHp = 120
		class = chooseClass[2]
	}

	spawnHp = maxHp / 2

	mainCharacter := InitCharacter(name, class, maxHp, spawnHp)

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
