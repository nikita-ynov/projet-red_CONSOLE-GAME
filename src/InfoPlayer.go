package main

import "fmt"

type item struct {
	name   string
	number int
}

type character struct {
	name      string
	classe    string
	lvl       int
	maxHp     int
	currentHp int
	inventory []item
}

func DisplayInfo(u character) character {
	return u
}

func main() {
	c := character{name: "Kevin", classe: "25", lvl: 1, maxHp: 1, currentHp: 1, inventory: []item{{name: "", number: 0}}}
	fmt.Println(DisplayInfo(c))
}
