package functions

import (
	"fmt"
)

func CharacterTurn() string {
	var choice int
	fmt.Println("======= PLAYER TURN =======")
	fmt.Println("1. Attack")
	fmt.Println("2. Use Skill")
	maxChoice := 2
	for choice < 1 || choice > maxChoice {
		fmt.Print("Enter your choice :   ")
		fmt.Scan(&choice)

	}
	if choice == 2 {
		return "skill"
	}
	return "attack"
}
