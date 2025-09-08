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
	c := character{name: "Kevin", age: 25}
	fmt.Println(DisplayInfo(c))
}
