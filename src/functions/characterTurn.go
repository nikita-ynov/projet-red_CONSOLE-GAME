package functions

import (
	"fmt"
)

func CharacterTurn() string {
	var choice int
	fmt.Println("======= PLAYER TURN =======")
	fmt.Println("1. Attack")
	fmt.Println("2. Use Skill")
	fmt.Println("3. Take Health Potion")
	maxChoice := 3
	for choice < 1 || choice > maxChoice {
		fmt.Print("Enter your choice :   ")
		fmt.Scan(&choice)

	}
	if choice == 2 {
		return "skill"
	}
	if choice == 3 {
		return "health potion"
	}
	return "attack"
}
