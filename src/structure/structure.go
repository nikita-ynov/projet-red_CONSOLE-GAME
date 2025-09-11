package structure

type Inventory struct {
	Name     string
	ChangeHp int
	Quantity int
}

type Spell struct {
	Name    string
	Dammage int
}
type Character struct {
	Name      string
	Class     string
	Lvl       int
	HpMax     int
	CurrentHp int
	Inventory []Inventory
	Skills    []Spell
}

type MerchantItems struct {
	Name     string
	ChangeHp int
	Quantity int
	Price    int
}
