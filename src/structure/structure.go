package structure

type Inventory struct {
	name     string
	changeHp int
	quantity int
}

type Character struct {
	name      string
	class     string
	lvl       int
	hpMax     int
	currentHp int
	inventory []Inventory
}
