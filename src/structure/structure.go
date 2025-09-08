package structure

type Inventory struct {
	Name     string
	ChangeHp int
	Quantity int
}

type Character struct {
	Name      string
	Class     string
	Lvl       int
	HpMax     int
	CurrentHp int
	Inventory []Inventory
}
