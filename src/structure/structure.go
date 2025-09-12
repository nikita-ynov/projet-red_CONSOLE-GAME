package structure

type Inventory struct {
	Name      string
	ChangeHp  int
	Quantity  int
	UniqueObj int
}

type Skills struct {
	Name    string
	Dammage int
}
type Character struct {
	Name           string
	Class          string
	Lvl            int
	HpMax          int
	CurrentHp      int
	Money          int
	InventoryLimit int
	Inventory      []Inventory
	Skills         []Skills
}

type MerchantItems struct {
	Name      string
	ChangeHp  int
	Quantity  int
	UniqueObj int
	Price     int
}

type Monster struct {
	Name      string
	HpMax     int
	CurrentHp int
	Damage    int
}
