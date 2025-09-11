package structure

type Inventory struct {
	Name     string
	ChangeHp int
	Quantity int
}

type Spells struct {
	Name         string
	SpellDammage int
}

type Character struct {
	Name       string
	Class      string
	Lvl        int
	HpMax      int
	CurrentHp  int
	Inventory  []Inventory
	ListSpells []Spells
}

type MerchantItems struct {
	Name             string
	ChangeHp         int
	Quantity         int
	MerchantQuantity int
	Price            int
}
