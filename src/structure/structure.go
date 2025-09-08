package structure

type Inventory struct {
	name     string
	quantity int
}

type Character struct {
	name      string
	class     string
	lvl       int
	hpMax     int
	hp        int
	inventory []Inventory
}
