package functions

import (
	structure "PROJETRED/structure"
	"fmt"
	"strings"
)

func CharacterCreation() structure.Character {
	chooseClass := []string{"Human", "Elf", "Dwarf", "Hobbit", "Orcs"}

	var name string
	var choice int
	var class string
	var maxHp int
	var spawnHp int

	// Ask for player name
	fmt.Print("\033[H\033[2J") // Clear console
	fmt.Println("\033[1;36m====== CHARACTER CREATION ======\033[0m")
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	name = StartWithUppercase(name)

	// Choose class
	fmt.Print("\033[H\033[2J")
	fmt.Println("\033[1;36mChoose your class:\033[0m")
	for i, classOption := range chooseClass {
		fmt.Printf("%d - %s\n", i+1, classOption)
	}

	maxChoice := len(chooseClass)
	for choice < 1 || choice > maxChoice {
		fmt.Print("\nEnter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("\033[1;33mInvalid input! Please enter a number.\033[0m")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		if choice < 1 || choice > maxChoice {
			fmt.Println("\033[1;33mInvalid choice! Please select a valid class.\033[0m")
		}
	}

	// Set class-specific stats
	switch choice {
	case 1:
		class = chooseClass[0]
		maxHp = 100
	case 2:
		class = chooseClass[1]
		maxHp = 80
	case 3:
		class = chooseClass[2]
		maxHp = 120
	case 4:
		class = chooseClass[3]
		maxHp = 50
	case 5:
		class = chooseClass[4]
		maxHp = 150
	}

	spawnHp = maxHp / 2

	fmt.Printf("\n\033[1;32mWelcome %s the %s! Your starting HP is %d/%d.\033[0m\n", name, class, spawnHp, maxHp)

	mainCharacter := InitCharacter(name, class, maxHp, spawnHp, 100)

	return mainCharacter
}

// Capitalize first letter, lowercase the rest
func StartWithUppercase(name string) string {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return ""
	}

	first := strings.ToUpper(string(name[0]))
	if len(name) > 1 {
		return first + strings.ToLower(name[1:])
	}
	return first
}
